package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	//"root@tcp(localhost:3306)/belajar_golang_restful_api"
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)))
	if err != nil {
		fmt.Printf("Cannot connect to databse\n")
		fmt.Printf("Connection Error :%v\n", err)
	} else {
		fmt.Println("Connect to database")
	}
	_ = DB.AutoMigrate(&User{})

}
