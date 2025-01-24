package promotion

import (
	"github.com/VictorNevola/internal/pkg/common"
	"github.com/VictorNevola/internal/pkg/entity/promotion"
	customvalidator "github.com/VictorNevola/internal/pkg/utils/custom-validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	config.App.Post("/promotion", handler.CreatePromotion)
}

func (h *httpHandler) CreatePromotion(c *fiber.Ctx) error {
	data := promotion.CreateData{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrBadRequestResponse)
	}

	if err := h.validator.Validator.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewValidationErrorResponse(err))
	}

	companyID, _ := uuid.Parse("01efc3d8-9193-6dae-8d17-00155d4d9af9")
	data.CompanyID = companyID

	promotionModel, err := h.service.CreatePromotion(c.Context(), &data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.ErrInternalServerResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(promotionModel)
}
