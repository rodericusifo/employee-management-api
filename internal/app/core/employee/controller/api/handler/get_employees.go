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

func (h *EmployeeHandler) GetEmployees(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqQuery := new(request.GetEmployeesRequestQuery)
	if err := validator.ValidateRequestQuery(ctx, reqQuery); err != nil {
		return err
	}

	getEmployeesDtoRes, meta, err := h.EmployeeService.GetEmployees(&input.GetEmployeesDTO{
		Page:   reqQuery.Page,
		Limit:  reqQuery.Limit,
		UserID: reqUser.ID,
	})
	if err != nil {
		return err
	}

	getEmployeesResponse := serializer.SerializeEmployeeDTOsToEmployeeResponses(getEmployeesDtoRes)

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get employees success", getEmployeesResponse, meta))
}
