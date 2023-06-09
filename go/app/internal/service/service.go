package service

import (
	"context"
	"microservices/app/internal/repository"
	"microservices/app/libraries/grpc_client"
	"microservices/app/libraries/logging"
)

type AdviseService interface {
	GetAdvise(ctx context.Context, req *grpc_client.GetAdviceRequest) ([]string, []string, error)
}

type Service struct {
	AdviseService
}

func NewService(repos *repository.Repository, logger *logging.Logger) *Service {
	return &Service{
		AdviseService: NewAdviser(logger),
	}
}
