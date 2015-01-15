package routers

import (
	"github.com/equoia/twitchtvaudio/controllers"
	"github.com/astaxie/beego"
	_ "github.com/yvasiyarov/newrelic_platform_go"
	"github.com/yvasiyarov/beego_gorelic"
)

func init() {
    beego.Router("/:channelname", &controllers.MainController{})
	beego.Router("/get", &controllers.MainAlternativeController{})
	beego.Router("/", &controllers.NoinputController{})
	beego.Router("/api/twitch", &controllers.ApiController{})
	beego.SetStaticPath("/css", "views/css")
	beego_gorelic.InitNewrelicAgent()
}
