package server

import (
	"bytes"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	v1 "hello/api/helloworld/v1"
	"hello/internal/conf"
	"hello/internal/service"
	"io"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// DefaultRequestDecoder decodes the request body to object.
func DefaultRequestDecoder(r *http.Request, v interface{}) error {
	//codec, ok := http.CodecForRequest(r, "Content-Type")
	//if !ok {
	//	return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
	//}
	data, err := io.ReadAll(r.Body)

	//fmt.Println(string(data))

	// reset body.
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	if err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	if len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	//if err = codec.Unmarshal(data, v); err != nil {
	//	return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
	//}
	//fmt.Println(v)
	return nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, bff *service.BffService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.RequestDecoder(DefaultRequestDecoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterBffServiceHTTPServer(srv, bff)
	return srv
}
