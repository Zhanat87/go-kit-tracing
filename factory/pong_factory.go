package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-tracing/middleware"
	"github.com/Zhanat87/go-kit-tracing/service/pong"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
)

type PongServiceFactory struct{}

func (s *PongServiceFactory) CreateHTTPService(packageName string, logger log.Logger, zipkinTracer *zipkin.Tracer) pong.HTTPService {
	srv := pong.NewHTTPService()
	srv = middleware.NewPongLoggingMiddleware(srv, packageName, log.With(logger, "component", packageName))
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(packageName)
	srv = middleware.NewPongInstrumentingMiddleware(srv, packageName, counter, duration, counterError)
	srv = middleware.NewPongZipkinTracingMiddleware(srv, packageName, zipkinTracer)

	return srv
}
