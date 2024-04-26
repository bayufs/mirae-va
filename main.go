package main

import (
	"fmt"
	"mirae-va/app/routers"

	"mirae-va/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {

	defer config.DisconnectDB(db)
	r := gin.Default()
	routers.InitRouter(r)
	r.Run(fmt.Sprintf("%s:%s", config.SERVICE_HOST, config.SERVICE_PORT))
}
