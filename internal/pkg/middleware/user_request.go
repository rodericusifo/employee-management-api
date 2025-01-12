package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/patcher"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"
)

func APIUserRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals(constant.CONTEXT_KEY_USER).(*jwt.Token).Claims
		user, ok := claims.(*types.JwtCustomClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("invalid claims type. correct type: %T", claims))
		}

		userModelRes, err := patcher.UserResource().FirstUser(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
				{Field: "name"},
				{Field: "email"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: user.XID},
				},
			},
			Joins: []types.JoinQuerySQLOperation{
				{
					Relation: "Role",
					Selects: []types.SelectJoinQuerySQLOperation{
						{Field: "id"},
						{Field: "slug"},
					},
				},
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fiber.NewError(fiber.StatusUnauthorized, "user not found")
			}
			return err
		}

		reqUser := new(types.RequestUser)
		reqUser = &types.RequestUser{
			ID:    userModelRes.ID,
			XID:   user.XID,
			Name:  userModelRes.Name,
			Email: userModelRes.Email,
			Role: types.RequestRole{
				ID:   userModelRes.Role.ID,
				Slug: userModelRes.Role.Slug,
			},
		}
		if err := validator.ValidateRequestUser(reqUser); err != nil {
			return err
		}

		c.Locals(constant.CONTEXT_KEY_REQUEST_USER, reqUser)
		return c.Next()
	}
}
