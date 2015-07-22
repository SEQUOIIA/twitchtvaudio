package controllers

import (
	"github.com/equoia/twitchtvaudio/retrieveaudio"
	"strings"
	"net/http"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"github.com/gorilla/mux"
	"log"
	"io/ioutil"
	"regexp"
)

var Version string

func HttpHeaderSet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Server", "github.com/equoia/twitchtvaudio 1.0")
	next(w, r)
}

func Root(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})
	ctx["version"] = Version

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

	ctx["version"] = Version

	lowercasedchannelname := strings.ToLower(vars["channelname"])
	ctx["channel"] = lowercasedchannelname
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
	
	ctx["version"] = Version

	lowercasedchannelname := strings.ToLower(channelname)
	ctx["channel"] = lowercasedchannelname
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

func GetM3U8andRewrite(w http.ResponseWriter, r *http.Request) {
	// Dirty CORS workaround, does not work too well right now
	// because CORS is also being applied to the .ts files
	// found in the .m3u8

	vars := mux.Vars(r)

	lowercasedchannelname := strings.ToLower(vars["channelname"])
	statuscode, url := retrieveaudio.Get(lowercasedchannelname)
	if statuscode == 0 {
		w.WriteHeader(404)
		w.Write([]byte("Stream not found."))
	} else if statuscode == 1 {
		cli := http.DefaultClient
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Something went wrong, contact developer at @equoia"))
			log.Println(err)
		}
		req.Header.Set("User-Agent", "letr.it/twitchaudio 1.0")

		resp, err := cli.Do(req)
		if err != nil {
			log.Println("Looks like something went wrong with the request:\n", err)
		}
		defer resp.Body.Close()
		tmpbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		lines := strings.Split(string(tmpbody), "\n")
		r := regexp.MustCompile(`(.*ttvnw.net\/.*\/audio_only\/)`)

		vodfileEndpoint := r.FindStringSubmatch(url)[1]

		for i, line := range lines {
			if strings.Contains(line, "index") {
				lines[i] = vodfileEndpoint + line
			}
		}

		output := strings.Join(lines, "\n")

		w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
		w.WriteHeader(200)
		w.Write([]byte(output))

	}
}

func M3U8Stream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var tmpbody string = `#EXTM3U
#EXT-X-MEDIA:TYPE=VIDEO,GROUP-ID="audio_only",NAME="Audio Only",AUTOSELECT=YES,DEFAULT=YES
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=128000,CODECS="mp4a.40.2",VIDEO="audio_only"`

	lines := strings.Split(string(tmpbody), "\n")
	url := "https://letr.it/twitchaudio/audiobits/" + vars["channelname"] + ".m3u8"
	lines = append(lines, url)

	output := strings.Join(lines, "\n")

	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	w.WriteHeader(200)
	w.Write([]byte(output))
}