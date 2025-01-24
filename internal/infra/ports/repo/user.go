package repo

import (
	"context"

	"github.com/VictorNevola/internal/pkg/entity/user"
)

type (
	UserRepo interface {
		CreateUser(ctx context.Context, model *user.Model) (*user.Model, error)
	}
)
