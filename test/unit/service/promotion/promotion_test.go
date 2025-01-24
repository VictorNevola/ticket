package promotion_test

import (
	"testing"

	"github.com/VictorNevola/internal/domain/promotion"
	promotionEntity "github.com/VictorNevola/internal/pkg/entity/promotion"
	"github.com/stretchr/testify/assert"
)

func TestCreatePromotion(t *testing.T) {
	t.Parallel()
	ctx, dbCleanup := testContext()
	defer dbCleanup()

	t.Run("should create a promotion successfully", func(t *testing.T) {
		defer clearDatabase(ctx)

		company := createInitialData(ctx)
		// before promotion not exists
		hasPromotion := getPromotion(ctx)
		assert.Nil(t, hasPromotion)

		promotionService := ctx.Value("promotionService").(promotion.Service)
		promotionCreated, err := promotionService.CreatePromotion(ctx, &promotionEntity.CreateData{
			CompanyID:             *company.ID,
			Name:                  "Promotion Test",
			TextMessageInProgress: "Promotion in progress",
			TextMessageSuccess:    "Promotion success",
			StartDate:             "2021-01-01",
			EndDate:               "2021-01-02",
			QtyMaxUsers:           100,
			VouchersPerUser:       10,
		})

		// after promotion exists
		assert.Nil(t, err)
		assert.NotNil(t, promotionCreated)
		assert.Equal(t, company.ID, promotionCreated.CompanyID)
		assert.Equal(t, "Promotion Test", promotionCreated.Name)
		assert.Equal(t, "Promotion in progress", promotionCreated.TextMessageInProgress)
		assert.Equal(t, "Promotion success", promotionCreated.TextMessageSuccess)
		assert.Equal(t, 100, promotionCreated.QtyMaxUsers)
		assert.Equal(t, 10, promotionCreated.VouchersPerUser)
	})
}
