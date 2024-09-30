package model

import (
	"database/sql"
)

type AvailablePackage struct {
	Id              string         `json:"id"`
	PackageName     sql.NullString `json:"packageName"`
	AvailableMenu   sql.NullString `json:"availableMenu"`
	AvailableDevice sql.NullInt64  `json:"availableDevice"`
	AvailableApps   sql.NullInt64  `json:"availableApps"`
}

func (AvailablePackage) TableName() string {
	return "available_packages"
}
