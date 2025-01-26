package voucher_test

import (
	"context"
	"log"
	"time"

	"github.com/VictorNevola/internal/domain/voucher"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	companyEntity "github.com/VictorNevola/internal/pkg/entity/company"
	promotionEntity "github.com/VictorNevola/internal/pkg/entity/promotion"
	userEntity "github.com/VictorNevola/internal/pkg/entity/user"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

const (
	voucherServiceKey testhelpers.ContextKey = "voucherService"
)

func testContext() (context.Context, func()) {
	ctx := context.TODO()
	db, dbCleanup, _ := testhelpers.ConnectionToDB(ctx)

	//repositories
	voucherRepository := postgresql.NewVoucherRepo(db)
	promotionRepository := postgresql.NewPromotionRepository(db)
	userInPromotionRepository := postgresql.NewUserInPromotionRepository(db)

	//services
	voucherService := voucher.NewService(voucher.ServiceParams{
		VoucherRepository:         voucherRepository,
		PromotionRepository:       promotionRepository,
		UserInPromotionRepository: userInPromotionRepository,
		SecretKey:                 "secret",
	})

	ctx = context.WithValue(ctx, testhelpers.DbKey, db)
	ctx = context.WithValue(ctx, voucherServiceKey, voucherService)

	return ctx, dbCleanup
}

func createInitialData(
	ctx context.Context,
	promotionEndDate time.Time,
	promotionQtyMaxUsers int,
) (*userEntity.Model, *promotionEntity.Model) {
	db := ctx.Value(testhelpers.DbKey).(*bun.DB)

	companyUUID := uuid.New()
	company := &companyEntity.Model{
		ID:    &companyUUID,
		Name:  "Company Test",
		TaxID: "123456789",
	}

	_, err := db.NewInsert().Model(company).Exec(ctx)
	if err != nil {
		log.Println("Error creating initial company data", err)
	}

	promotionUUID := uuid.New()
	promotion := &promotionEntity.Model{
		ID:                    &promotionUUID,
		CompanyID:             company.ID,
		Name:                  "Promotion Test",
		TextMessageInProgress: "Test",
		TextMessageSuccess:    "Test",
		StartDate:             time.Now(),
		EndDate:               promotionEndDate,
		QtyMaxUsers:           promotionQtyMaxUsers,
		VouchersPerUser:       1,
	}
	_, err = db.NewInsert().Model(promotion).Exec(ctx)
	if err != nil {
		log.Println("Error creating initial promotion data", err)
	}

	userUUID := uuid.New()
	user := &userEntity.Model{
		ID:       &userUUID,
		Email:    "test@test.com",
		Username: "test",
	}

	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		log.Println("Error creating initial user data", err)
	}

	return user, promotion
}
