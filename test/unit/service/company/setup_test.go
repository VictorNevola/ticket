package company_test

import (
	"context"
	"log"

	"github.com/VictorNevola/internal/domain/company"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	companyEntity "github.com/VictorNevola/internal/pkg/entity/company"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/uptrace/bun"
)

func testContext() (context.Context, func()) {
	ctx := context.TODO()
	db, dbCleanup, _ := testhelpers.ConnectionToDB(ctx)

	companyRepository := postgresql.NewCompanyRepository(db)
	companyService := company.NewService(&company.ServiceParams{
		CompanyRepository: companyRepository,
	})

	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "companyService", companyService)

	return ctx, dbCleanup
}

func getCompany(ctx context.Context) *companyEntity.Model {
	db := ctx.Value("db").(*bun.DB)
	company := &companyEntity.Model{}

	err := db.NewSelect().Model(company).Scan(ctx)
	if err != nil {
		return nil
	}

	return company
}

func clearDatabase(ctx context.Context) {
	db := ctx.Value("db").(*bun.DB)
	db.NewDelete().Model(&companyEntity.Model{}).Where("1=1").Exec(ctx)

	log.Println("Database cleared")
}
