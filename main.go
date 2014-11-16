package main

import (
	_ "github.com/equoia/twitchtvaudio/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.HttpPort = 8089
	beego.Run()
}

