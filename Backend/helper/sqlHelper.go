package helper

import (
	"database/sql"
	"time"
)

func ValidateSqlString(data string) *sql.NullString {
	if data == "" {
		return &sql.NullString{Valid: false}
	}

	return &sql.NullString{String: data, Valid: true}
}

func ValidateSqlTime(data time.Time) *sql.NullTime {
	if data.IsZero() {
		return &sql.NullTime{Valid: false}
	}

	return &sql.NullTime{Time: data, Valid: true}
}

func ValidateSqlInt64(data int64) *sql.NullInt64 {
	if data == 0 {
		return &sql.NullInt64{Valid: false}
	}

	return &sql.NullInt64{Int64: data, Valid: true}
}
