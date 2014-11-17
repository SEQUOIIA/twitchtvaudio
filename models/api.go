package models

import (
	"github.com/equoia/twitchtvaudio/retrieveaudio"
)

type Twitchaudio struct {
	Statuscode int
	Url		   string
}

func GetAudiostream(channelname string) Twitchaudio {
	statuscode, url := retrieveaudio.Get(channelname)
	result := Twitchaudio{Statuscode: statuscode, Url: url}
	resultlist := make(map[string]*Twitchaudio)
	resultlist["audio_only"] = &result

	return result
}

/*func GetAudiostream(channelname string) map[string]*Twitchaudio {
	statuscode, url := retrieveaudio.Get(channelname)
	result := Twitchaudio{Statuscode: statuscode, Url: url}
	resultlist := make(map[string]*Twitchaudio)
	resultlist["audio_only"] = &result

	return resultlist
}
*/
