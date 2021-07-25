package pong

import (
	"context"
	"fmt"

	"github.com/openzipkin/zipkin-go"

	"github.com/Zhanat87/go-kit-tracing/transport"
)

type HTTPService interface {
	HTTP(ctx context.Context, req interface{}) (response interface{}, err error)
	Grpc(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct {
	service      Service
	zipkinTracer *zipkin.Tracer
}

func NewHTTPService(zipkinTracer *zipkin.Tracer) HTTPService {
	return &httpService{service: NewService(zipkinTracer), zipkinTracer: zipkinTracer}
}

func (s *httpService) HTTP(ctx context.Context, req interface{}) (interface{}, error) {
	span, ctx := s.zipkinTracer.StartSpanFromContext(ctx, "pong http service response")
	defer span.Finish()
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
