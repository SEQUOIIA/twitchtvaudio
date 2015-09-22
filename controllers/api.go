package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strings"
	"github.com/equoia/twitchtvaudio/retrieveaudio"
	"encoding/json"
)

func Apisettings(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Twitch audio-only legacy")
	w.Header().Set("Access-Control-Allow-Origin", "http://equoia.github.io")
	next(w, r)
}

func GetStream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lowercasedchannelname := strings.ToLower(vars["channelname"])
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		http.Error(w, "", http.StatusNotFound)
	} else if statuscode == 1 {
		type result struct {
			Url string `json:"url"`
		}
		tmpresult := result{Url: url}

		tmpresultJson, err := json.Marshal(tmpresult)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(tmpresultJson)
	}
}