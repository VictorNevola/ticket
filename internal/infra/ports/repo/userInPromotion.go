package repo

import (
	"context"

	userinpromotionEntity "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	"github.com/google/uuid"
)

type UserInPromotionRepo interface {
	BindUserToPromotion(ctx context.Context, model userinpromotionEntity.Model) error
	GetUserInPromotionByUserIDAndPromotionID(ctx context.Context, userID, promotionID uuid.UUID) (*userinpromotionEntity.Model, error)
}
