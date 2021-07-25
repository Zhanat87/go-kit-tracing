package middleware

import (
	"context"
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	"github.com/Zhanat87/go-kit-tracing/service/pong"
	"github.com/go-kit/kit/metrics"
)

type pongInstrumentingMiddleware struct {
	next  pong.HTTPService
	saver gokitmiddlewares.Saver
}

func NewPongInstrumentingMiddleware(s pong.HTTPService, packageName string,
	counter metrics.Counter, latency metrics.Histogram, counterE metrics.Counter) pong.HTTPService {
	return &pongInstrumentingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
	}
}

func (s *pongInstrumentingMiddleware) HTTP(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "http")
	}(time.Now())

	return s.next.HTTP(ctx, req)
}

func (s *pongInstrumentingMiddleware) Grpc(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "grpc")
	}(time.Now())

	return s.next.Grpc(ctx, req)
}
