package config

import (
    "github.com/SeptianSamdany/go-fiber-rest-api/entities"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB
var Database_URI = "root:@tcp(localhost:3306)/go_fiber_restapi?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
    var err error

    DB, err = gorm.Open(mysql.Open(Database_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic("Database tidak terhubung")
    }

    DB.AutoMigrate(&entities.User{}, &entities.Job{}) // Tambahkan migrasi untuk Job
    return nil
}
