package user

import (
	"context"

	"github.com/VictorNevola/internal/infra/ports/repo"
	"github.com/VictorNevola/internal/pkg/entity/user"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service interface {
		CreateUser(ctx context.Context, data *user.CreateData) (*user.Model, error)
	}

	ServiceParams struct {
		UserRepository repo.UserRepo
	}

	service struct {
		userRepository repo.UserRepo
	}
)

func NewService(params *ServiceParams) Service {
	return &service{
		userRepository: params.UserRepository,
	}
}

func (s *service) CreateUser(
	ctx context.Context,
	data *user.CreateData,
) (*user.Model, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password),
		bcrypt.MinCost,
	)
	if err != nil {
		return nil, err
	}

	return s.userRepository.CreateUser(ctx, &user.Model{
		Username: data.Username,
		Email:    data.Email,
		Password: string(hashedPassword),
	})
}
