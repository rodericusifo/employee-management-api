package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/employee/controller/api/handler"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"

	jwtware "github.com/gofiber/contrib/jwt"

	lib_wire_core_service_employee "github.com/rodericusifo/employee-management-api/lib/injector/core/service/employee"
)

func InitAPI(router fiber.Router) {
	employee := router.Group("/employees")
	employee.Use(jwtware.New(*getter.GetJWTAuthConfig()))
	employeeService := lib_wire_core_service_employee.EmployeeService()
	employeeHandler := handler.InitEmployeeHandler(employeeService)
	employeeHandler.Mount(employee)
}
