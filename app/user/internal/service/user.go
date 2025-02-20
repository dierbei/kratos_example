package service

import (
	"context"
	"hello/app/user/ent"
	"hello/app/user/internal/biz"

	pb "hello/api/helloworld/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.UserterUsecase
}

func NewUserterService(uc *biz.UserterUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	u, err := s.uc.CreateUserter(ctx, &ent.User{
		//ID:   re,
		Age:  int(req.User.Age),
		Name: req.User.Name,
	})
	if err != nil {
		return nil, err
	}

	//fmt.Println(u)
	return &pb.CreateUserReply{
		User: &pb.UserModel{
			Name: u.Name,
			Age:  int64(u.Age),
			Id:   int64(u.ID),
		},
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	u, err := s.uc.UpdateUserter(ctx, &ent.User{
		ID:   int(req.User.Id),
		Age:  int(req.User.Age),
		Name: req.User.Name,
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserReply{
		User: &pb.UserModel{
			Name: u.Name,
			Age:  int64(u.Age),
			Id:   int64(u.ID),
		},
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, err := s.uc.FindByIDUserter(ctx, int(req.User.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		User: &pb.UserModel{
			Name: u.Name,
			Age:  int64(u.Age),
			Id:   int64(u.ID),
		},
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	users, err := s.uc.ListAllUserter(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]*pb.UserModel, 0)
	for _, u := range users {
		data = append(data, &pb.UserModel{
			Name: u.Name,
			Age:  int64(u.Age),
			Id:   int64(u.ID),
		})
	}

	return &pb.ListUserReply{
		User: data,
	}, nil
}
