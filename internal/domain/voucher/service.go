package voucher

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/VictorNevola/internal/infra/ports/repo"
	promotionEntity "github.com/VictorNevola/internal/pkg/entity/promotion"
	userinpromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	voucherEntity "github.com/VictorNevola/internal/pkg/entity/voucher"
	"github.com/VictorNevola/internal/pkg/utils/crypto"
	time_location "github.com/VictorNevola/internal/pkg/utils/time-location"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type (
	Service interface {
		GenerateVoucher(
			ctx context.Context,
			data voucherEntity.CreateVoucherDataBody,
			userID uuid.UUID,
		) (*voucherEntity.Model, error)
	}

	ServiceParams struct {
		VoucherRepository         repo.VoucherRepo
		PromotionRepository       repo.PromotionRepo
		UserInPromotionRepository repo.UserInPromotionRepo
		SecretKey                 string
	}

	service struct {
		voucherRepository         repo.VoucherRepo
		promotionRepository       repo.PromotionRepo
		userInPromotionRepository repo.UserInPromotionRepo
		secretKey                 string
	}
)

func NewService(params ServiceParams) Service {
	return &service{
		voucherRepository:         params.VoucherRepository,
		promotionRepository:       params.PromotionRepository,
		userInPromotionRepository: params.UserInPromotionRepository,
		secretKey:                 params.SecretKey,
	}
}

func (s *service) GenerateVoucher(
	ctx context.Context,
	data voucherEntity.CreateVoucherDataBody,
	userID uuid.UUID,
) (*voucherEntity.Model, error) {
	var (
		promotionDetails *promotionEntity.Model
		userInPromotion  *userinpromotionEntity.Model
		vouchersUsage    []voucherEntity.Model
		errorGroup       errgroup.Group
	)

	errorGroup.Go(func() error {
		promotionDetailsDB, err := s.promotionRepository.GetPromotionByID(ctx, data.PromotionID)
		if err != nil {
			return err
		}
		promotionDetails = promotionDetailsDB
		return nil
	})

	errorGroup.Go(func() error {
		userInPromotionDB, err := s.userInPromotionRepository.GetUserInPromotionByUserIDAndPromotionID(
			ctx,
			userID,
			data.PromotionID,
		)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		userInPromotion = userInPromotionDB
		return nil
	})

	errorGroup.Go(func() error {
		vouchers, err := s.voucherRepository.GetAllVouchersByFilters(ctx, repo.GetAllVouchersByFilters{
			PromotionID:      &data.PromotionID,
			ConfirmedAtIsNil: false,
		})
		if err != nil {
			return err
		}
		vouchersUsage = vouchers
		return nil
	})

	if err := errorGroup.Wait(); err != nil {
		return nil, err
	}

	// check if campaing already is active
	if promotionDetails.EndDate.Before(time_location.Now()) {
		return nil, ErrPromotionIsNotActive
	}

	// check if the promotion has available quantity of vouchers
	qtyMaxOfVouchersToPromotion := promotionDetails.QtyMaxUsers * promotionDetails.VouchersPerUser
	if len(vouchersUsage) >= qtyMaxOfVouchersToPromotion {
		return nil, ErrPromotionHasNoAvailableQuantityOfVouchers
	}

	// check if the user is in the promotion
	if userInPromotion == nil {
		return nil, ErrUserNotInPromotion
	}

	return s.saveNewVoucherToUser(
		ctx,
		*promotionDetails.CompanyID,
		userID,
		data.PromotionID,
	)
}

func (s *service) saveNewVoucherToUser(
	ctx context.Context,
	companyID uuid.UUID,
	userID uuid.UUID,
	promotionID uuid.UUID,
) (*voucherEntity.Model, error) {
	voucherID := uuid.New()
	voucher := fmt.Sprintf("%s-%s", companyID.String(), voucherID.String())
	voucherHash, err := crypto.Encrypt(voucher, s.secretKey)
	if err != nil {
		return nil, err
	}

	return s.voucherRepository.CreateVoucher(ctx, &voucherEntity.Model{
		ID:          &voucherID,
		UserID:      &userID,
		PromotionID: &promotionID,
		VoucherHash: voucherHash,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(time.Hour), // expires at 1 hour
	})
}
