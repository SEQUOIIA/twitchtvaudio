package routers

import (
	"github.com/equoia/twitchtvaudio/controllers"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"net/http"
)

var Router = mux.NewRouter()

func init() {
	subpath := "/"
	appRouterbase := mux.NewRouter().StrictSlash(true)
	apirouterbase := mux.NewRouter()


	Router.PathPrefix(subpath).Handler(negroni.New(
		negroni.Wrap(appRouterbase),
	))

	appRouter := appRouterbase.PathPrefix(subpath).Subrouter()

	// API routes

	appRouter.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(controllers.Apisettings),
		negroni.Wrap(apirouterbase),
	))
	apirouter := apirouterbase.PathPrefix("/api").Subrouter()
	apirouter.Handle("/stream/{channelname}", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.GetStream)),
	))
}
