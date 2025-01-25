package userinpromotion_test

import (
	"testing"

	userinpromotion "github.com/VictorNevola/internal/domain/userInPromotion"
	userinpromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserInPromotion(t *testing.T) {
	t.Parallel()
	ctx, dbCleanup := testContext()
	defer dbCleanup()

	t.Run("should create a user in promotion successfully", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		user, promotion := createInitialData(ctx)

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

		user, promotion := createInitialData(ctx)

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
}
