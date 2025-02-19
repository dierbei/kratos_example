package service

import (
	"context"
	v1 "hello/api/helloworld/v1"
	"hello/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	//conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint("127.0.0.1:9001"))
	//if err != nil {
	//	panic(err)
	//}
	//userClient := v1.NewUserClient(conn)
	//userClient.GetUser()

	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
