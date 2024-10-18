package main

import (
	"fmt"
	"go-demo/config"
)

func main() {
	config.InitConfig()
	fmt.Println(config.AppConfig.Database.User)
	fmt.Println(config.AppConfig.App.Name)
}
