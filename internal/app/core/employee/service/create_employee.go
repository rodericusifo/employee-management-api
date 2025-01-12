package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func (s *EmployeeService) CreateEmployee(payload *input.CreateEmployeeDTO) error {
	employeeModelRes, err := s.EmployeeResource.FirstEmployee(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
		WithDeleted: true,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if employeeModelRes != nil {
		return fiber.NewError(fiber.StatusConflict, "employee already registered")
	}

	employeeModel := &sql.Employee{
		Name:     payload.Name,
		Email:    payload.Email,
		Address:  payload.Address,
		Age:      payload.Age,
		Birthday: payload.Birthday,
		UserID:   payload.UserID,
	}
	err = s.EmployeeResource.SaveEmployee(employeeModel)
	if err != nil {
		return err
	}

	return nil
}
