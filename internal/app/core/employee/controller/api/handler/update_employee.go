package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/controller/api/request"
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/response"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"
)

func (h *EmployeeHandler) UpdateEmployee(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqBody := new(request.UpdateEmployeeRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	reqParams := new(request.UpdateEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	birthdayString := *reqBody.Birthday
	birthdayTime, err := time.Parse(constant.DEFAULT_TIME_LAYOUT.(string), birthdayString)
	if err != nil {
		return err
	}

	if err := h.EmployeeService.UpdateEmployee(&input.UpdateEmployeeDTO{
		XID:      reqParams.XID,
		Address:  reqBody.Address,
		Age:      reqBody.Age,
		Birthday: &birthdayTime,
		UserID:   reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("update employee success", nil, nil))
}
