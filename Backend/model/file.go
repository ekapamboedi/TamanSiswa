package model

import (
	"database/sql"
)

type File struct {
	Id        int64          `json:"id"`
	Url       sql.NullString `json:"url"`
	Filename  sql.NullString `json:"filename"`
	Uuid      sql.NullString `json:"uuid"`
	UpdatedAt sql.NullTime   `json:"updatedAt"`
	CreatedAt sql.NullTime   `json:"createdAt"`
	DeletedAt sql.NullTime   `json:"deletedAt"`
}

func (File) TableName() string {
	return "files"
}
