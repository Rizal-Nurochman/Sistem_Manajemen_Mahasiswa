package database

import (
	"belajar-api/students"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectedDatabase()  {
	dsn := "root:mrn07052005@tcp(127.0.0.1:3306)/database_mahasiswa?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database failed to connection: ",err)
	}
	DB.AutoMigrate(&students.Students{})
	fmt.Println("Database connection succeed!")
}