package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/auth/controller/api/handler"

	internal_registry_service_auth "github.com/rodericusifo/employee-management-api/internal/registry/service/auth"
)

func InitAPI(router fiber.Router) {
	auth := router.Group("/auth")
	authService := internal_registry_service_auth.AuthService()
	authHandler := handler.InitAuthHandler(authService)
	authHandler.Mount(auth)
}
