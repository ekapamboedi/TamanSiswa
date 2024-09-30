package middleware

import (
	"database/sql"
	"time"

	"TamanSiswa/helper"
	"TamanSiswa/model"
)

func LogError(endpoint string, method string, message string, identity helper.TokenPayload) {
	var errorLog = model.ErrorLog{
		// Endpoint:     sql.NullString{String: endpoint, Valid: true},
		Method:       sql.NullString{String: method, Valid: true},
		Username:     sql.NullString{String: identity.Username, Valid: true},
		Role:         sql.NullString{String: identity.Role, Valid: true},
		ErrorMessage: sql.NullString{String: message, Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		// EmployeeId:   sql.NullString{String: identity.EmployeeId, Valid: true},
		// CompanyId:    sql.NullString{String: identity.CompanyId, Valid: true},
		// CompanyName:  sql.NullString{String: identity.CompanyName, Valid: true},
		// Description:  sql.NullString{String: description, Valid: true},
	}

	model.DB.Table("error_logs").Create(&errorLog)
}

func LogErrorWithDescription(endpoint string, method string, message string, description string, identity helper.TokenPayload) {
	var errorLog = model.ErrorLog{
		// Endpoint:     sql.NullString{String: endpoint, Valid: true},
		Method:       sql.NullString{String: method, Valid: true},
		Username:     sql.NullString{String: identity.Username, Valid: true},
		Role:         sql.NullString{String: identity.Role, Valid: true},
		ErrorMessage: sql.NullString{String: message, Valid: true},
		CreatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
		// EmployeeId:   sql.NullString{String: identity.EmployeeId, Valid: true},
		// CompanyId:    sql.NullString{String: identity.CompanyId, Valid: true},
		// CompanyName:  sql.NullString{String: identity.CompanyName, Valid: true},
		// Description:  sql.NullString{String: description, Valid: true},
	}

	model.DB.Table("error_logs").Create(&errorLog)
}
