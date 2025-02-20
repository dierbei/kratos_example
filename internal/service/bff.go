package service

import (
	"context"
	v1 "hello/api/helloworld/v1"
	"hello/internal/data"
)

type BffService struct {
	v1.UnimplementedUserServer
	uc *data.Data
	//uc *biz.GreeterUsecase
}

func NewBffService(uc *data.Data) *BffService {
	return &BffService{uc: uc}
}

func (s *BffService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BffService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BffService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BffService) ListUsers(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BffService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.GetUserReply, error) {
	//TODO implement me
	panic("implement me")
}

//// SayHello implements helloworld.GreeterServer.
//func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
//	conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint("127.0.0.1:9001"))
//	if err != nil {
//		panic(err)
//	}
//	userClient := v1.NewUserClient(conn)
//	userClient.
//
//	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
//	if err != nil {
//		return nil, err
//	}
//	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
//}
