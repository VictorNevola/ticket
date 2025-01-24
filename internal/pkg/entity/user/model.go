package user

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
		bun.BaseModel `bun:"users,alias:users"`
		ID            *uuid.UUID `bun:"id,pk"`
		Username      string     `bun:"username"`
		Email         string     `bun:"email,unique"`
		Password      string     `bun:"password"`

		CreatedAt time.Time `bun:"created_at"`
		UpdatedAt time.Time `bun:"updated_at"`
		DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero"`
	}

	ModelUserInPromotion struct {
		bun.BaseModel `bun:"users_in_promotions,alias:users_in_promotions"`
		ID            *uuid.UUID `bun:"id,pk"`
		UserID        uuid.UUID  `bun:"user_id,notnull"`
		PromotionID   uuid.UUID  `bun:"promotion_id,notnull"`
		CreatedAt     string     `bun:"created_at"`
		UpdateAt      string     `bun:"update_at"`
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
