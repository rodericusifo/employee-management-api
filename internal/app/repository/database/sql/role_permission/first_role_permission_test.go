package role_permission

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func init() {
	SetupTestMysqlRolePermissionDatabaseSQLRepository()
}

func TestMysqlRolePermissionDatabaseSQLRepository_FirstRolePermission(t *testing.T) {
	type (
		args struct {
			query *types.QuerySQL
		}
		result struct {
			value *sql.RolePermission
			err   error
		}
	)

	testCases := []struct {
		desc   string
		input  args
		output result
		before func()
		after  func()
	}{
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				query: &types.QuerySQL{
					Selects: []types.SelectQuerySQLOperation{
						{Field: "id"},
					},
					Searches: [][]types.SearchQuerySQLOperation{
						{
							{Field: "role_id", Operator: "=", Value: 1},
							{Field: "permission_id", Operator: "=", Value: 2},
						},
					},
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = 1
						arg2 = 2
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `role_permissions`.`id` FROM `role_permissions` WHERE `role_permissions`.`role_id` = ? AND `role_permissions`.`permission_id` = ? ORDER BY `role_permissions`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_first_role_permission",
			input: args{
				query: &types.QuerySQL{
					Selects: []types.SelectQuerySQLOperation{
						{Field: "id"},
					},
					Searches: [][]types.SearchQuerySQLOperation{
						{
							{Field: "role_id", Operator: "=", Value: 1},
							{Field: "permission_id", Operator: "=", Value: 2},
						},
					},
				},
			},
			output: result{
				value: &sql.RolePermission{
					ID: 1,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1         = 1
						arg2         = 2
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `role_permissions`.`id` FROM `role_permissions` WHERE `role_permissions`.`role_id` = ? AND `role_permissions`.`permission_id` = ? ORDER BY `role_permissions`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := rolePermissionDatabaseSQLRepository.FirstRolePermission(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
