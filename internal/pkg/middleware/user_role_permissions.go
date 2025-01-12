package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/getter"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/patcher"
)

func APIUserRolePermissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqUser := getter.GetRequestUser(c)

		if reqUser.Role.Slug == "super_admin" {
			return c.Next()
		}

		path := strings.Split(c.OriginalURL(), "?")[0]

		permission, err := patcher.PermissionResource().FirstPermission(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "path", Operator: "=", Value: path},
				},
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("permission with path %s not found", path))
			}
			return err
		}

		_, err = patcher.RolePermissionResource().FirstRolePermission(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "role_id", Operator: "=", Value: reqUser.Role.ID},
					{Field: "permission_id", Operator: "=", Value: permission.ID},
				},
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fiber.NewError(fiber.StatusForbidden, "access not allowed")
			}
			return err
		}

		return c.Next()
	}
}
