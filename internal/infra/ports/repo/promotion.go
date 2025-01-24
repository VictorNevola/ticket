package repo

import (
	"context"

	"github.com/VictorNevola/internal/pkg/entity/promotion"
)

type (
	PromotionRepo interface {
		CreatePromotion(ctx context.Context, model *promotion.Model) (*promotion.Model, error)
	}
)
