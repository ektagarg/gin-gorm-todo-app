package main

import (
	"./Config"
	"./Models"
	"./Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err := gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
