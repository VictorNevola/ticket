package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	ErrorsResponse struct {
		Value   string `json:"value"`
		Message string `json:"message"`
	}
	ValidationErrorResponse struct {
		Code    int                          `json:"code"`
		Message string                       `json:"message"`
		Errors  map[string][]*ErrorsResponse `json:"errors"`
	}
)

var (
	ErrBadRequestResponse     = fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	ErrInternalServerResponse = fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
)

func NewValidationErrorResponse(err error) *ValidationErrorResponse {
	validationErrors := err.(validator.ValidationErrors)
	errors := make(map[string][]*ErrorsResponse)

	for _, e := range validationErrors {
		errors[e.Field()] = append(errors[e.StructField()], &ErrorsResponse{
			Value:   e.Value().(string),
			Message: e.Error(),
		})
	}

	return &ValidationErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: "Validation error",
		Errors:  errors,
	}
}
