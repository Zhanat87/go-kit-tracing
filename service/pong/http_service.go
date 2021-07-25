package pong

import (
	"context"
	"fmt"

	"github.com/Zhanat87/go-kit-tracing/transport"
)

type HTTPService interface {
	HTTP(ctx context.Context, req interface{}) (response interface{}, err error)
	Grpc(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct {
	service Service
}

func NewHTTPService() HTTPService {
	return &httpService{service: NewService()}
}

func (s *httpService) HTTP(ctx context.Context, req interface{}) (interface{}, error) {
	pongRequest, ok := req.(*transport.PongRequest)
	if !ok {
		return nil, fmt.Errorf("error convert transport.pongRequest: %#v", req)
	}
	pong, err := s.service.Pong(ctx, pongRequest.Data)
	if err != nil {
		return nil, err
	}

	return &transport.PongResponse{Data: pong}, nil
}

func (s *httpService) Grpc(ctx context.Context, req interface{}) (interface{}, error) {
	return &transport.PongRequest{Data: "grpc pong response"}, nil
}
