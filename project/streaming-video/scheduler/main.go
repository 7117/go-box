package scheduler

import (
	"github.com/julienschmidt/httprouter"
	"goPractice/video/scheduler/task"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videorecord/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	// 启动taskrunner
	go task.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
