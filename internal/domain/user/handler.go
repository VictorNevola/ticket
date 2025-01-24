package user

import (
	"github.com/VictorNevola/internal/pkg/common"
	"github.com/VictorNevola/internal/pkg/entity/user"
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

	config.App.Post("/user", handler.CreateUser)
}

func (h *httpHandler) CreateUser(c *fiber.Ctx) error {
	data := user.CreateData{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrBadRequestResponse)
	}

	if err := h.validator.Validator.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewValidationErrorResponse(err))
	}

	userModel, err := h.service.CreateUser(c.Context(), &data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.ErrInternalServerResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(userModel)
}
