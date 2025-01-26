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
		VoucherHash   string     `bun:"voucher_hash,unique,notnull"`
		UserID        *uuid.UUID `bun:"user_id,notnull"`
		PromotionID   *uuid.UUID `bun:"promotion_id,notnull"`
		CreatedAt     time.Time  `bun:"created_at"`
		ExpiresAt     time.Time  `bun:"expires_at"`
		ConfirmedAt   time.Time  `bun:"confirmed_at,nullzero"`
		DeletedAt     time.Time  `bun:"deleted_at,soft_delete,nullzero"`
	}
)
