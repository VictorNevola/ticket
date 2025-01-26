package postgresql

import (
	"context"

	"github.com/VictorNevola/internal/infra/ports/repo"
	voucherEntity "github.com/VictorNevola/internal/pkg/entity/voucher"
	"github.com/uptrace/bun"
)

type (
	VoucherRepository struct {
		bun.DB
	}
)

func NewVoucherRepo(db *bun.DB) *VoucherRepository {
	return &VoucherRepository{
		DB: *db,
	}
}

func (r *VoucherRepository) CreateVoucher(ctx context.Context, data *voucherEntity.Model) (*voucherEntity.Model, error) {
	_, err := r.DB.NewInsert().Model(data).Exec(ctx)
	return data, err
}

func (r *VoucherRepository) GetAllVouchersByFilters(
	ctx context.Context,
	filters repo.GetAllVouchersByFilters,
) ([]voucherEntity.Model, error) {
	var vouchers []voucherEntity.Model

	query := r.DB.NewSelect().Model(&vouchers)

	if filters.PromotionID != nil {
		query.Where("promotion_id = ?", filters.PromotionID)
	}

	if filters.ConfirmedAtIsNil {
		query.Where("confirmed_at IS NULL")
	} else {
		query.Where("confirmed_at IS NOT NULL")
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return vouchers, nil
}
