package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/controller/api/request"
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/response"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"
)

func (h *EmployeeHandler) DeleteEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqParams := new(request.DeleteEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	if err := h.EmployeeService.DeleteEmployee(&input.DeleteEmployeeDTO{
		XID:    reqParams.XID,
		UserID: reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("delete employee success", nil, nil))
}
