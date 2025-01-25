package company

import (
	"context"

	"github.com/VictorNevola/internal/infra/ports/repo"
	"github.com/VictorNevola/internal/pkg/entity/company"
	"github.com/VictorNevola/internal/pkg/utils/crypto"
)

type (
	Service interface {
		CreateCompany(ctx context.Context, data *company.CreateData) (*company.Model, error)
	}

	ServiceParams struct {
		CompanyRepository repo.CompanyRepo
		SecretKey         string
	}

	service struct {
		companyRepository repo.CompanyRepo
		secretKey         string
	}
)

func NewService(params *ServiceParams) Service {
	return &service{
		companyRepository: params.CompanyRepository,
		secretKey:         params.SecretKey,
	}
}

func (s *service) CreateCompany(ctx context.Context, data *company.CreateData) (*company.Model, error) {
	randomKey, err := crypto.GenerateSecretKey(32)
	if err != nil {
		return nil, err
	}

	encryptedSecretKey, err := crypto.Encrypt(randomKey, s.secretKey)
	if err != nil {
		return nil, err
	}

	return s.companyRepository.CreateCompany(ctx, &company.Model{
		Name:      data.Name,
		TaxID:     data.TaxID,
		SecretKey: encryptedSecretKey,
	})
}
