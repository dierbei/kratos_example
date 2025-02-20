package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	//grpc2 "google.golang.org/grpc"
	v1 "hello/api/helloworld/v1"
	"hello/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is data providers.
// var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	// TODO wrapped database client
	UserClient v1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.
			WithEndpoint("127.0.0.1:9001"),
	)
	if err != nil {
		panic(err)
	}

	//conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint("127.0.0.1:9001"), u=Withtrans)
	//if err != nil {
	//	panic(err)
	//}
	userClient := v1.NewUserClient(conn)

	return &Data{
		UserClient: userClient,
	}, cleanup, nil
}
