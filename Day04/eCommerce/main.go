package main

import (
	"eCommerce/Config"
	_ "eCommerce/Controllers"
	"eCommerce/Models"
	"eCommerce/Routes"
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
	Config.DB.AutoMigrate(&Models.Customer{}, &Models.Retailer{}, &Models.Product{}, &Models.Transaction{})

	r := Routes.SetupRouter()
	//Services.NewCustomerService(Repository.NewRepo())

	err := r.Run()
	if err != nil {
		return
	}
}


