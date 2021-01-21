package main

import (
	"go-gin-gorm-fx-skeleton/config"
	"go-gin-gorm-fx-skeleton/utils"
)

func main() {
	utils.LoadEnv()
	db := config.InitializeDB()
	db.Begin()
}
