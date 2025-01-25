package user_test

import (
	"context"

	"github.com/VictorNevola/internal/domain/user"
	"github.com/VictorNevola/internal/infra/adapters/postgresql"
	userEntity "github.com/VictorNevola/internal/pkg/entity/user"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/uptrace/bun"
)

const (
	userServiceKey testhelpers.ContextKey = "userService"
)

func testContext() (context.Context, func()) {
	ctx := context.TODO()
	db, dbCleanup, _ := testhelpers.ConnectionToDB(ctx)

	userRepository := postgresql.NewUserRepository(db)
	userService := user.NewService(&user.ServiceParams{
		UserRepository: userRepository,
	})

	ctx = context.WithValue(ctx, testhelpers.DbKey, db)
	ctx = context.WithValue(ctx, userServiceKey, userService)

	return ctx, dbCleanup
}

func getUser(ctx context.Context) *userEntity.Model {
	db := ctx.Value(testhelpers.DbKey).(*bun.DB)
	userModel := &userEntity.Model{}
	err := db.NewSelect().Model(userModel).Scan(ctx)
	if err != nil {
		return nil
	}
	return userModel
}
