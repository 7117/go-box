package api

import (
	"github.com/gin-gonic/gin"
)

func Router(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/axios", Create)
	apiRouter.GET("/detail", Detail)
}
