package postgresql

import (
	"context"

	"github.com/VictorNevola/internal/pkg/entity/company"
	"github.com/uptrace/bun"
)

type (
	CompanyRepository struct {
		bun.DB
	}
)

func NewCompanyRepository(db *bun.DB) *CompanyRepository {
	return &CompanyRepository{
		DB: *db,
	}
}

func (r *CompanyRepository) CreateCompany(ctx context.Context, model *company.Model) (*company.Model, error) {
	_, err := r.DB.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}
