package model

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Salary struct {
	Id                     string         `gorm:"primaryKey" json:"id"`
	EmployeeId             sql.NullInt64  `json:"employeeId"`
	EmployeeName           sql.NullString `json:"employeeName"`
	CompanyId              sql.NullString `json:"companyId"`
	DepartmentId           sql.NullString `json:"departmentId"`
	DepartmentName         sql.NullString `json:"departmentName"`
	PositionId             sql.NullString `json:"positionId"`
	PositionName           sql.NullString `json:"positionName"`
	Earning                sql.NullInt64  `json:"earning"`
	MartialAndDependents   sql.NullString `json:"martialAndDependents"`
	Currency               sql.NullString `json:"currency"`
	BankName               sql.NullString `json:"bankName"`
	BankAccount            sql.NullString `json:"bankAccount"`
	PaymentMethod          sql.NullString `json:"paymentMethod"`
	BpjsEmployeeNo         sql.NullString `json:"bpjsEmployeeNo"`
	BpjsEmployeeStartPay   sql.NullTime   `json:"bpjsEmployeeStartPay"`
	BpjsEmployeeEndPay     sql.NullTime   `json:"bpjsEmployeeEndPay"`
	BpjsHealthCareNo       sql.NullString `json:"bpjsHealthCareNo"`
	BpjsHealthCareStartPay sql.NullTime   `json:"bpjsHealthCareStartPay"`
	BpjsHealthCareEndPay   sql.NullTime   `json:"bpjsHealthCareEndPay"`
	NpwpNo                 sql.NullString `json:"npwpNo"`
	TaxStartPay            sql.NullTime   `json:"taxStartPay"`
	TaxEndPay              sql.NullTime   `json:"taxEndPay"`
	CreatedAt              sql.NullTime   `json:"createdAt"`
	UpdatedAt              sql.NullTime   `json:"updatedAt"`
}

func (Salary) TableName() string {
	return "salaries"
}

func (salary *Salary) BeforeCreate(tx *gorm.DB) (err error) {
	salary.Id = uuid.NewString()
	return
}
