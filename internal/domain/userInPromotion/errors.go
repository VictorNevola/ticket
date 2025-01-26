package userinpromotion

import "github.com/gofiber/fiber/v2"

var (
	ErrUserAlreadyInPromotion                 = fiber.NewError(fiber.StatusBadRequest, "User already in promotion")
	ErrPromotionHasNoAvailableQuantityOfUsers = fiber.NewError(fiber.StatusBadRequest, "Promotion has no available quantity of users")
	ErrPromotionHasExpired                    = fiber.NewError(fiber.StatusBadRequest, "Promotion has expired")
	ErrPromotionNotFound                      = fiber.NewError(fiber.StatusBadRequest, "Promotion not found")
)
