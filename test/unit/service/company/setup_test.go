package company_test

import (
	"context"
	"os"

	"github.com/VictorNevola/internal/domain/company"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	companyEntity "github.com/VictorNevola/internal/pkg/entity/company"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/uptrace/bun"
)

const (
	companyServiceKey testhelpers.ContextKey = "companyService"
)

func testContext() (context.Context, func()) {
	testhelpers.LoadEnv()
	ctx := context.TODO()
	db, dbCleanup, _ := testhelpers.ConnectionToDB(ctx)

	companyRepository := postgresql.NewCompanyRepository(db)
	companyService := company.NewService(&company.ServiceParams{
		CompanyRepository: companyRepository,
		SecretKey:         os.Getenv("SecretKey"),
	})

	ctx = context.WithValue(ctx, testhelpers.DbKey, db)
	ctx = context.WithValue(ctx, companyServiceKey, companyService)

	return ctx, dbCleanup
}

func getCompany(ctx context.Context) *companyEntity.Model {
	db := ctx.Value(testhelpers.DbKey).(*bun.DB)
	company := &companyEntity.Model{}

	err := db.NewSelect().Model(company).Scan(ctx)
	if err != nil {
		return nil
	}

	return company
}
