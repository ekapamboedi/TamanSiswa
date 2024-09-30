package model

import (
	"database/sql"
)

type Position struct {
	Id             string         `gorm:"primaryKey" json:"id"`
	DepartmentCode sql.NullString `json:"departmentCode"`
	DepartmentName sql.NullString `json:"departmentName"`
	PositionCode   sql.NullString `json:"positionCode"`
	PositionName   sql.NullString `json:"positionName"`
	CompanyName    sql.NullString `json:"companyName"`
	CompanyId      sql.NullString `json:"companyId"`
	CreatedAt      sql.NullTime   `json:"createdAt"`
	CreatedBy      sql.NullString `json:"createdBy"`
	UpdatedAt      sql.NullTime   `json:"updatedAt"`
	UpdatedBy      sql.NullString `json:"updatedBy"`
}

func (Position) TableName() string {
	return "master_positions"
}
