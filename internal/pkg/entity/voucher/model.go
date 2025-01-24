package voucher

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Model struct {
		bun.BaseModel `bun:"vouchers,alias:vouchers"`
		ID            *uuid.UUID `bun:"id,pk"`
		VauncherHash  string     `bun:"vauncher_hash,unique,notnull"`
		CreatedAt     time.Time  `bun:"created_at"`
		ConfirmedAt   time.Time  `bun:"confirmed_at"`
		DeletedAt     time.Time  `bun:"deleted_at,soft_delete,nullzero"`
	}

	ModelVouncherUsage struct {
		bun.BaseModel `bun:"voucher_usages,alias:voucher_usages"`

		ID          *uuid.UUID `bun:"id,pk"`
		PromotionID *uuid.UUID `bun:"promotion_id,notnull"`
		VoucherID   uuid.UUID  `bun:"voucher_id,notnull"`
		UserID      uuid.UUID  `bun:"user_id,notnull"`
	}
)
