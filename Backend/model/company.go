package model

import (
	"database/sql"
)

type Company struct {
	Id                  string         `json:"id"`
	CompanyName         sql.NullString `json:"companyName"`
	CompanyLocation     sql.NullString `json:"companyLocation"`
	CompanyCity         sql.NullString `json:"companyCity"`
	Longitude           sql.NullString `json:"longitude"`
	Latitude            sql.NullString `json:"latitude"`
	CompanyBusinessType sql.NullString `json:"companyBusinessType"`
	Director            sql.NullString `json:"director"`
	ContactNumber       sql.NullString `json:"contactNumber"`
	PhoneNumber1        sql.NullString `json:"phoneNumber1"`
	PhoneNumber2        sql.NullString `json:"phoneNumber2"`
	Pic                 sql.NullString `json:"pic"`
	Icon                sql.NullString `json:"icon"`
	IsActive            sql.NullBool   `json:"isActive"`
	IsLock              sql.NullBool   `json:"isLock"`
	SubcribeDate        sql.NullTime   `json:"subcribeDate"`
	ExpiredDate         sql.NullTime   `json:"expiredDate"`
	CreatedAt           sql.NullTime   `json:"createdAt"`
	UpdatedAt           sql.NullTime   `json:"updatedAt"`
}

func (Company) TableName() string {
	return "master_companies"
}
