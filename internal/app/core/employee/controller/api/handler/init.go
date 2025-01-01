package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/service"
	"github.com/rodericusifo/employee-management-api/internal/pkg/middleware"
)

type EmployeeHandler struct {
	EmployeeService service.IEmployeeService
}

func InitEmployeeHandler(employeeService service.IEmployeeService) *EmployeeHandler {
	return &EmployeeHandler{EmployeeService: employeeService}
}

func (employeeHandler *EmployeeHandler) Mount(group fiber.Router) {
	group.Post("/create", middleware.APIUserRequest(), middleware.APIUserRolePermissions(), employeeHandler.CreateEmployee)
	group.Get("/list", middleware.APIUserRequest(), middleware.APIUserRolePermissions(), employeeHandler.GetEmployees)
	group.Get("/:xid/detail", middleware.APIUserRequest(), middleware.APIUserRolePermissions(), employeeHandler.GetEmployee)
	group.Put("/:xid/update", middleware.APIUserRequest(), middleware.APIUserRolePermissions(), employeeHandler.UpdateEmployee)
	group.Delete("/:xid/delete", middleware.APIUserRequest(), middleware.APIUserRolePermissions(), employeeHandler.DeleteEmployee)
}
