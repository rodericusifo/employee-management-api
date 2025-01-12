package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/app/core/auth/controller/api/handler"

	registry_service_auth "github.com/rodericusifo/employee-management-api/registry/service/auth"
)

func InitAPI(router fiber.Router) {
	auth := router.Group("/auth")
	authService := registry_service_auth.AuthService()
	authHandler := handler.InitAuthHandler(authService)
	authHandler.Mount(auth)
}
