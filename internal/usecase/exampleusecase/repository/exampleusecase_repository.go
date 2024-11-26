package repository

import (
	"context"
	"go-template-clean/config"
	"go-template-clean/internal/domain"
	"go-template-clean/internal/domain/entity"
	"go-template-clean/pkg/db"
	"go-template-clean/pkg/logger"
)

type exampleRepository struct {
}

func NewExampleRepository(db *db.DB, logger logger.Logger, config *config.Configuration) domain.ExampleRepository {

	return &exampleRepository{}
}

func (r exampleRepository) FindUser(ctx context.Context, key string) (entity.User, error) {
	return entity.User{Name: "test-name"}, nil
}
