package role

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/rodericusifo/employee-management-api/internal/app/model/database/sql"
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/builder"
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/validator"
)

func (r *RoleDatabaseSeederSQLRepository) Seed(db *gorm.DB) error {
	roles := make([]*sql.Role, 0)
	for _, RoleSeed := range RoleSeedData {
		err := validator.ValidatePayload(RoleSeed)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"message": fmt.Sprintf("validation failed: role with xid %s", RoleSeed.XID),
				"detail":  err,
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		role := new(sql.Role)

		q := db

		query := &types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: RoleSeed.XID},
					{Field: "slug", Operator: "=", Value: RoleSeed.Slug},
				},
			},
		}

		q = builder.BuildQuerySQL(r.model.TableName(), q, query, r.dialect)

		err = q.Table(r.model.TableName()).First(role).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			logrus.WithFields(logrus.Fields{
				"message": "get role fail",
				"detail":  err,
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}
		if role.ID != 0 {
			logrus.WithFields(logrus.Fields{
				"message": fmt.Sprintf("role with xid %s and slug %s already registered", RoleSeed.XID, RoleSeed.Slug),
			}).Errorln("[ROLE DATABASE SEEDER SQL REPOSITORY] [SEED]")
			continue
		}

		roles = append(roles, &sql.Role{
			XID:  RoleSeed.XID,
			Name: RoleSeed.Name,
			Slug: RoleSeed.Slug,
		})
	}
	return db.CreateInBatches(roles, len(roles)).Error
}
