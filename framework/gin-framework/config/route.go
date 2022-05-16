package config

import (
	"gin-framework/controller/api"
	"gin-framework/controller/create"
	"gin-framework/controller/eat"
	"gin-framework/controller/err"
	"gin-framework/controller/first"
	"gin-framework/controller/log"
	"gin-framework/controller/login"
	"gin-framework/controller/m"
	"gin-framework/controller/mm"
	"gin-framework/controller/one"
	"gin-framework/controller/raw"
	"gin-framework/controller/session"
	"gin-framework/controller/tran"
	"gin-framework/controller/user"
	"gin-framework/middle"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func Router(router *gin.Engine) {

	//设置html路径
	//**表示文件夹
	//*表示文件
	router.LoadHTMLGlob("template/**/*")
	//设置静态文件的路径
	//第一个url 第二个是实际的路径
	router.Static("/static", "static")

	eatRouter := router.Group("/eat")
	eat.Router(eatRouter)

	loginRouter := router.Group("/login")
	login.Router(loginRouter)

	userRouter := router.Group("/reg")
	user.Router(userRouter)

	logRouter := router.Group("/log")
	log.Router(logRouter)

	sessRouter := router.Group("/session")
	session.Router(sessRouter)

	createRouter := router.Group("/create")
	create.Router(createRouter)

	oneRouter := router.Group("/one")
	one.Router(oneRouter)

	mRouter := router.Group("/m")
	m.Router(mRouter)

	mmRouter := router.Group("/mm")
	mm.Router(mmRouter)

	firstRouter := router.Group("/first")
	first.Router(firstRouter)

	errRouter := router.Group("/err")
	err.Router(errRouter)

	tranRouter := router.Group("/tran")
	tran.Router(tranRouter)

	rawRouter := router.Group("/raw")
	raw.Router(rawRouter)

	apiRouter := router.Group("/api")
	apiRouter.Use(middle.GetCrosQues)
	api.Router(apiRouter)
}
