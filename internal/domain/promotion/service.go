package promotion

import (
	"context"

	"github.com/VictorNevola/internal/infra/ports/repo"
	"github.com/VictorNevola/internal/pkg/entity/promotion"
	time_location "github.com/VictorNevola/internal/pkg/utils/time-location"
)

type (
	Service interface {
		CreatePromotion(ctx context.Context, data *promotion.CreateData) (*promotion.Model, error)
	}

	ServiceParams struct {
		PromotionRepository repo.PromotionRepo
	}

	service struct {
		promotionRepository repo.PromotionRepo
	}
)

func NewService(params *ServiceParams) Service {
	return &service{
		promotionRepository: params.PromotionRepository,
	}
}

func (s *service) CreatePromotion(
	ctx context.Context,
	data *promotion.CreateData,
) (*promotion.Model, error) {
	startDate, _ := time_location.StringToTime(data.StartDate)
	endDate, _ := time_location.StringToTime(data.EndDate)

	return s.promotionRepository.CreatePromotion(ctx, &promotion.Model{
		CompanyID:             &data.CompanyID,
		Name:                  data.Name,
		TextMessageInProgress: data.TextMessageInProgress,
		TextMessageSuccess:    data.TextMessageSuccess,
		StartDate:             startDate,
		EndDate:               endDate,
		QtyMaxUsers:           data.QtyMaxUsers,
		VouchersPerUser:       data.VouchersPerUser,
	})
}
