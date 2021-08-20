package api

import (
	"gokemon/abdulsalam/db"
	"os"

	"gorm.io/gorm"
)

func DatabaseManager() *gorm.DB {
	return db.DbManager()
}

func BaseUrl() string {
	return os.Getenv("BASE_POKEMON_API")
}

func BaseLimit() string {
	return os.Getenv("LIMIT_POKEMON_API")
}
