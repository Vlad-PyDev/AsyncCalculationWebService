package agent

import (
	"context"
	"log"
	"time"

	pb "github.com/Vlad-PyDev/AsyncCalculationWebService/api/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func handleStream(client pb.OrchestratorClient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataStream, err := client.Calculate(ctx)
	if err != nil {
		return err
	}

	completionChan := make(chan struct{})
	defer close(completionChan)

	go func() {
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			case <-completionChan:
				return
			default:
				receivedTask, err := dataStream.Recv()
				if err != nil {
					log.Printf("Error receiving task: %v", err)
					return
				}

				tasksCh <- &Task{
					ID:   int(receivedTask.Id),
					Arg1: receivedTask.Arg1,
					Arg2: receivedTask.Arg2,
					Type: receivedTask.Operator,
				}
			}
		}
	}()

	go func() {
		defer cancel()
		for {
			select {
			case taskResult := <-resultsCh:
				err := dataStream.Send(&pb.AgentResponse{
					Id:     int32(taskResult.ID),
					Result: float32(taskResult.Result),
					Error:  taskResult.Error,
				})
				if err != nil {
					log.Printf("Error sending response: %v", err)
					return
				}
			case <-ctx.Done():
				return
			case <-completionChan:
				return
			}
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}

func (a *Agent) Connect() {
	for {
		grpcConn, err := grpc.NewClient(
			a.config.OrchestratorAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Printf("Failed to connect to server: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		grpcClient := pb.NewOrchestratorClient(grpcConn)
		streamErr := handleStream(grpcClient)
		grpcConn.Close()

		if streamErr != nil {
			log.Printf("Stream processing error: %v", streamErr)
		}
		time.Sleep(1 * time.Second)
	}
}
