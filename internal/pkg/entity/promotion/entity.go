package promotion

import "github.com/google/uuid"

type (
	CreateData struct {
		Name                  string `json:"name" validate:"required,min=3,max=255"`
		TextMessageInProgress string `json:"text_message_in_progress" validate:"required,min=3,max=3000"`
		TextMessageSuccess    string `json:"text_message_success" validate:"required,min=3,max=3000"`
		StartDate             string `json:"start_date" validate:"required,datetime=2006-01-02 15:04:05"`
		EndDate               string `json:"end_date" validate:"required,datetime=2006-01-02 15:04:05"`
		QtyMaxUsers           int    `json:"qty_max_users" validate:"required,min=1"`
		VouchersPerUser       int    `json:"vouchers_per_user" validate:"required,min=1"`
		CompanyID             uuid.UUID
	}
)
