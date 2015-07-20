package main

import (
	"github.com/codegangsta/negroni"
	"github.com/equoia/twitchtvaudio/routers"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(routers.Router)
	n.Run("0.0.0.0:8089")
}

