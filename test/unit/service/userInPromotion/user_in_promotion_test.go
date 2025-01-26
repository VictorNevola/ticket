package userinpromotion_test

import (
	"testing"
	"time"

	userinpromotion "github.com/VictorNevola/internal/domain/userInPromotion"
	userinpromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	time_location "github.com/VictorNevola/internal/pkg/utils/time-location"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserInPromotion(t *testing.T) {
	t.Parallel()
	ctx, dbCleanup := testContext()
	defer dbCleanup()

	t.Run("should create a user in promotion successfully", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(
			ctx,
			time_location.Now().Add(time.Hour*24), // promotion end date is 24 hours from now,
			10,                                    // promotion qty max users is 10
		)

		// before user in promotion not exists
		hasUserInPromotion := getUserInPromotionByUserID(ctx, *user.ID)
		assert.Nil(t, hasUserInPromotion)

		userInPromotionService := ctx.Value(userInPromotionServiceKey).(userinpromotion.Service)
		err := userInPromotionService.BindUserToPromotion(ctx, &userinpromotionEntity.BindUserToPromotionData{
			UserID:      *user.ID,
			PromotionID: *promotion.ID,
		})
		assert.Nil(t, err)

		// after user in promotion exists
		hasUserInPromotion = getUserInPromotionByUserID(ctx, *user.ID)
		assert.NotNil(t, hasUserInPromotion)
		assert.Equal(t, *promotion.ID, hasUserInPromotion.PromotionID)
		assert.Equal(t, *user.ID, hasUserInPromotion.UserID)
	})

	t.Run("should not create a user in promotion if user already exists", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(
			ctx,
			time_location.Now().Add(time.Hour*24), // promotion end date is 24 hours from now,
			10,                                    // promotion qty max users is 10
		)

		userInPromotionService := ctx.Value(userInPromotionServiceKey).(userinpromotion.Service)
		err := userInPromotionService.BindUserToPromotion(ctx, &userinpromotionEntity.BindUserToPromotionData{
			UserID:      *user.ID,
			PromotionID: *promotion.ID,
		})
		assert.Nil(t, err)

		err = userInPromotionService.BindUserToPromotion(ctx, &userinpromotionEntity.BindUserToPromotionData{
			UserID:      *user.ID,
			PromotionID: *promotion.ID,
		})
		assert.NotNil(t, err)
		assert.EqualValues(t, userinpromotion.ErrUserAlreadyInPromotion, err)
	})

	t.Run("should not create a user in promotion if promotion has expired", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(
			ctx,
			time_location.Now().Add(-1*time.Hour), // promotion end date is 1 hour ago
			10,                                    // promotion qty max users is 10
		)

		userInPromotionService := ctx.Value(userInPromotionServiceKey).(userinpromotion.Service)
		err := userInPromotionService.BindUserToPromotion(ctx, &userinpromotionEntity.BindUserToPromotionData{
			UserID:      *user.ID,
			PromotionID: *promotion.ID,
		})

		assert.NotNil(t, err)
		assert.EqualValues(t, userinpromotion.ErrPromotionHasExpired, err)
	})

	t.Run("should return an error if promotion does not exist", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, _ := createInitialData(
			ctx,
			time_location.Now().Add(time.Hour*24), // promotion end date is 24 hours from now
			10,                                    // promotion qty max users is 10
		)

		userInPromotionService := ctx.Value(userInPromotionServiceKey).(userinpromotion.Service)
		err := userInPromotionService.BindUserToPromotion(ctx, &userinpromotionEntity.BindUserToPromotionData{
			UserID:      *user.ID,
			PromotionID: uuid.New(),
		})

		assert.NotNil(t, err)
		assert.EqualValues(t, userinpromotion.ErrPromotionNotFound, err)
	})

	t.Run("should return an error if promotion has no available quantity of users", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(
			ctx,
			time_location.Now().Add(time.Hour*24),
			1, // promotion qty max users is 1
		)

		confirmedAt := time_location.Now()

		createVoucher(ctx, *user.ID, *promotion.ID, &confirmedAt)
		createVoucher(ctx, *user.ID, *promotion.ID, &confirmedAt)

		userInPromotionService := ctx.Value(userInPromotionServiceKey).(userinpromotion.Service)
		err := userInPromotionService.BindUserToPromotion(ctx, &userinpromotionEntity.BindUserToPromotionData{
			UserID:      *user.ID,
			PromotionID: *promotion.ID,
		})

		assert.NotNil(t, err)
		assert.EqualValues(t, userinpromotion.ErrPromotionHasNoAvailableQuantityOfUsers, err)
	})
}
