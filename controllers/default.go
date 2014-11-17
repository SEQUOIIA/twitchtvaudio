package controllers

import (
	"github.com/astaxie/beego"
	"github.com/equoia/twitchtvaudio/retrieveaudio"
	"github.com/equoia/twitchtvaudio/models"
)

type MainController struct {
	beego.Controller
}


type NoinputController struct {
	beego.Controller
}

type ApiController struct {
	beego.Controller
}

func (u *ApiController) Get() {
	content := models.GetAudiostream(u.Ctx.Input.Param(":channelname"))
	u.Data["json"] = content
	u.ServeJson()
}


func (this *MainController) Get() {
	this.Data["Website"] = "hummel.yt"

	statuscode, url := retrieveaudio.Get(this.Ctx.Input.Param(":channelname"))
	if statuscode == 0 {
		this.Data["channelresult"] = 0
		this.TplNames = "failure.tpl"
	} else if statuscode == 1 {
		this.Data["channelresult"] = 1
		this.Data["resulturl"] = url
		this.TplNames = "success.tpl"
	}
}

func (this *NoinputController) Get() {
	this.Data["Website"] = "hummel.yt"
	this.TplNames = "noinput.tpl"
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
