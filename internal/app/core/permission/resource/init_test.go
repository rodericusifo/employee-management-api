package resource

import (
	"time"

	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/mocks"
)

var (
	mockPermissionDatabaseSQLRepository *mocks.IPermissionDatabaseSQLRepository
	permissionResource                  IPermissionResource
)

var (
	mockDateTime                               time.Time
	mockUUID, mockHashPassword, mockDateString string
)

func SetupTestPermissionResource() {
	mockPermissionDatabaseSQLRepository = new(mocks.IPermissionDatabaseSQLRepository)

	permissionResource = InitPermissionResource(mockPermissionDatabaseSQLRepository)

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
