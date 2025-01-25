package userinpromotion

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
		bun.BaseModel `bun:"users_in_promotions"`

		ID          *uuid.UUID `bun:"id,pk"`
		PromotionID uuid.UUID  `bun:"promotion_id"`
		UserID      uuid.UUID  `bun:"user_id"`
		CreatedAt   time.Time  `bun:"created_at"`
		UpdatedAt   time.Time  `bun:"updated_at"`
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
	}

	return nil
}
