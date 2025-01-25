package postgresql

import (
	"context"

	userinpromotion "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	UserInPromotionRepository struct {
		bun.DB
	}
)

func NewUserInPromotionRepository(db *bun.DB) *UserInPromotionRepository {
	return &UserInPromotionRepository{
		DB: *db,
	}
}

func (r *UserInPromotionRepository) BindUserToPromotion(ctx context.Context, model userinpromotion.Model) error {
	_, err := r.DB.NewInsert().Model(&model).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserInPromotionRepository) GetUserInPromotionByUserIDAndPromotionID(
	ctx context.Context,
	userID, promotionID uuid.UUID,
) (*userinpromotion.Model, error) {
	model := userinpromotion.Model{}

	err := r.DB.NewSelect().Model(&model).
		Where("user_id = ?", userID).
		Where("promotion_id = ?", promotionID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
