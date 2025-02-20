package service

import (
	"context"
	v1 "hello/api/helloworld/v1"
	"hello/internal/data"
	"log"
)

type BffService struct {
	v1.UnsafeBffServiceServer
	uc *data.Data
	//uc *biz.GreeterUsecase
}

func (s *BffService) GetUser(ctx context.Context, in *v1.GetUserReq) (*v1.GetUserRes, error) {
	log.Println("GetUser", in)
	u, err := s.uc.UserClient.GetUser(ctx, &v1.GetUserRequest{
		User: &v1.UserModel{
			Id: in.GetId(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetUserRes{
		User: &v1.UserInfo{
			Id:   u.User.Id,
			Name: u.User.Name,
			Age:  u.User.Age,
		},
	}, nil
}

func NewBffService(uc *data.Data) *BffService {
	return &BffService{uc: uc}
}

// 创建用户 (POST)
func (s *BffService) CreateUser(ctx context.Context, in *v1.CreateUserReq) (*v1.GetUserRes, error) {
	log.Println("CreateUser", in)
	u, err := s.uc.UserClient.CreateUser(ctx, &v1.CreateUserRequest{
		User: &v1.UserModel{
			Name: in.Name,
			Age:  in.Age,
		},
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetUserRes{
		User: &v1.UserInfo{
			Id:   u.User.Id,
			Name: u.User.Name,
			Age:  u.User.Age,
		},
	}, nil
}

// 更新用户信息 (PUT)
func (s *BffService) UpdateUser(ctx context.Context, in *v1.UpdateUserReq) (*v1.GetUserRes, error) {
	log.Println("UpdateUser", in)

	u, err := s.uc.UserClient.UpdateUser(ctx, &v1.UpdateUserRequest{
		User: &v1.UserModel{
			Id:   in.Id,
			Name: in.Name,
			Age:  in.Age,
		},
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetUserRes{
		User: &v1.UserInfo{
			Id:   u.User.Id,
			Name: u.User.Name,
			Age:  u.User.Age,
		},
	}, nil
}

// 删除用户 (DELETE)
func (s *BffService) DeleteUser(ctx context.Context, in *v1.DelUserReq) (*v1.DelUserReply, error) {
	//_, err := s.uc.UserClient.DeleteUser(ctx, &v1.DeleteUserReq{Id: in.Id})
	//if err != nil {
	//	return nil, err
	//}

	return &v1.DelUserReply{
		Message: "User deleted successfully",
	}, nil
}

// 获取用户列表 (GET)
func (s *BffService) ListUsers(ctx context.Context, in *v1.ListUserReq) (*v1.ListUserRes, error) {
	log.Println("ListUsers", in)

	res, err := s.uc.UserClient.ListUser(ctx, &v1.ListUserRequest{})
	if err != nil {
		return nil, err
	}

	users := make([]*v1.UserInfo, len(res.User))
	for i, u := range res.User {
		users[i] = &v1.UserInfo{
			Id:   u.Id,
			Name: u.Name,
			Age:  u.Age,
		}
	}

	return &v1.ListUserRes{
		Users: users,
	}, nil
}
