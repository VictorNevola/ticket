package repo

import (
	"context"

	"github.com/VictorNevola/internal/pkg/entity/company"
)

type (
	CompanyRepo interface {
		CreateCompany(ctx context.Context, model *company.Model) (*company.Model, error)
	}
)
