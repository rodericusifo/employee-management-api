package service

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/counter"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/definer"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/serializer"
)

func (s *EmployeeService) GetEmployees(payload *input.GetEmployeesDTO) (output.GetEmployeesDTO, *types.Meta, error) {
	page, limit := definer.DefinePaginationPageLimit(payload.Page, payload.Limit)

	employeeListModelRes, err := s.EmployeeResource.FindEmployees(&types.QuerySQL{
		Offset: counter.CountPaginationOffset(page, limit),
		Limit:  limit,
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}
	countEmployeeListModelRes := len(employeeListModelRes)

	if len(employeeListModelRes) < 1 {
		return nil, nil, fiber.NewError(fiber.StatusNotFound, "employees not found")
	}

	countEmployeeAllModelRes, err := s.EmployeeResource.CountEmployees(&types.QuerySQL{
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "user_id", Operator: "=", Value: payload.UserID},
				{Field: "deleted_at", Operator: "IS NULL"},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	employeeListDto := serializer.SerializeEmployeesToEmployeeDTOs(employeeListModelRes)

	meta := &types.Meta{
		CurrentPage:      page,
		CountDataPerPage: countEmployeeListModelRes,
		TotalData:        int(countEmployeeAllModelRes),
	}

	meta.TotalPage = counter.CountPaginationTotalPage(meta.CountDataPerPage, meta.TotalData)

	return employeeListDto, meta, nil
}
