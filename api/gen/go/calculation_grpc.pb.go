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
	cc grpc.ClientConnInterface
}

func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) Calculate(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AgentResponse, TaskRequest], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Orchestrator_ServiceDesc.Streams[0], Orchestrator_Calculate_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AgentResponse, TaskRequest]{ClientStream: stream}
	return x, nil
}

type Orchestrator_CalculateClient = grpc.BidiStreamingClient[AgentResponse, TaskRequest]

type OrchestratorServer interface {
	Calculate(grpc.BidiStreamingServer[AgentResponse, TaskRequest]) error
	mustEmbedUnimplementedOrchestratorServer()
}

type UnimplementedOrchestratorServer struct{}

func (UnimplementedOrchestratorServer) Calculate(grpc.BidiStreamingServer[AgentResponse, TaskRequest]) error {
	return status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedOrchestratorServer) mustEmbedUnimplementedOrchestratorServer() {}
func (UnimplementedOrchestratorServer) testEmbeddedByValue()                      {}

type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}

func RegisterOrchestratorServer(s grpc.ServiceRegistrar, srv OrchestratorServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Orchestrator_ServiceDesc, srv)
}

func _Orchestrator_Calculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrchestratorServer).Calculate(&grpc.GenericServerStream[AgentResponse, TaskRequest]{ServerStream: stream})
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
