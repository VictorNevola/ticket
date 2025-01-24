package company

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
		bun.BaseModel `bun:"companies"`

		ID        *uuid.UUID `bun:"id,pk"`
		Name      string     `bun:"name"`
		TaxID     string     `bun:"tax_id,unique"`
		CreatedAt time.Time  `bun:"created_at"`
		UpdatedAt time.Time  `bun:"updated_at"`
		DeletedAt time.Time  `bun:"deleted_at,soft_delete,nullzero"`
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
