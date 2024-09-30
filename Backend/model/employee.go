package model

import (
	"database/sql"
)

type Employee struct {
	Id               sql.NullString `json:"id"`
	EmployeeId       sql.NullString `json:"employeeId"`
	EmployeeName     sql.NullString `json:"employeeName"`
	AccountTypeId    sql.NullInt64  `json:"accountTypeId"`
	DepartmentId     sql.NullString `json:"departmentId"`
	DepartmentName   sql.NullString `json:"departmentName"`
	DepartmentCode   sql.NullString `json:"departmentCode"`
	PositionId       sql.NullString `json:"positionId"`
	PositionName     sql.NullString `json:"positionName"`
	PositionCode     sql.NullString `json:"positionCode"`
	EmployeeTypeName sql.NullString `json:"employeeTypeName"`
	EmployeeTypeCode sql.NullString `json:"employeeTypeCode"`
	CompanyName      sql.NullString `json:"companyName"`
	CompanyId        sql.NullString `json:"companyId"`
	Gender           sql.NullString `json:"gender"`
	Birthday         sql.NullTime   `json:"birthday"`
	PhoneNumber      sql.NullString `json:"phoneNumber"`
	Address          sql.NullString `json:"address"`
	PostalCode       sql.NullString `json:"postalCode"`
	National         sql.NullString `json:"national"`
	Religion         sql.NullString `json:"religion"`
	EmailAddress     sql.NullString `json:"emailAddress"`
	ProfilePicture   sql.NullString `json:"profilePicture"`
	DevicePhoto      sql.NullString `json:"devicePhoto"`
	Avatar           sql.NullString `json:"avatar"`
	WorkTypeId       sql.NullString `json:"workTypeId"`
	WorkTypeName     sql.NullString `json:"workTypeName"`
	AreaId           sql.NullString `json:"areaId"`
	AreaCode         sql.NullString `json:"areaCode"`
	AreaName         sql.NullString `json:"areaName"`
	ShiftId          sql.NullString `json:"shiftId"`
	ShiftName        sql.NullString `json:"shiftName"`
	ShiftCode        sql.NullString `json:"shiftCode"`
	IsResign         sql.NullBool   `json:"isResign"`
	JoinDate         sql.NullTime   `json:"joinDate"`
	ResignDate       sql.NullTime   `json:"resignDate"`
	EffectiveStart   sql.NullTime   `json:"effectiveStart"`
	EffectiveEnd     sql.NullTime   `json:"effectiveEnd"`
	CreatedBy        sql.NullString `json:"createdBy"`
	UpdatedBy        sql.NullString `json:"updatedBy"`
	CreatedAt        sql.NullTime   `json:"createdAt"`
	UpdatedAt        sql.NullTime   `json:"updatedAt"`
}

func (Employee) TableName() string {
	return "employees"
}
