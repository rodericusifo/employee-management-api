package role

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
	SetupTestMysqlRoleDatabaseSQLRepository()
}

func TestMysqlRoleDatabaseSQLRepository_FirstRole(t *testing.T) {
	type (
		args struct {
			query *types.QuerySQL
		}
		result struct {
			value *sql.Role
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
						{Field: "xid"},
						{Field: "name"},
						{Field: "slug"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]types.SearchQuerySQLOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "slug", Operator: "=", Value: "super_admin"},
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
						arg1 = mockUUID
						arg2 = "super_admin"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id`,`roles`.`xid`,`roles`.`name`,`roles`.`slug`,`roles`.`created_at`,`roles`.`updated_at` FROM `roles` WHERE `roles`.`xid` = ? AND `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_first_role",
			input: args{
				query: &types.QuerySQL{
					Selects: []types.SelectQuerySQLOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "slug"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]types.SearchQuerySQLOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "slug", Operator: "=", Value: "super_admin"},
						},
					},
				},
			},
			output: result{
				value: &sql.Role{
					ID:        1,
					XID:       mockUUID,
					Name:      "Super Admin",
					Slug:      "super_admin",
					CreatedAt: mockDateTime,
					UpdatedAt: mockDateTime,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1         = mockUUID
						arg2         = "super_admin"
						rowsInstance = sqlmock.NewRows([]string{"id", "xid", "name", "slug", "created_at", "updated_at"})
					)
					rowsInstance.AddRow(1, mockUUID, "Super Admin", "super_admin", mockDateTime, mockDateTime)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							"SELECT `roles`.`id`,`roles`.`xid`,`roles`.`name`,`roles`.`slug`,`roles`.`created_at`,`roles`.`updated_at` FROM `roles` WHERE `roles`.`xid` = ? AND `roles`.`slug` = ? ORDER BY `roles`.`id` LIMIT 1",
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

			value, err := roleDatabaseSQLRepository.FirstRole(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
