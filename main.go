package main

import (
	"first/src/first-api/Config"
	models "first/src/first-api/Model"
	"first/src/first-api/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	//	Config.DB.DropTableIfExists(&Model.User{}, &Model.Account{})
	Config.DB.AutoMigrate(&models.User{})
	r := Routes.SetUpRouter()

	r.Run()
}
