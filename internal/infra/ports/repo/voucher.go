package repo

import (
	"context"

	voucherEntity "github.com/VictorNevola/internal/pkg/entity/voucher"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	VoucherRepo interface {
		bun.IDB
		CreateVoucher(ctx context.Context, data *voucherEntity.Model) (*voucherEntity.Model, error)
		GetAllVouchersByFilters(
			ctx context.Context,
			filters GetAllVouchersByFilters,
		) ([]voucherEntity.Model, error)
	}

	GetAllVouchersByFilters struct {
		PromotionID      *uuid.UUID
		ConfirmedAtIsNil bool
	}
)
