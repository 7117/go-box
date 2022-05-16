package routers

import (
	"github.com/astaxie/beego"
	"spider/controllers"
)

func init() {

	beego.Router("/", &controllers.CrawlController{}, "*:Crawl")

}
