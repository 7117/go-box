package m

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	{
		router.GET("/m", M)
	}
}
