package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go-template-clean/config"
	"go-template-clean/internal/domain"
	"go-template-clean/internal/domain/entity"
	"testing"
)

type ExampleRepositorySuccess struct {
	domain.ExampleRepository
}
type ExampleRepositoryFail struct {
	domain.ExampleRepository
}

func (r ExampleRepositorySuccess) FindUser(ctx context.Context, key string) (entity.User, error) {
	return entity.User{Name: "test-name"}, nil
}

func (r ExampleRepositoryFail) FindUser(ctx context.Context, key string) (entity.User, error) {
	return entity.User{}, errors.New("test error")
}

func TestExampleService_GetUserTest(t *testing.T) {
	// Prepare configurations
	mockConfig := &config.Configuration{}

	// Create test cases
	tests := []struct {
		name          string
		repository    domain.ExampleRepository
		expectedUser  domain.User
		expectedError error
	}{
		{
			name:          "Success - User Found",
			repository:    ExampleRepositorySuccess{},
			expectedUser:  domain.User{Name: "test-name"},
			expectedError: nil,
		},
		{
			name:          "Error - Repository Failure",
			repository:    ExampleRepositoryFail{},
			expectedUser:  domain.User{},
			expectedError: errors.New("test error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exampleService := NewExampleService(tt.repository, mockConfig, nil, nil)
			result, err := exampleService.GetUserTest(context.Background())

			assert.Equal(t, tt.expectedUser, result)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
