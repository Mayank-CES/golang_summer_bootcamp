package main

import (
	"fmt"
	"studentAPI/Config"
	"studentAPI/Models"
	"studentAPI/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{}, &Models.Marks{})
	//Config.DB.AutoMigrate(&Models.Marks{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
