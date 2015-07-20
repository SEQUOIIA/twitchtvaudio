package controllers

import (
	"github.com/equoia/twitchtvaudio/retrieveaudio"
	"os/exec"
	"fmt"
	"strings"
	"net/http"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"github.com/gorilla/mux"
)

func HttpHeaderSet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Server", "github.com/equoia/twitchtvaudio 1.0")
	next(w, r)
}

func Root(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	ctx["version"] = string(currentcommit)

	templateBox, err := rice.FindBox("../views")
	if err != nil {
		panic(err)
	}

	templateString, err := templateBox.String("noinput.tpl")
	if err != nil {
		panic(err)
	}

	htmlTemplate, err := template.New("index").Parse(templateString)
	if err != nil {
		panic(err)
	}

	htmlTemplate.Execute(w, ctx)
}

func GetChannel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := make(map[string]interface{})

	templateBox, err := rice.FindBox("../views")
	if err != nil {
		panic(err)
	}

	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	ctx["Version"] = string(currentcommit)

	lowercasedchannelname := strings.ToLower(vars["channelname"])
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		ctx["channelresult"] = 0
		templateString, err := templateBox.String("failure.tpl")
		if err != nil {
			panic(err)
		}

		htmlTemplate, err := template.New("index").Parse(templateString)
		if err != nil {
			panic(err)
		}

		htmlTemplate.Execute(w, ctx)
	} else if statuscode == 1 {
		ctx["channelresult"] = 1
		ctx["resulturl"] = url
		templateString, err := templateBox.String("success.tpl")
		if err != nil {
			panic(err)
		}

		htmlTemplate, err := template.New("index").Parse(templateString)
		if err != nil {
			panic(err)
		}

		htmlTemplate.Execute(w, ctx)
	}
}

func GetChannelAlternative(w http.ResponseWriter, r *http.Request) {
	channelname := r.URL.Query().Get("channelname")

	ctx := make(map[string]interface{})

	templateBox, err := rice.FindBox("../views")
	if err != nil {
		panic(err)
	}

	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	ctx["Version"] = string(currentcommit)

	lowercasedchannelname := strings.ToLower(channelname)
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		ctx["channelresult"] = 0
		templateString, err := templateBox.String("failure.tpl")
		if err != nil {
			panic(err)
		}

		htmlTemplate, err := template.New("index").Parse(templateString)
		if err != nil {
			panic(err)
		}

		htmlTemplate.Execute(w, ctx)
	} else if statuscode == 1 {
		ctx["channelresult"] = 1
		ctx["resulturl"] = url
		templateString, err := templateBox.String("success.tpl")
		if err != nil {
			panic(err)
		}

		htmlTemplate, err := template.New("index").Parse(templateString)
		if err != nil {
			panic(err)
		}

		htmlTemplate.Execute(w, ctx)
	}
}