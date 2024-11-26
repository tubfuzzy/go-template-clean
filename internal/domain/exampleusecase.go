package domain

import (
	"context"
	"go-template-clean/internal/domain/entity"
)

type ExampleService interface {
	GetUserTest(ctx context.Context) (User, error)
}

type ExampleRepository interface {
	FindUser(ctx context.Context, key string) (entity.User, error)
}

type User struct {
	Name string `json:"name"`
}
