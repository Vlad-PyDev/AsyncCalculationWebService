package orchestrator

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/Vlad-PyDev/AsyncCalculationWebService/api/gen/go"
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/models"
	"google.golang.org/grpc"
)

const (
	tcp         = "tcp"
	addr string = ":5000"
)

type Server struct {
	pb.UnimplementedOrchestratorServer
	mu sync.Mutex
}

func runGRPC() {
	log.Println("Initializing TCP server...")
	listener, err := net.Listen(tcp, addr)
	if err != nil {
		log.Fatalf("failed to start TCP server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrchestratorServer(grpcServer, NewServer())

	log.Printf("TCP server running on: %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}

func (s *Server) Calculate(stream pb.Orchestrator_CalculateServer) error {
	log.Println("gRPC agent connected")
	defer log.Println("gRPC agent disconnected")

	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()

	completion := make(chan struct{})
	defer close(completion)

	go func() {
		defer cancel()
		for {
			select {
			case task := <-tasksCh:
				s.mu.Lock()
				err := stream.Send(&pb.TaskRequest{
					Id:       int32(task.ID),
					Arg1:     task.Left.Value,
					Arg2:     task.Right.Value,
					Operator: task.Value,
				})
				s.mu.Unlock()

				if err != nil {
					log.Printf("Error sending task: %v", err)
					return
				}
			case <-ctx.Done():
				return
			case <-completion:
				return
			}
		}
	}()

	go func() {
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				response, err := stream.Recv()
				if err != nil {
					log.Printf("Error receiving response: %v", err)
					return
				}
				resultsCh <- models.Result{
					ID:     int(response.Id),
					Result: float64(response.Result),
					Error:  response.Error,
				}
			}
		}
	}()

	<-ctx.Done()
	return nil
}

func NewServer() *Server {
	return &Server{mu: sync.Mutex{}}
}
