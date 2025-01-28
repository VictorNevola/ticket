package voucher_test

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/VictorNevola/internal/domain/voucher"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	companyEntity "github.com/VictorNevola/internal/pkg/entity/company"
	promotionEntity "github.com/VictorNevola/internal/pkg/entity/promotion"
	userEntity "github.com/VictorNevola/internal/pkg/entity/user"
	userinpromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	voucherEntity "github.com/VictorNevola/internal/pkg/entity/voucher"
	time_location "github.com/VictorNevola/internal/pkg/utils/time-location"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

const (
	voucherServiceKey testhelpers.ContextKey = "voucherService"
)

func testContext() (context.Context, func()) {
	testhelpers.LoadEnv()
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
		SecretKey:                 os.Getenv("SecretKey"),
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

func createUserInPromotion(ctx context.Context, user *userEntity.Model, promotion *promotionEntity.Model) *userinpromotionEntity.Model {
	db := ctx.Value(testhelpers.DbKey).(*bun.DB)

	userInPromotion := &userinpromotionEntity.Model{
		UserID:      *user.ID,
		PromotionID: *promotion.ID,
	}
	_, err := db.NewInsert().Model(userInPromotion).Exec(ctx)
	if err != nil {
		log.Println("Error creating initial user in promotion data", err)
	}

	return userInPromotion
}

func createVoucher(ctx context.Context, promotion *promotionEntity.Model, user *userEntity.Model) *voucherEntity.Model {
	db := ctx.Value(testhelpers.DbKey).(*bun.DB)
	randomHash := uuid.New().String()
	randomID := uuid.New()

	voucher := &voucherEntity.Model{
		ID:          &randomID,
		VoucherHash: randomHash,
		UserID:      user.ID,
		PromotionID: promotion.ID,
		ExpiresAt:   time_location.Now().Add(time.Hour),
		ConfirmedAt: time_location.Now(),
	}

	_, err := db.NewInsert().Model(voucher).Exec(ctx)
	if err != nil {
		log.Println("Error creating initial voucher data", err)
	}

	return voucher
}
