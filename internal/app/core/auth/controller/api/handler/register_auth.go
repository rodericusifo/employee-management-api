package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/auth/controller/api/request"
	"github.com/rodericusifo/employee-management-api/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/response"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"
)

func (h *AuthHandler) RegisterAuth(ctx *fiber.Ctx) error {
	reqBody := new(request.RegisterAuthRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	err := h.AuthService.RegisterAuth(&input.RegisterAuthDTO{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
		// RoleSlug: reqBody.RoleSlug,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("auth register success", nil, nil))
}
