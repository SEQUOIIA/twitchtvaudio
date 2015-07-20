package main

import (
	"github.com/codegangsta/negroni"
	"github.com/equoia/twitchtvaudio/routers"
	"github.com/equoia/twitchtvaudio/controllers"
)

func main() {
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(controllers.HttpHeaderSet))
	n.UseHandler(routers.Router)
	n.Run("0.0.0.0:8089")
}

