package company

import (
	"context"

	"github.com/VictorNevola/internal/infra/ports/repo"
	"github.com/VictorNevola/internal/pkg/entity/company"
)

type (
	Service interface {
		CreateCompany(ctx context.Context, data *company.CreateData) (*company.Model, error)
	}

	ServiceParams struct {
		CompanyRepository repo.CompanyRepo
	}

	service struct {
		companyRepository repo.CompanyRepo
	}
)

func NewService(params *ServiceParams) Service {
	return &service{
		companyRepository: params.CompanyRepository,
	}
}

func (s *service) CreateCompany(ctx context.Context, data *company.CreateData) (*company.Model, error) {
	return s.companyRepository.CreateCompany(ctx, &company.Model{
		Name:  data.Name,
		TaxID: data.TaxID,
	})
}
