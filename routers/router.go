package routers

import (
	"github.com/equoia/twitchtvaudio/controllers"
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

	Router.Handle("/get", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.GetChannelAlternative)),
	))

	Router.Handle("/{channelname}", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.GetChannel)),
	))



	cssfileServer := http.StripPrefix("/css/", http.FileServer(rice.MustFindBox("../views/css").HTTPBox()))
	Router.PathPrefix("/css/").Handler(cssfileServer)
}
