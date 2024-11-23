package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBI struct {
	db *gorm.DB
}

var DB DBI

func ConnectDB() {
	dsn := fmt.Sprintf()
	gorm.Open(postgres.Open(), gorm_options)
}
