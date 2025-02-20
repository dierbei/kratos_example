package data

import (
	"context"
	"hello/app/user/ent"

	"hello/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userterRepo struct {
	data *Data
	log  *log.Helper
}

// NewuserterRepo .
func NewuserterRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userterRepo) Save(context.Context, *ent.User) (*ent.User, error) {
	return nil, nil
}

func (r *userterRepo) Update(context.Context, *ent.User) (*ent.User, error) {
	return nil, nil
}

func (r *userterRepo) FindByID(context.Context, int64) (*ent.User, error) {
	return nil, nil
}

func (r *userterRepo) ListByHello(context.Context, string) ([]*ent.User, error) {
	return nil, nil
}

func (r *userterRepo) ListAll(context.Context) ([]*ent.User, error) {
	return nil, nil
}
