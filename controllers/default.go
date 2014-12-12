package controllers

import (
	"github.com/astaxie/beego"
	"github.com/equoia/twitchtvaudio/retrieveaudio"
	"github.com/equoia/twitchtvaudio/models"
	"os/exec"
	"fmt"
	"strings"
)

type MainController struct {
	beego.Controller
}

type MainAlternativeController struct {
	beego.Controller
}

type NoinputController struct {
	beego.Controller
}

type ApiController struct {
	beego.Controller
}

type Channelnamepost struct {
	Channelname	string	`json:"channelname"`
}

/*
func (u *ApiController) Post() {
	var channelnamepost Channelnamepost
	//var channelnamepost map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &channelnamepost)
	content := models.GetAudiostream(channelnamepost.Channelname)
	u.Data["json"] = content
	u.ServeJson()
}
*/

func (u *ApiController) Post() {
	//var channelnamepost map[string]interface{}
	content := models.GetAudiostream(u.GetString("channelname"))
	u.Data["json"] = content
	u.ServeJson()
}

func (this *MainController) Get() {
	this.Data["Website"] = "hummel.yt"
	lowercasedchannelname := strings.ToLower(this.Ctx.Input.Param(":channelname"))
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		this.Data["channelresult"] = "Can't find audio-only stream"
		this.TplNames = "failure.tpl"
	} else if statuscode == 1 {
		this.Data["channelresult"] = 1
		this.Data["resulturl"] = url
		this.TplNames = "success.tpl"
	} else if statuscode == 2 {
		this.Data["channelresult"] = "Something is wrong with the Twitch API server"
		this.TplNames = "failure.tpl"
	}
	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Version"] = string(currentcommit)
}

func (this *MainAlternativeController) Get() {
	this.Data["Website"] = "hummel.yt"
	lowercasedchannelname := strings.ToLower(this.GetString("channelname"))
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		this.Data["channelresult"] = "Can't find audio-only stream"
		this.TplNames = "failure.tpl"
	} else if statuscode == 1 {
		this.Data["channelresult"] = 1
		this.Data["resulturl"] = url
		this.TplNames = "success.tpl"
	} else if statuscode == 2 {
		this.Data["channelresult"] = "Something is wrong with the Twitch API server"
		this.TplNames = "failure.tpl"
	}
	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Version"] = string(currentcommit)
}


func (this *NoinputController) Get() {
	this.Data["Website"] = "hummel.yt"
	this.TplNames = "noinput.tpl"
	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Version"] = string(currentcommit)
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
