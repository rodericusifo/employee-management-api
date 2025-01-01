package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/controller/api/request"
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"

	pkg_util_response "github.com/rodericusifo/employee-management-api/pkg/util/response"
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

	return ctx.Status(fiber.StatusOK).JSON(pkg_util_response.ResponseSuccess[any]("delete employee success", nil, nil))
}
