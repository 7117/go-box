package routers

import (
	"github.com/astaxie/beego"
	"goPractice/spider02/controllers"
)

func init() {
	beego.Router("/", &controllers.SpiderController{}, "*:CrawlMovie")
}
