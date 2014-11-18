package routers

import (
	"github.com/equoia/twitchtvaudio/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/:channelname", &controllers.MainController{})
	beego.Router("/get", &controllers.MainAlternativeController{})
	beego.Router("/", &controllers.NoinputController{})
	beego.Router("/api/twitch", &controllers.ApiController{})
	beego.SetStaticPath("/css", "views/css")
}
