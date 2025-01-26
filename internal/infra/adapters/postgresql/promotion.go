package postgresql

import (
	"context"

	"github.com/VictorNevola/internal/pkg/entity/promotion"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	PromotionRepository struct {
		bun.DB
	}
)

func NewPromotionRepository(db *bun.DB) *PromotionRepository {
	return &PromotionRepository{
		DB: *db,
	}
}

func (r *PromotionRepository) CreatePromotion(ctx context.Context, model *promotion.Model) (*promotion.Model, error) {
	_, err := r.DB.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, err
}

func (r *PromotionRepository) GetPromotionByID(ctx context.Context, id uuid.UUID) (*promotion.Model, error) {
	model := promotion.Model{}

	err := r.DB.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
