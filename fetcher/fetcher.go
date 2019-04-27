package fetcher

import (
	"encoding/json"
	"github.com/OGFris/ATTBot/twist"
	"io/ioutil"
	"net/http"
	"time"
)

func EveryFetch(f func(animes *twist.AnimesFeed, episodes *twist.EpisodesFeed)) {
	for {
		r, err := http.Get("https://twist.moe/feed/anime?format=json")
		if err != nil {
			panic(err)
		}

		bytes, err := ioutil.ReadAll(r.Body)

		err = r.Body.Close()
		if err != nil {
			panic(err)
		}

		animes := new(twist.AnimesFeed)
		err = json.Unmarshal(bytes, &animes)
		if err != nil {
			panic(err)
		}

		r, err = http.Get("https://twist.moe/feed/episodes?format=json")
		if err != nil {
			panic(err)
		}

		bytes, err = ioutil.ReadAll(r.Body)

		err = r.Body.Close()
		if err != nil {
			panic(err)
		}

		episodes := new(twist.EpisodesFeed)
		err = json.Unmarshal(bytes, &episodes)
		if err != nil {
			panic(err)
		}

		f(animes, episodes)
		time.Sleep(time.Minute)
	}
}
