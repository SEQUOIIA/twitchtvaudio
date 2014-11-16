package retrieveaudio

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bufio"
	"github.com/grafov/m3u8"
)

type Accesstoken struct {
	Token				string	`json:"token"`
	Sig					string	`json:"sig"`
	Mobile_restricted	bool	`json:"mobile_restricted"`
}

func Get(channelname string) (int, string) {
	cli := http.DefaultClient

	accessTokenurl := "http://api.twitch.tv/api/channels/" + channelname + "/access_token"
	req, err := http.NewRequest("GET", accessTokenurl, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", "letr.it/twitchaudio 1.0")

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("Looks like something went wrong with the request:\n", err)
	}
	defer resp.Body.Close()

	var accesstoken Accesstoken
	tempbody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(tempbody, &accesstoken)
	if err != nil {
		fmt.Println("Something went wrong with the JSON:\n", err)
	}

	m3u8url := "http://usher.twitch.tv/select/" + channelname + ".json?nauthsig=" + accesstoken.Sig + "&nauth=" + accesstoken.Token +"&allow_source=true&allow_audio_only=true"
	req, err = http.NewRequest("GET", m3u8url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", "letr.it/twitchaudio 1.0")

	resp, err = cli.Do(req)
	if err != nil {
		fmt.Println("Looks like something went wrong with the request:\n", err)
	}
	defer resp.Body.Close()

	//tempbody, err = ioutil.ReadAll(resp.Body)

	//allstreamsraw := string(tempbody)

	p, listType, err := m3u8.DecodeFrom(bufio.NewReader(resp.Body), true)
	if err != nil {
		panic(err)
	}
	switch listType {
	case m3u8.MEDIA:
		mediapl := p.(*m3u8.MediaPlaylist)
		fmt.Printf("%+v\n", mediapl)
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		//fmt.Printf("%+v\n", masterpl.Variants[5])
	for _, data := range masterpl.Variants {
		if data.Video == "audio_only" {
			return 1, data.URI
		}

	}
	}
	return 0, "no go"
}

func maintest() {
	cli := http.DefaultClient

	accessTokenurl := "http://api.twitch.tv/api/channels/" + "taketv" + "/access_token"
	req, err := http.NewRequest("GET", accessTokenurl, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", "letr.it/twitchaudio 1.0")

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("Looks like something went wrong with the request:\n", err)
	}
	defer resp.Body.Close()

	var accesstoken Accesstoken
	tempbody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(tempbody, &accesstoken)
	if err != nil {
		fmt.Println("Something went wrong with the JSON:\n", err)
	}

	m3u8url := "http://usher.twitch.tv/select/" + "taketv" + ".json?nauthsig=" + accesstoken.Sig + "&nauth=" + accesstoken.Token +"&allow_source=true&allow_audio_only=true"
	req, err = http.NewRequest("GET", m3u8url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", "letr.it/twitchaudio 1.0")

	resp, err = cli.Do(req)
	if err != nil {
		fmt.Println("Looks like something went wrong with the request:\n", err)
	}
	defer resp.Body.Close()

	//tempbody, err = ioutil.ReadAll(resp.Body)

	//allstreamsraw := string(tempbody)

	p, listType, err := m3u8.DecodeFrom(bufio.NewReader(resp.Body), true)
	if err != nil {
		panic(err)
	}
	switch listType {
	case m3u8.MEDIA:
		mediapl := p.(*m3u8.MediaPlaylist)
		fmt.Printf("%+v\n", mediapl)
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		//fmt.Printf("%+v\n", masterpl.Variants[5])
		for _, data := range masterpl.Variants {

			if data.Video == "audio_only" {
				fmt.Println("Audio only stream found!")
				fmt.Println(data.URI)
			} else {
				fmt.Println("Audio only stream not found :/")
			}

		}
	}
}
