package user_test

import (
	"testing"

	"github.com/VictorNevola/internal/domain/user"
	userEntity "github.com/VictorNevola/internal/pkg/entity/user"
	"github.com/VictorNevola/test/testhelpers"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	ctx, dbCleanup := testContext()
	defer dbCleanup()

	t.Run("should create a user successfully", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		// before user not exists
		hasUser := getUser(ctx)
		assert.Nil(t, hasUser)

		userService := ctx.Value(userServiceKey).(user.Service)
		userCreated, err := userService.CreateUser(ctx, &userEntity.CreateData{
			Username: "user_test",
			Email:    "test@test.com",
			Password: "123456",
		})

		// after user exists
		assert.Nil(t, err)
		assert.NotNil(t, userCreated)
		assert.Equal(t, "user_test", userCreated.Username)
		assert.Equal(t, "test@test.com", userCreated.Email)

		// Validate that the password is hashed
		err = bcrypt.CompareHashAndPassword([]byte(userCreated.Password), []byte("123456"))
		assert.Nil(t, err)
	})

	t.Run("should not create a user with the same email", func(t *testing.T) {
		defer testhelpers.ClearAllDataBase(ctx)

		userService := ctx.Value(userServiceKey).(user.Service)
		_, err := userService.CreateUser(ctx, &userEntity.CreateData{
			Username: "user_test1",
			Email:    "test@test.com",
			Password: "123456",
		})
		assert.Nil(t, err)

		_, err = userService.CreateUser(ctx, &userEntity.CreateData{
			Username: "user_test2",
			Email:    "test@test.com",
			Password: "654321",
		})
		assert.NotNil(t, err)
	})
}
