package conn

import (
	"github.com/GugahBrz/go-api/models"
	"github.com/jinzhu/gorm"

	// Gorm mysql connector
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitDb init db connection
func InitDb() *gorm.DB {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/go-api")

	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}

	return db
}
