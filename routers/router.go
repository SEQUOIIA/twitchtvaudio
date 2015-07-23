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
	subpath := "/"
	appRouterbase := mux.NewRouter().StrictSlash(true)


	Router.PathPrefix(subpath).Handler(negroni.New(
		negroni.Wrap(appRouterbase),
	))
	/*
	Router.Handle(subpath, negroni.New(
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			http.Redirect(w, r, subpath + "/", 301)
		})),
	))
	*/

	appRouter := appRouterbase.PathPrefix(subpath).Subrouter()

	appRouter.Path("/").Handler(negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.Root)),
	))

	appRouter.Handle("/get", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.GetChannelAlternative)),
	))

	appRouter.Handle("/{channelname}", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.GetChannel)),
	))

	appRouter.Handle("/audiobits/{channelname}.m3u8", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.GetM3U8andRewrite)),
	))

	appRouter.Handle("/stream/{channelname}.m3u8", negroni.New(
		negroni.Wrap(http.HandlerFunc(controllers.M3U8Stream)),
	))

	var csspath string
	var jspath string
	if subpath == "/" {
		csspath = "/css/"
		jspath = "/js/"
	} else {
		csspath = subpath + "/css/"
		jspath = subpath + "/js/"
	}

	cssfileServer := http.StripPrefix(csspath, http.FileServer(rice.MustFindBox("../views/css").HTTPBox()))
	appRouter.PathPrefix("/css/").Handler(cssfileServer)
	jsfileServer := http.StripPrefix(jspath, http.FileServer(rice.MustFindBox("../views/js").HTTPBox()))
	appRouter.PathPrefix("/js/").Handler(jsfileServer)
}
