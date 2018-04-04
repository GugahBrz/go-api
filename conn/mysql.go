package conn

import (
	"fmt"

	"github.com/GugahBrz/go-api/models"
	"github.com/jinzhu/gorm"

	// Gorm mysql connector
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type connection struct {
	DatabaseHost string `default:"localhost"`
	DatabasePort string `default:"3306"`
	DatabaseName string `default:"go-api"`
	DatabaseUser string `default:"root"`
	DatabasePass string `default:"root"`
}

// InitDb init db connection
func InitDb() *gorm.DB {

	var c connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DatabaseUser, c.DatabasePass, c.DatabaseHost, c.DatabasePort, c.DatabaseName)
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}

	return db
}
