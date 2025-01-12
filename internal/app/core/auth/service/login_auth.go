package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/app/core/auth/service/dto/output"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/patcher"
)

func (s *AuthService) LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error) {
	userModelRes, err := s.UserResource.FirstUser(&types.QuerySQL{
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return nil, err
	}

	match := patcher.CompareHashAndPassword(userModelRes.Password, payload.Password)
	if !match {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "email and password not match")
	}

	claims := &types.JwtCustomClaims{
		XID: userModelRes.XID,
	}

	token, err := patcher.GenerateJWTTokenFromClaims(claims)
	if err != nil {
		return nil, err
	}

	loginAuthDto := &output.LoginAuthDTO{
		Token: token,
	}

	return loginAuthDto, nil
}
