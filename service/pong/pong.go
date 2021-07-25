package pong

import (
	"context"
	"os"
	"time"

	"github.com/openzipkin/zipkin-go"

	"github.com/Zhanat87/common-libs/utils"
)

const (
	PackageName = "pong"
	BaseURL     = "/api/v1/pong/"
)

type Service interface {
	Pong(ctx context.Context, ping string) (string, error)
}

type service struct {
	zipkinTracer *zipkin.Tracer
}

func NewService(zipkinTracer *zipkin.Tracer) Service {
	return &service{zipkinTracer: zipkinTracer}
}

func (s *service) Pong(ctx context.Context, ping string) (string, error) {
	span, _ := s.zipkinTracer.StartSpanFromContext(ctx, "pong service response on ping request: "+ping)
	defer span.Finish()
	time.Sleep(22 * time.Millisecond)
	dateTime, err := utils.GetCurrentDateTime(os.Getenv("TIME_ZONE"))
	if err != nil {
		return "", err
	}

	return "current date and time in pong service: " + dateTime + ", req data: " + ping, nil
}
