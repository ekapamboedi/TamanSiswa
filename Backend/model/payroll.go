package model

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payroll struct {
	Id           string         `gorm:"primaryKey" json:"id"`
	CompanyId    sql.NullString `json:"companyId"`
	PositionId   sql.NullString `gorm:"uniqueIndex" json:"positionId"`
	PositionName sql.NullString `json:"positionName"`
	Properties   sql.NullString `json:"properties"`
	Deductions   sql.NullString `json:"deductions"`
	CreatedBy    sql.NullString `json:"createdBy"`
	UpdatedBy    sql.NullString `json:"updatedBy"`
	CreatedAt    sql.NullTime   `json:"createdAt"`
	UpdatedAt    sql.NullTime   `json:"updatedAt"`
	EmployeeId   sql.NullString `json:"employeeId"`
	EmployeeName sql.NullString `json:"employeeName"`
}

func (Payroll) TableName() string {
	return "master_payrolls"
}

func (payroll *Payroll) BeforeCreate(tx *gorm.DB) (err error) {
	payroll.Id = uuid.NewString()
	return
}
