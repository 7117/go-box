package one

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	{
		router.GET("/one", One)
	}
}
