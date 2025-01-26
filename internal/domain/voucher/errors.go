package voucher

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrPromotionHasNoAvailableQuantityOfVouchers = fiber.NewError(
		fiber.StatusUnprocessableEntity,
		"Promotion has no available quantity of vouchers",
	)
	ErrPromotionHasNoAvailableQuantityOfUsers = fiber.NewError(
		fiber.StatusUnprocessableEntity,
		"Promotion has of maximum users reached",
	)
	ErrPromotionIsNotActive = fiber.NewError(
		fiber.StatusUnprocessableEntity,
		"Promotion has expired",
	)
	ErrUserNotInPromotion = fiber.NewError(
		fiber.StatusUnprocessableEntity,
		"User not in promotion",
	)
)
