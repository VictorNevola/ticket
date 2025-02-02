package userinpromotion_test

import (
	"context"
	"log"
	"time"

	userinpromotion "github.com/VictorNevola/internal/domain/userInPromotion"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	companyEntity "github.com/VictorNevola/internal/pkg/entity/company"
	promotionEntity "github.com/VictorNevola/internal/pkg/entity/promotion"
	userEntity "github.com/VictorNevola/internal/pkg/entity/user"
	userInPromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	voucherEntity "github.com/VictorNevola/internal/pkg/entity/voucher"
	customuuid "github.com/VictorNevola/internal/pkg/utils/custom-UUID"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

const (
	userInPromotionServiceKey testhelpers.ContextKey = "userInPromotionService"
)

func testContext() (context.Context, func()) {
	ctx := context.TODO()
	db, dbCleanup, _ := testhelpers.ConnectionToDB(ctx)

	//repositories
	userInPromotionRepository := postgresql.NewUserInPromotionRepository(db)
	promotionRepository := postgresql.NewPromotionRepository(db)
	voucherRepository := postgresql.NewVoucherRepo(db)

	//services
	userInPromotionService := userinpromotion.NewService(&userinpromotion.ServiceParams{
		UserInPromotionRepository: userInPromotionRepository,
		PromotionRepository:       promotionRepository,
		VoucherRepository:         voucherRepository,
	})

	ctx = context.WithValue(ctx, testhelpers.DbKey, db)
	ctx = context.WithValue(ctx, userInPromotionServiceKey, userInPromotionService)

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

func createVoucher(
	ctx context.Context,
	userID uuid.UUID,
	promotionID uuid.UUID,
	confirmedAt *time.Time,
) {
	db := ctx.Value(testhelpers.DbKey).(*bun.DB)

	VoucherUUID := uuid.New()
	randomHash := customuuid.GenerateV6()

	_, err := db.NewInsert().Model(&voucherEntity.Model{
		ID:          &VoucherUUID,
		VoucherHash: randomHash.String(),
		UserID:      &userID,
		PromotionID: &promotionID,
		ConfirmedAt: *confirmedAt,
	}).Exec(ctx)
	if err != nil {
		log.Println("Error creating initial voucher data", err)
	}
}

func getUserInPromotionByUserID(ctx context.Context, userID uuid.UUID) *userInPromotionEntity.Model {
	db := ctx.Value(testhelpers.DbKey).(*bun.DB)
	userInPromotionModel := &userInPromotionEntity.Model{}
	err := db.NewSelect().Model(userInPromotionModel).Where("user_id = ?", userID).Scan(ctx)
	if err != nil {
		return nil
	}
	return userInPromotionModel
}
