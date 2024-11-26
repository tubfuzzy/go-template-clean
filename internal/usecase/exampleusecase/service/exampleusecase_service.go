package service

import (
	"context"
	"go-template-clean/config"
	"go-template-clean/internal/domain"
	"go-template-clean/pkg/cache"
	"go-template-clean/pkg/logger"
)

type exampleService struct {
	domain.ExampleRepository
	*config.Configuration
}

func NewExampleService(exampleRepository domain.ExampleRepository, config *config.Configuration, cache cache.Engine, logger logger.Logger) domain.ExampleService {
	return &exampleService{
		ExampleRepository: exampleRepository,
		Configuration:     config,
	}
}

func (s exampleService) GetUserTest(ctx context.Context) (domain.User, error) {
	result, err := s.ExampleRepository.FindUser(ctx, "test")
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Name: result.Name,
	}, nil
}
