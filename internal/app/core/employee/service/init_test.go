package service

import (
	"time"

	"github.com/rodericusifo/employee-management-api/internal/pkg/constant"
	"github.com/rodericusifo/employee-management-api/internal/mocks"
)

var (
	mockEmployeeResource *mocks.IEmployeeResource
	employeeService      IEmployeeService
)

var (
	mockDate, mockBirthday       time.Time
	mockUUID, mockAddress        string
	mockAge, mockPage, mockLimit int
)

func SetupTestEmployeeService() {
	mockEmployeeResource = new(mocks.IEmployeeResource)

	employeeService = InitEmployeeService(mockEmployeeResource)

	layoutFormat := constant.DEFAULT_TIME_LAYOUT.(string)

	valueDate := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, valueDate)

	valueBirthday := "1999-03-12 00:00:00"
	mockBirthday, _ = time.Parse(layoutFormat, valueBirthday)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
	mockAddress = "Street A, City B"
	mockAge = 25
	mockPage = 1
	mockLimit = 10
}
