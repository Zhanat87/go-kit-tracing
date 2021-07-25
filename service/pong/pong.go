package pong

import (
	"context"
	"os"

	"github.com/Zhanat87/common-libs/utils"
)

const (
	PackageName = "pong"
	BaseURL     = "/api/v1/pong/"
)

type Service interface {
	Pong(ctx context.Context, ping string) (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Pong(ctx context.Context, ping string) (string, error) {
	dateTime, err := utils.GetCurrentDateTime(os.Getenv("TIME_ZONE"))
	if err != nil {
		return "", err
	}

	return dateTime + ", req data: " + ping, nil
}
