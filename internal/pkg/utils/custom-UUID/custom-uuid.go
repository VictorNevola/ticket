package customuuid

import "github.com/google/uuid"

func GenerateV6() *uuid.UUID {
	uuidV6, _ := uuid.NewV6()

	return &uuidV6
}
