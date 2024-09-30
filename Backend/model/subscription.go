package model

import (
	"database/sql"
)

type Subscription struct {
	Id                      string         `json:"id"`
	CompanyId               sql.NullString `json:"companyId"`
	CompanyName             sql.NullString `json:"companyName"`
	SubscribedPackage       sql.NullString `json:"subscribedPackage"`
	DeviceTotal             sql.NullInt64  `json:"deviceTotal"`
	DeviceUsed              sql.NullInt64  `json:"deviceUsed"`
	RemainingDevice         sql.NullInt64  `json:"remainingDevice"`
	AppsTotal               sql.NullInt64  `json:"appsTotal"`
	AppsUsed                sql.NullInt64  `json:"appsUsed"`
	RemainingApps           sql.NullInt64  `json:"remainingApps"`
	CompanyStartSubscribe   sql.NullTime   `json:"companyStartSubscribe"`
	CompanyExpiredSubscribe sql.NullTime   `json:"companyExpiredSubscribe"`
	CreatedBy               sql.NullString `json:"createdBy"`
	UpdatedBy               sql.NullString `json:"updatedBy"`
	UpdatedAt               sql.NullTime   `json:"updatedAt"`
	CreatedAt               sql.NullTime   `json:"createdAt"`
}

func (Subscription) TableName() string {
	return "detail_subscriptions"
}
