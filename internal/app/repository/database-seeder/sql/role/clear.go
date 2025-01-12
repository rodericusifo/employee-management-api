package role

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
)

func (r *RoleDatabaseSeederSQLRepository) Clear(db *gorm.DB) error {
	roleSeedSlugs := make([]string, 0)
	roles := make([]*sql.Role, 0)

	for _, RoleSeed := range RoleSeedData {
		roleSeedSlugs = append(roleSeedSlugs, RoleSeed.Slug)
	}

	q := db

	query := &types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "IN", Value: roleSeedSlugs},
			},
		},
	}

	q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)

	if err := q.Table(r.model.TableName()).Find(&roles).Error; err != nil {
		return err
	}

	for _, role := range roles {
		if err := db.Table(r.model.TableName()).Delete(role).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"message": "delete role fail",
				"detail":  err,
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [CLEAR]")
			continue
		}
	}

	return nil
}
