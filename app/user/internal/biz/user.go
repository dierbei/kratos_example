package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hello/app/user/ent"
	//"os/user"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

type Userter struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

type UserterRepo interface {
	Save(context.Context, *ent.User) (*ent.User, error)
	Update(context.Context, *ent.User) (*ent.User, error)
	FindByID(context.Context, int64) (*ent.User, error)
	ListByHello(context.Context, string) ([]*ent.User, error)
	ListAll(context.Context) ([]*ent.User, error)
}

type UserterUsecase struct {
	repo UserterRepo
	log  *log.Helper
}

func NewUserterUsecase(repo UserterRepo, logger log.Logger) *UserterUsecase {
	return &UserterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserterUsecase) CreateUserter(ctx context.Context, u *ent.User) (*ent.User, error) {
	uc.log.WithContext(ctx).Infof("CreateUserter: %v", u.Name)
	return uc.repo.Save(ctx, u)
}
