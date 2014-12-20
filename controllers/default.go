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
	lowercasedchannelname := strings.ToLower(this.Ctx.Input.Param(":channelname"))
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		this.Data["channelresult"] = 0
		this.TplNames = "failure.tpl"
	} else if statuscode == 1 {
		this.Data["channelresult"] = 1
		this.Data["resulturl"] = url
		this.TplNames = "success.tpl"
	}
	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Version"] = string(currentcommit)
}

func (this *MainAlternativeController) Get() {
	lowercasedchannelname := strings.ToLower(this.GetString("channelname"))
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		this.Data["channelresult"] = 0
		this.TplNames = "failure.tpl"
	} else if statuscode == 1 {
		this.Data["channelresult"] = 1
		this.Data["resulturl"] = url
		this.TplNames = "success.tpl"
	}
	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Version"] = string(currentcommit)
}


func (this *NoinputController) Get() {
	this.TplNames = "noinput.tpl"
	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Version"] = string(currentcommit)
}
