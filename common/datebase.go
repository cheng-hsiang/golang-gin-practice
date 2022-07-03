package common

import (
	"fmt"
	"gin_api/model"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	port := viper.GetString("datasource.port")
	timezone := viper.GetString("datasource.timezone")

	result := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, username, password, database, port, timezone)
	fmt.Println(result)
	fmt.Println("host=localhost user=mo password=password dbname=ec_api_v1 port=5432 sslmode=disable TimeZone=Asia/Taipei")
	// dsn := "host=localhost user=mo password=password dbname=ec_api_v1 port=5432 sslmode=disable TimeZone=Asia/Taipei"
	dsn := result
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
