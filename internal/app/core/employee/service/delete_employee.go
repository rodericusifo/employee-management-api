package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (s *EmployeeService) DeleteEmployee(payload *input.DeleteEmployeeDTO) error {
	employeeModelRes, err := s.EmployeeResource.FirstEmployee(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "employee not found")
		}
		return err
	}

	employeeModel := employeeModelRes

	err = s.EmployeeResource.DeleteEmployee(employeeModel)
	if err != nil {
		return err
	}

	return nil
}
