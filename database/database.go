package database

import (
	"fmt"
	"log"
	"my-orm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_Username = "root"
	DB_Password = ""
	DB_Name     = "ormdb"
	DB_Host     = "localhost"
	DB_Port     = "3306"
)

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_Username,
		DB_Password,
		DB_Host,
		DB_Port,
		DB_Name,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error when connecting to the database", err)
	}
	log.Println("connected to the database")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Package{}, &models.User{})
	if err != nil {
		log.Fatal("migration failed", err)
	}
	log.Println("migration success")

}
