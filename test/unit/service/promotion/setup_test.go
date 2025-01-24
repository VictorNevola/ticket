package promotion_test

import (
	"context"
	"log"

	"github.com/VictorNevola/internal/domain/promotion"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	companyEntity "github.com/VictorNevola/internal/pkg/entity/company"
	promotionEntity "github.com/VictorNevola/internal/pkg/entity/promotion"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func testContext() (context.Context, func()) {
	ctx := context.TODO()
	db, dbCleanup, _ := testhelpers.ConnectionToDB(ctx)

	promotionRepository := postgresql.NewPromotionRepository(db)
	promotionService := promotion.NewService(&promotion.ServiceParams{
		PromotionRepository: promotionRepository,
	})

	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "promotionService", promotionService)

	return ctx, dbCleanup
}

func clearDatabase(ctx context.Context) {
	db := ctx.Value("db").(*bun.DB)
	db.NewDelete().Model(&companyEntity.Model{}).Where("1=1").Exec(ctx)
	db.NewDelete().Model(&promotionEntity.Model{}).Where("1=1").Exec(ctx)

	log.Println("Database cleared")
}

func createInitialData(ctx context.Context) *companyEntity.Model {
	db := ctx.Value("db").(*bun.DB)

	companyUUID := uuid.New()
	company := &companyEntity.Model{
		ID:    &companyUUID,
		Name:  "Company Test",
		TaxID: "123456789",
	}

	_, err := db.NewInsert().Model(company).Exec(ctx)
	if err != nil {
		log.Println("Error creating initial data")
	}

	return company
}

func getPromotion(ctx context.Context) *promotionEntity.Model {
	db := ctx.Value("db").(*bun.DB)
	promotionModel := &promotionEntity.Model{}

	err := db.NewSelect().Model(promotionModel).Where("1=1").Scan(ctx)
	if err != nil {
		return nil
	}

	return promotionModel
}
