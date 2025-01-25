package userinpromotion

import "github.com/google/uuid"

type (
	BindUserToPromotionData struct {
		PromotionID uuid.UUID `json:"promotion_id" validate:"required,uuid"`
		UserID      uuid.UUID `json:"user_id" validate:"required,uuid"`
	}
)
