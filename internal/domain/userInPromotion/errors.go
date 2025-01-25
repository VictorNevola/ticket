package userinpromotion

import "github.com/gofiber/fiber/v2"

var (
	ErrUserAlreadyInPromotion = fiber.NewError(fiber.StatusBadRequest, "User already in promotion")
)
