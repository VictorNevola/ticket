package userinpromotion

import (
	"context"
	"database/sql"

	"github.com/VictorNevola/internal/infra/ports/repo"
	userinpromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	time_location "github.com/VictorNevola/internal/pkg/utils/time-location"
	"github.com/google/uuid"
)

type (
	Service interface {
		BindUserToPromotion(ctx context.Context, data *userinpromotionEntity.BindUserToPromotionData) error
	}

	ServiceParams struct {
		UserInPromotionRepository repo.UserInPromotionRepo
		VoucherRepository         repo.VoucherRepo
		PromotionRepository       repo.PromotionRepo
	}

	service struct {
		userInPromotionRepository repo.UserInPromotionRepo
		voucherRepository         repo.VoucherRepo
		promotionRepository       repo.PromotionRepo
	}
)

func NewService(params *ServiceParams) Service {
	return &service{
		userInPromotionRepository: params.UserInPromotionRepository,
		voucherRepository:         params.VoucherRepository,
		promotionRepository:       params.PromotionRepository,
	}
}

func (s *service) BindUserToPromotion(ctx context.Context, data *userinpromotionEntity.BindUserToPromotionData) error {
	promotionDetails, err := s.promotionRepository.GetPromotionByID(ctx, data.PromotionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrPromotionNotFound
		}
		return err
	}

	// check if the promotion has expired
	if promotionDetails.EndDate.Before(time_location.Now()) {
		return ErrPromotionHasExpired
	}

	userInPromotion, err := s.userInPromotionRepository.GetUserInPromotionByUserIDAndPromotionID(ctx, data.UserID, data.PromotionID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if userInPromotion != nil {
		return ErrUserAlreadyInPromotion
	}

	// check if the promotion has available quantity of users
	vouchersUsage, err := s.voucherRepository.GetAllVouchersByFilters(ctx, repo.GetAllVouchersByFilters{
		PromotionID:      &data.PromotionID,
		ConfirmedAtIsNil: false,
	})
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	usersInPromotionSet := make(map[uuid.UUID]bool)
	qtyOfUniqueUsersInPromotion := 0

	for _, voucherUsage := range vouchersUsage {
		if !usersInPromotionSet[*voucherUsage.UserID] {
			usersInPromotionSet[*voucherUsage.UserID] = true
			qtyOfUniqueUsersInPromotion++
		}
	}

	if qtyOfUniqueUsersInPromotion >= promotionDetails.QtyMaxUsers {
		return ErrPromotionHasNoAvailableQuantityOfUsers
	}

	return s.userInPromotionRepository.BindUserToPromotion(ctx, userinpromotionEntity.Model{
		UserID:      data.UserID,
		PromotionID: data.PromotionID,
	})
}
