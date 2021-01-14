package main

import (
	"fmt"
	"log"
	"os"

	"go-gorm/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	utils.LoadEnv()
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	log.Print(fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
