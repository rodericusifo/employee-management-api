package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodericusifo/employee-management-api/internal/pkg/util/handler"
	"github.com/rodericusifo/employee-management-api/mocks"
)

var (
	mockApp         *fiber.App
	mockAuthService *mocks.IAuthService
	authHandler     *AuthHandler
)

var (
	mockEmail, mockPassword, mockJWTToken string
)

func SetupTestAuthHandler() {
	mockApp = fiber.New(fiber.Config{
		ErrorHandler: handler.APIError,
	})

	mockAuthService = new(mocks.IAuthService)

	auth := mockApp.Group("/auth")
	authHandler = InitAuthHandler(mockAuthService)
	authHandler.Mount(auth)

	mockEmail = "john@gmail.com"
	mockPassword = "john1223"

	mockJWTToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjhlYTc3OGJjLTM5NTgtNGU5Zi04ZmEyLWE4YTlhZDhmMmFiMSIsIm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AZ21haWwuY29tIiwicm9sZSI6IkFETUlOIiwiZXhwIjoxNjc3MDc5NzgxfQ.bndXk_BggjadIF2Rwluxc-3tPr-ArfWVYTZ5y03wHU8"
}
