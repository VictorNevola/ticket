package voucher

import "github.com/google/uuid"

type (
	CreateVoucherDataBody struct {
		PromotionID uuid.UUID `json:"promotion_id" validate:"required,uuid"`
	}
)
