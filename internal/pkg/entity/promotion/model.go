package promotion

import (
	"context"
	"time"

	customuuid "github.com/VictorNevola/internal/pkg/utils/custom-UUID"
	time_location "github.com/VictorNevola/internal/pkg/utils/time-location"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Model struct {
		bun.BaseModel `bun:"promotions,alias:promotions"`

		ID                    *uuid.UUID `bun:"id,pk"`
		CompanyID             *uuid.UUID `bun:"company_id"`
		Name                  string     `bun:"name"`
		TextMessageInProgress string     `bun:"text_message_in_progress"`
		TextMessageSuccess    string     `bun:"text_message_success"`
		StartDate             time.Time  `bun:"start_date"`
		EndDate               time.Time  `bun:"end_date"`
		QtyMaxUsers           int        `bun:"qty_max_users"`
		VouchersPerUser       int        `bun:"vouchers_per_user"`

		CreatedAt time.Time `bun:"created_at"`
		UpdatedAt time.Time `bun:"updated_at"`
		DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero"`
	}
)

var _ bun.BeforeAppendModelHook = (*Model)(nil)

func (m *Model) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.ID = customuuid.GenerateV6()
		m.CreatedAt = time_location.Now()
		m.UpdatedAt = time_location.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time_location.Now()
	case *bun.DeleteQuery:
		m.DeletedAt = time_location.Now()
	}

	return nil
}
