package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"hello/app/user/ent"
	"hello/app/user/internal/conf"

	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
//var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

var ProviderSet = wire.NewSet(NewData, NewUserterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &Data{
		db: client,
	}, cleanup, nil
}
