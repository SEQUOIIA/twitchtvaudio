package controllers

import (
	"github.com/astaxie/beego"
	"github.com/equoia/twitchtvaudio/retrieveaudio"
)

type MainController struct {
	beego.Controller
}


type NoinputController struct {
	beego.Controller
}


func (this *MainController) Get() {
	this.Data["Website"] = "hummel.yt"

	statuscode, url := retrieveaudio.Get(this.Ctx.Input.Param(":channelname"))
	if statuscode == 0 {
		this.Data["channelresult"] = 0
		this.TplNames = "failure.tpl"
		this.Layout = "failure.tpl"
	} else if statuscode == 1 {
		this.Data["channelresult"] = 1
		this.Data["resulturl"] = url
		this.TplNames = "success.tpl"
		this.Layout = "success.tpl"
	}
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Gascript"] = "ga_script.tpl"
}

func (this *NoinputController) Get() {
	this.Data["Website"] = "hummel.yt"
	this.Layout = "noinput.tpl"
	this.TplNames = "noinput.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Gascript"] = "ga_script.tpl"
}


/*
func (this *CvrController) Get() {
	this.TplNames = "cvr.tpl"
	this.Data["Website"] = "hummel.yt"

	virksomheddata := cvr.Get(this.Ctx.Input.Param(":cvr_nr"))
	var virksomhed Virksomhed
	err := json.Unmarshal(virksomheddata, &virksomhed)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	this.Data["virksomhed"] = virksomhed
}


func (this *TwitchController) Get() {
	this.TplNames = "twitch.tpl"
	this.Data["Website"] = "hummel.yt"

	url := "https://api.twitch.tv/api/videos/a" + this.Ctx.Input.Param(":vod_id") + ".json"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data Twitchvodjson
	json.Unmarshal(body, &data)

	vodurls := ""

	for _, vod := range data.Chunks.Live.Url {
		fmt.Println(vod)
		//vodurls = +vod
	}

	this.Data["Channel"] = vodurls

}
*/
