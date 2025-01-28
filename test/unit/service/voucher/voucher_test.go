package voucher_test

import (
	"testing"
	"time"

	"github.com/VictorNevola/internal/domain/voucher"
	voucherEntity "github.com/VictorNevola/internal/pkg/entity/voucher"
	time_location "github.com/VictorNevola/internal/pkg/utils/time-location"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGenerateVoucher(t *testing.T) {
	t.Parallel()
	ctx, dbCleanup := testContext()
	defer dbCleanup()

	t.Run("should generate a voucher successfully", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(
			ctx,
			time_location.Now().Add(time.Hour*24), // promotion end date is 24 hours from now,
			10,                                    // promotion qty max users is 10
		)
		createUserInPromotion(ctx, user, promotion)

		voucherService := ctx.Value(voucherServiceKey).(voucher.Service)
		createdVoucher, err := voucherService.GenerateVoucher(ctx, voucherEntity.CreateVoucherDataBody{
			PromotionID: *promotion.ID,
		}, *user.ID)

		assert.Nil(t, err)
		assert.NotNil(t, createdVoucher)
		assert.Equal(t, *promotion.ID, *createdVoucher.PromotionID)
		assert.Equal(t, *user.ID, *createdVoucher.UserID)
		assert.NotEmpty(t, createdVoucher.VoucherHash)
		assert.Len(t, createdVoucher.VoucherHash, 120)
		assert.Equal(t,
			createdVoucher.ExpiresAt.Format("2006-01-02 15:04:05"),
			time_location.Now().Add(time.Hour).Format("2006-01-02 15:04:05"),
		)
	})

	t.Run("should return error if the promotion is not active", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(ctx, time_location.Now().Add(-time.Hour*24), 10)
		createUserInPromotion(ctx, user, promotion)
		voucherService := ctx.Value(voucherServiceKey).(voucher.Service)
		_, err := voucherService.GenerateVoucher(ctx, voucherEntity.CreateVoucherDataBody{
			PromotionID: *promotion.ID,
		}, *user.ID)

		assert.NotNil(t, err)
		assert.EqualValues(t, err, voucher.ErrPromotionIsNotActive)
	})

	t.Run("should return error if the promotion has no available quantity of vouchers", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(ctx, time_location.Now().Add(time.Hour*24), 1)
		createUserInPromotion(ctx, user, promotion)
		createVoucher(ctx, promotion, user) // create 1 voucher for the promotion, so the promotion has no available quantity of vouchers

		voucherService := ctx.Value(voucherServiceKey).(voucher.Service)
		_, err := voucherService.GenerateVoucher(ctx, voucherEntity.CreateVoucherDataBody{
			PromotionID: *promotion.ID,
		}, *user.ID)

		assert.NotNil(t, err)
		assert.EqualValues(t, err, voucher.ErrPromotionHasNoAvailableQuantityOfVouchers)
	})

	t.Run("should return error if the user is not in the promotion", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(ctx, time_location.Now().Add(time.Hour*24), 10)

		voucherService := ctx.Value(voucherServiceKey).(voucher.Service)
		_, err := voucherService.GenerateVoucher(ctx, voucherEntity.CreateVoucherDataBody{
			PromotionID: *promotion.ID,
		}, *user.ID)

		assert.NotNil(t, err)
		assert.EqualValues(t, err, voucher.ErrUserNotInPromotion)
	})
}
