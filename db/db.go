package db

import (
	"fmt"
	"gokemon/abdulsalam/config"
	"gokemon/abdulsalam/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s.db", configuration.DB_NAME)
	db, err = gorm.Open(sqlite.Open(connect_string), &gorm.Config{})

	if err != nil {
		panic("Database Connection Error")
	}

	db.AutoMigrate(&model.Pokemon{})
}

func DbManager() *gorm.DB {
	return db
}
