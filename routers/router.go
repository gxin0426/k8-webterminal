package routers

import (
	"k8-webterminal/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/terminal", &controllers.TerminalController{}, "get:Get")
	beego.Handler("/terminal/ws", &controllers.TerminalSockjs{}, true)
	beego.Router("/podlog", &controllers.PodLogController{}, "get:Get")
	beego.Handler("/podlog/ws", &controllers.PodLogSockjs{}, true)
	beego.Router("/searchpod", &controllers.SearchPodController{}, "get:Get")
}
