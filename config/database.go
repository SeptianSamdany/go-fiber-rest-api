package config

import (
	"github.com/SeptianSamdany/go-fiber-rest-api/entities"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var Database *gorm.DB
var Database_URI = "root:@tcp(localhost:3306)/go_fiber_restapi?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(Database_URI), &gorm.Config{
	 SkipDefaultTransaction: true,
	 PrepareStmt:            true,
	})
   
	if err != nil {
	 panic("Database tidak tehubung")
	}
   
	Database.AutoMigrate(&entities.User{})
   
	return nil   
}