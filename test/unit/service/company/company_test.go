package company_test

import (
	"testing"

	"github.com/VictorNevola/internal/domain/company"
	companyEntity "github.com/VictorNevola/internal/pkg/entity/company"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
	t.Parallel()
	ctx, dbCleanup := testContext()
	defer dbCleanup()

	t.Run("should create a company successfully", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)
		// before company not exists
		hasCompany := getCompany(ctx)
		assert.Nil(t, hasCompany)

		companyService := ctx.Value(companyServiceKey).(company.Service)
		companyCreated, err := companyService.CreateCompany(ctx, &companyEntity.CreateData{
			Name:  "Company Test",
			TaxID: "123456789",
		})

		assert.Nil(t, err)
		assert.NotNil(t, companyCreated)
		assert.Equal(t, "Company Test", companyCreated.Name)
		assert.Equal(t, "123456789", companyCreated.TaxID)
		assert.NotEmpty(t, companyCreated.SecretKey)
	})

	t.Run("should not create a company with the same tax id", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)
		companyService := ctx.Value(companyServiceKey).(company.Service)
		companyService.CreateCompany(ctx, &companyEntity.CreateData{
			Name:  "Company Test",
			TaxID: "123456789",
		})

		companyCreated, err := companyService.CreateCompany(ctx, &companyEntity.CreateData{
			Name:  "Company Test 2",
			TaxID: "123456789",
		})
		assert.NotNil(t, err)
		assert.Nil(t, companyCreated)
	})
}
