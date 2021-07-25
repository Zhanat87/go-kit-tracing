package middleware

import (
	"context"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	"github.com/Zhanat87/common-libs/tracers"
	"github.com/Zhanat87/go-kit-tracing/service/pong"
	"github.com/Zhanat87/go-kit-tracing/transport"
	"github.com/go-kit/kit/endpoint"
)

type PongEndpoints struct {
	HTTPEndpoint endpoint.Endpoint
	GrpcEndpoint endpoint.Endpoint
}

func MakePongEndpoints(s pong.HTTPService) PongEndpoints {
	return PongEndpoints{
		HTTPEndpoint: gokitmiddlewares.GetDefaultEndpoint(MakePongHTTPEndpoint(s), "pong http", tracers.ZipkinTracer),
		GrpcEndpoint: gokitmiddlewares.GetDefaultEndpoint(MakePongGrpcEndpoint(s), "pong grpc", tracers.ZipkinTracer),
	}
}

func MakePongHTTPEndpoint(next pong.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.PongRequest)
		resp, err := next.HTTP(ctx, &req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func MakePongGrpcEndpoint(next pong.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.PongRequest)
		resp, err := next.Grpc(ctx, &req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
