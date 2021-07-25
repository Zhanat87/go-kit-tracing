package middleware

import (
	"context"
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	"github.com/Zhanat87/go-kit-tracing/service/pong"
	"github.com/go-kit/kit/log"
)

type pongLoggingMiddleware struct {
	next  pong.HTTPService
	saver gokitmiddlewares.Saver
}

func NewPongLoggingMiddleware(s pong.HTTPService, packageName string, logger log.Logger) pong.HTTPService {
	return &pongLoggingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewLogging(logger, packageName),
	}
}

func (s *pongLoggingMiddleware) HTTP(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "http")
	}(time.Now())

	return s.next.HTTP(ctx, req)
}

func (s *pongLoggingMiddleware) Grpc(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "grpc")
	}(time.Now())

	return s.next.Grpc(ctx, req)
}
