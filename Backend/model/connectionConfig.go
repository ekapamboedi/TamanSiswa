package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"TamanSiswa/config"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println(err.Error())
	}

	connPool, _ := database.DB()
	connPool.SetMaxOpenConns(5)

	DB = database
}
