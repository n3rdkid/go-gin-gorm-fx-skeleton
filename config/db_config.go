package config

import (
	"fmt"
	"go-gin-gorm-fx-skeleton/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitializeDB -> Initialize DB connection
func InitializeDB() *gorm.DB {
	get := utils.GetEnvByKey
	DBUSER := get("DBUSER")
	PASSWORD := get("PASSWORD")
	HOST := get("HOST")
	DATABASE := get("DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DBUSER,
		PASSWORD,
		HOST,
		3306,
		DATABASE)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + DATABASE + ";")
	if err != nil {
		// panic(err.Error())
	}
	return database
}
