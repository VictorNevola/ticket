package userinpromotion

import (
	"context"
	"database/sql"

	"github.com/VictorNevola/internal/infra/ports/repo"
	userinpromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
)

type (
	Service interface {
		BindUserToPromotion(ctx context.Context, data *userinpromotionEntity.BindUserToPromotionData) error
	}

	ServiceParams struct {
		UserInPromotionRepository repo.UserInPromotionRepo
	}

	service struct {
		userInPromotionRepository repo.UserInPromotionRepo
	}
)

func NewService(params *ServiceParams) Service {
	return &service{
		userInPromotionRepository: params.UserInPromotionRepository,
	}
}

func (s *service) BindUserToPromotion(ctx context.Context, data *userinpromotionEntity.BindUserToPromotionData) error {
	userInPromotion, err := s.userInPromotionRepository.GetUserInPromotionByUserIDAndPromotionID(ctx, data.UserID, data.PromotionID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if userInPromotion != nil {
		return ErrUserAlreadyInPromotion
	}

	return s.userInPromotionRepository.BindUserToPromotion(ctx, userinpromotionEntity.Model{
		UserID:      data.UserID,
		PromotionID: data.PromotionID,
	})
}
