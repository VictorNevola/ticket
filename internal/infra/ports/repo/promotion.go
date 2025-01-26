package repo

import (
	"context"

	"github.com/VictorNevola/internal/pkg/entity/promotion"
	"github.com/google/uuid"
)

type (
	PromotionRepo interface {
		CreatePromotion(ctx context.Context, model *promotion.Model) (*promotion.Model, error)
		GetPromotionByID(ctx context.Context, id uuid.UUID) (*promotion.Model, error)
	}
)
