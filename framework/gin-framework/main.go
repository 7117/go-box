package main

import (
	"gin-framework/config"
	_ "gin-framework/data_source"
	_ "gin-framework/logs_source"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.Router(router)

	router.Run(":8090")
}
