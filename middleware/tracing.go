package middleware

import (
	"context"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	"github.com/Zhanat87/go-kit-tracing/service/pong"
	"github.com/openzipkin/zipkin-go"
)

type pongZipkinTracingMiddleware struct {
	next         pong.HTTPService
	zipkinTracer gokitmiddlewares.Tracer
}

func NewPongZipkinTracingMiddleware(s pong.HTTPService,
	packageName string, zipkinTracer *zipkin.Tracer) pong.HTTPService {
	return &pongZipkinTracingMiddleware{
		next:         s,
		zipkinTracer: gokitmiddlewares.NewZipkinTracing(zipkinTracer, packageName),
	}
}

func (s *pongZipkinTracingMiddleware) HTTP(ctx context.Context, req interface{}) (_ interface{}, err error) {
	span, ctx := s.zipkinTracer.Trace(ctx, "http")
	defer span.Finish()

	return s.next.HTTP(ctx, req)
}

func (s *pongZipkinTracingMiddleware) Grpc(ctx context.Context, req interface{}) (_ interface{}, err error) {
	span, ctx := s.zipkinTracer.Trace(ctx, "grpc")
	defer span.Finish()

	return s.next.Grpc(ctx, req)
}
