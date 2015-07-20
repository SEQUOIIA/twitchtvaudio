package routers

import (
	"github.com/equoia/twitchtvaudio/controllers"
	//"github.com/astaxie/beego"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/GeertJohan/go.rice"
)

var Router = mux.NewRouter()

func init() {
	Router.Handle("/", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.Root)),
	))


	cssfileServer := http.StripPrefix("/css/", http.FileServer(rice.MustFindBox("../views/css").HTTPBox()))
	Router.PathPrefix("/css/").Handler(cssfileServer)

	/*
    beego.Router("/:channelname", &controllers.MainController{})
	beego.Router("/get", &controllers.MainAlternativeController{})
	beego.Router("/", &controllers.NoinputController{})
	beego.Router("/api/twitch", &controllers.ApiController{})
	beego.SetStaticPath("/css", "views/css")
	*/
}
