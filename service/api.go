package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
}

// Service interface
type Service interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
	Subtract(ctx context.Context, numA, numB float32) (float32, error)
	Multiply(ctx context.Context, numA, numB float32) (float32, error)
	Divide(ctx context.Context, numA, numB float32) (float32, error)
}

// NewService func initializes a service
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}

func (s service) Subtract(ctx context.Context, numA, numB float32) (float32, error) {
	return numA - numB, nil
}

func (s service) Multiply(ctx context.Context, numA, numB float32) (float32, error) {
	return numA * numB, nil
}

func (s service) Divide(ctx context.Context, numA, numB float32) (float32, error) {
	return numA / numB, nil
}
