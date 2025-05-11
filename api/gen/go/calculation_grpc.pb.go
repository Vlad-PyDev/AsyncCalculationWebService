package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	Orchestrator_Calculate_FullMethodName = "/calculate.Orchestrator/Calculate"
)

type OrchestratorClient interface {
	Calculate(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AgentResponse, TaskRequest], error)
}

type orchestratorClient struct {
	conn grpc.ClientConnInterface
}

func (c *orchestratorClient) Calculate(ctx_ context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AgentResponse, TaskRequest], error) {
	callOptions := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.conn.NewStream(ctx_, &Orchestrator_ServiceDesc.Streams[0], Orchestrator_Calculate_FullMethodName, callOptions...)
	if err != nil {
		return nil, err
	}
	streamWrapper := &grpc.GenericClientStream[AgentResponse, TaskRequest]{ClientStream: stream}
	return streamWrapper, nil
}

func NewOrchestratorClient(conn grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{conn}
}

type Orchestrator_CalculateClient = grpc.BidiStreamingClient[AgentResponse, TaskRequest]

type OrchestratorServer interface {
	Calculate(grpc.BidiStreamingServer[AgentResponse, TaskRequest]) error
	mustEmbedUnimplementedOrchestratorServer()
}

type UnimplementedOrchestratorServer struct{}

func (UnimplementedOrchestratorServer) Calculate(stream grpc.BidiStreamingServer[AgentResponse, TaskRequest]) error {
	return status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}

func (UnimplementedOrchestratorServer) mustEmbedUnimplementedOrchestratorServer() {}
func (UnimplementedOrchestratorServer) verifyEmbeddedByValue()                    {}

type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}

func RegisterOrchestratorServer(registrar grpc.ServiceRegistrar, service OrchestratorServer) {
	if t, ok := service.(interface{ verifyEmbeddedByValue() }); ok {
		t.verifyEmbeddedByValue()
	}
	registrar.RegisterService(&Orchestrator_ServiceDesc, service)
}

func _Orchestrator_Calculate_Handler(service interface{}, serverStream grpc.ServerStream) error {
	return service.(OrchestratorServer).Calculate(&grpc.GenericServerStream[AgentResponse, TaskRequest]{ServerStream: serverStream})
}

type Orchestrator_CalculateServer = grpc.BidiStreamingServer[AgentResponse, TaskRequest]

var Orchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculate.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Calculate",
			Handler:       _Orchestrator_Calculate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculation.proto",
}
