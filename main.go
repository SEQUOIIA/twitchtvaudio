package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Accesstoken struct {
	Token				string	`json:"token"`
	Sig					string	`json:"sig"`
	Mobile_restricted	bool	`json:"mobile_restricted"`
}

func main() {
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

	fmt.Println(accesstoken)

	

}
