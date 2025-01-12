package resource

import (
	"time"

	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/mocks"
)

var (
	mockEmployeeDatabaseSQLRepository *mocks.IEmployeeDatabaseSQLRepository
	employeeResource                  IEmployeeResource
)

var (
	mockDateTime, mockBirthday                              time.Time
	mockUUID, mockAddress, mockHashPassword, mockDateString string
	mockAge                                                 int
)

func SetupTestEmployeeResource() {
	mockEmployeeDatabaseSQLRepository = new(mocks.IEmployeeDatabaseSQLRepository)

	employeeResource = InitEmployeeResource(mockEmployeeDatabaseSQLRepository)

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	mockDateString = "2015-09-02 08:04:00"
	mockDateTime, _ = time.Parse(layoutFormat, mockDateString)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}
