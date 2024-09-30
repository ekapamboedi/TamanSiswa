package model

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuAccess struct {
	Id         string         `json:"id"`
	CompanyId  sql.NullString `json:"companyId"`
	PositionId sql.NullString `json:"positionId"`
	Access     sql.NullString `json:"access"`
	CreatedAt  sql.NullTime   `json:"createdAt"`
	UpdatedAt  sql.NullTime   `json:"updatedAt"`
}

func (MenuAccess) TableName() string {
	return "menu_access"
}

func (menuAccess *MenuAccess) BeforeCreate(tx *gorm.DB) (err error) {
	menuAccess.Id = uuid.NewString()
	return
}
