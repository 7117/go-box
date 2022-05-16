package eat

import (
	"github.com/gin-gonic/gin"
)

func Router(eatRouter *gin.RouterGroup) {
	{
		eatRouter.GET("/world", Hello)
	}
}
