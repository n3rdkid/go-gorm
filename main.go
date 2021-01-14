package main

import (
	"fmt"
	"log"
	"os"

	"go-gorm/utils"

	"go-gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	utils.LoadEnv()
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	bike := models.Vehicle{TypeOfVehicle: "2 Wheeler", NoOfWheels: 2}
	db.Create(&bike)
	log.Printf("Connection Successful")
	if err != nil {
		panic(err.Error())
	}
}
