package voucher

import (
	"github.com/VictorNevola/internal/pkg/common"
	"github.com/VictorNevola/internal/pkg/entity/voucher"
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

	config.App.Post("/voucher", handler.GenerateVoucher)
}

func (h *httpHandler) GenerateVoucher(c *fiber.Ctx) error {
	data := voucher.CreateVoucherDataBody{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrBadRequestResponse)
	}

	if err := h.validator.Validator.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewValidationErrorResponse(err))
	}

	err := h.service.GenerateVoucher(c.Context(), data, uuid.MustParse("01efc3d8-c008-61c5-9aa2-00155d4d9af9"))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
