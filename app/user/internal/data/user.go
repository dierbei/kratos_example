package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"hello/app/user/ent"
	"hello/app/user/ent/user"
	"hello/app/user/internal/biz"
)

type userterRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserterRepo .
func NewUserterRepo(data *Data, logger log.Logger) biz.UserterRepo {
	return &userterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

//func NewUserterRepo(data *data2.Data, logger log.Logger) biz.UserRepo {
//	return &userterRepo{
//		data: data,
//		log:  log.NewHelper(logger),
//	}
//}

func (r *userterRepo) Save(ctx context.Context, req *ent.User) (*ent.User, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	//// Run the auto migration tool.
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	u, err := client.User.Create().SetAge(req.Age).SetName(req.Name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	fmt.Println("user was created: ", u)

	return nil, nil
}

func (r *userterRepo) Update(ctx context.Context, req *ent.User) (*ent.User, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	user, err := client.User.
		UpdateOneID(req.ID).
		SetName(req.Name).
		SetAge(req.Age).Save(ctx)

	return user, nil
}

func (r *userterRepo) FindByID(ctx context.Context, id int64) (*ent.User, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	user, err := client.User.Get(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("failed FindByID user: %v", err)
	}

	return user, nil
}

func (r *userterRepo) ListByHello(ctx context.Context, name string) ([]*ent.User, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	users, err := client.User.
		Query().
		Where(
			user.NameContains(name),
		).
		Limit(10).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed ListByHello user: %v", err)
	}

	return users, nil
}

func (r *userterRepo) ListAll(ctx context.Context) ([]*ent.User, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	users, err := client.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed ListByHello user: %v", err)
	}

	return users, nil
}
