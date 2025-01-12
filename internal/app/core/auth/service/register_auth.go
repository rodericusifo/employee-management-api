package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/patcher"
)

func (s *AuthService) RegisterAuth(payload *input.RegisterAuthDTO) error {
	payload.RoleSlug = "super_admin"

	userModelRes, err := s.UserResource.FirstUser(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if userModelRes != nil {
		return fiber.NewError(fiber.StatusConflict, "user already registered")
	}

	roleModelRes, err := s.RoleResource.FirstRole(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "=", Value: payload.RoleSlug},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "role not found")
		}
		return err
	}

	hashedPassword, err := patcher.GenerateHashFromPassword(payload.Password)
	if err != nil {
		return err
	}

	userModel := &sql.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		RoleID:   roleModelRes.ID,
	}
	err = s.UserResource.SaveUser(userModel)
	if err != nil {
		return err
	}

	return nil
}
