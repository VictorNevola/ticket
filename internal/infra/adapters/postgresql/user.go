package postgresql

import (
	"context"

	"github.com/VictorNevola/internal/pkg/entity/user"
	"github.com/uptrace/bun"
)

type (
	UserRepository struct {
		bun.DB
	}
)

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{
		DB: *db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, model *user.Model) (*user.Model, error) {
	_, err := r.DB.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, err
}
