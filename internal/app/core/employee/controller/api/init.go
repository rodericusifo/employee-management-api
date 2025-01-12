package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/controller/api/handler"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	jwtware "github.com/gofiber/contrib/jwt"

	internal_registry_service_employee "github.com/rodericusifo/employee-management-api/internal/registry/service/employee"
)

func InitAPI(router fiber.Router) {
	employee := router.Group("/employees")
	employee.Use(jwtware.New(*getter.GetJWTAuthConfig()))
	employeeService := internal_registry_service_employee.EmployeeService()
	employeeHandler := handler.InitEmployeeHandler(employeeService)
	employeeHandler.Mount(employee)
}
