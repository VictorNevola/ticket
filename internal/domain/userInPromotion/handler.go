package userinpromotion

import (
	"errors"

	"github.com/VictorNevola/internal/pkg/common"
	userinpromotion "github.com/VictorNevola/internal/pkg/entity/userInPromotion"
	customvalidator "github.com/VictorNevola/internal/pkg/utils/custom-validator"
	"github.com/gofiber/fiber/v2"
)

type (
	httpHandler struct {
		service   Service
		validator *customvalidator.CustomValidator
	}

	HttpHandlerParams struct {
		Service   Service
		App       *fiber.App
		Validator *customvalidator.CustomValidator
	}
)

func NewHTTPHandler(config *HttpHandlerParams) {
	handler := &httpHandler{
		service:   config.Service,
		validator: config.Validator,
	}

	config.App.Post("/user-in-promotion", handler.BindUserToPromotion)
}

func (h *httpHandler) BindUserToPromotion(c *fiber.Ctx) error {
	data := userinpromotion.BindUserToPromotionData{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrBadRequestResponse)
	}

	if err := h.validator.Validator.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewValidationErrorResponse(err))
	}

	err := h.service.BindUserToPromotion(c.Context(), &data)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyInPromotion) {
			return c.Status(fiber.StatusBadRequest).JSON(ErrUserAlreadyInPromotion)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(common.ErrInternalServerResponse)
	}

	return c.SendStatus(fiber.StatusCreated)
}
