package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"

	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *UserDatabaseSeederSQLRepository) Clear(db *gorm.DB) error {
	role := new(sql.Role)

	q := db

	query := &types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "=", Value: "super_admin"},
			},
		},
	}

	q = builder.BuildQuerySQL(r.models.Role.TableName(), q, query, r.dialect)

	err := q.Table(r.models.Role.TableName()).First(role).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithFields(logrus.Fields{
			"message": "get role fail",
			"detail":  err,
		}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
		return err
	}

	users := make([]*sql.User, 0)

	q = db

	query = &types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "role_id", Operator: "=", Value: role.ID},
			},
		},
	}

	q = builder.BuildQuerySQL(r.models.User.TableName(), q, query, r.dialect)

	if err := q.Table(r.models.User.TableName()).Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		if err := db.Table(r.models.User.TableName()).Delete(user).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"message": "delete user fail",
				"detail":  err,
			}).Errorln("[USER DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
			continue
		}
	}

	return nil
}
