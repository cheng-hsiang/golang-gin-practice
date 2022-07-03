package main

import (
	"fmt"
	"gin_api/common"
	"gin_api/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	fmt.Println(db)
	r := gin.Default()
	r = router.CollectRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")       // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(workDir + "/config") // path to look for the config file in
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
