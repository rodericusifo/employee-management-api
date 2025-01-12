package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/controller/api/request"
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/response"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/serializer"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"
)

func (h *EmployeeHandler) GetEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqParams := new(request.GetEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	employeeDtoRes, err := h.EmployeeService.GetEmployee(&input.GetEmployeeDTO{
		XID:    reqParams.XID,
		UserID: reqUser.ID,
	})
	if err != nil {
		return err
	}

	getEmployeeResponse := serializer.SerializeEmployeeDTOToEmployeeResponse(employeeDtoRes)

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get employee success", getEmployeeResponse, nil))
}
