package company

import (
	"net/http"

	"github.com/VictorNevola/internal/pkg/common"
	"github.com/VictorNevola/internal/pkg/entity/company"
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

	config.App.Post("/company", handler.CreateCompany)
}

func (h *httpHandler) CreateCompany(c *fiber.Ctx) error {
	data := &company.CreateData{}

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrBadRequestResponse)
	}

	if err := h.validator.Validator.Struct(data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.NewValidationErrorResponse(err))
	}

	companyCreated, err := h.service.CreateCompany(c.Context(), data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(common.ErrInternalServerResponse)
	}

	return c.Status(http.StatusCreated).JSON(companyCreated)
}
