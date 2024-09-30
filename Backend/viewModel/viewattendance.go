package viewmodel

type ViewAttendances struct {
	EmployeeName    string `json:"employeeName"`
	Date            string `json:"date"`
	MaximumLate     string `json:"maximumLate"`
	Late            string `json:"late"`
	StartWorkTime   string `json:"startWorkTime"`
	ClockIn         string `json:"clockIn"`
	ClockInMethod   string `json:"clockInMethod"`
	ClockInAddress  string `json:"clockInAddress"`
	ClockInNote     string `json:"clockInNote"`
	ClockInPhoto    string `json:"clockInPhoto"`
	ClockOut        string `json:"clockOut"`
	ClockOutMethod  string `json:"clockOutMethod"`
	ClockOutAddress string `json:"clockOutAddress"`
	ClockOutNote    string `json:"clockOutNote"`
	ClockOutPhoto   string `json:"clockOutPhoto"`
	DepartmentName  string `json:"departmentName"`
	PositionName    string `json:"positionName"`
	EmployeeShift   string `json:"employeeShift"`
	EmployeeId      string `json:"employeeId"`
	CompanyId       string `json:"companyId"`
	Id              int64  `json:"id"`
}

func (ViewAttendances) TableName() string {
	return "view_attendances"
}
