package main

import (
	"encoding/base64"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/OGFris/ATTBot/fetcher"
	"github.com/OGFris/ATTBot/twist"
	"github.com/animenotifier/kitsu"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	bot := anaconda.NewTwitterApiWithCredentials(os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"))

	lastFetch := new(twist.EpisodesFeed)
	fetcher.EveryFetch(func(animes *twist.AnimesFeed, episodes *twist.EpisodesFeed) {
		if lastFetch.Lastbuilddate == "" {
			lastFetch = episodes
		} else {
			for n, item := range episodes.Items {
				if item.AnimetwistID == lastFetch.Items[0].AnimetwistID {
					if n != 0 {
						for i := 0; i < n; i++ {
							episode := episodes.Items[i]

							fmt.Println("Found a new anime release:", episode.Description)

							kitsuAnime, err := kitsu.GetAnime(fmt.Sprint(episode.KitsuID))
							if err != nil {
								panic(err)
							}

							// fmt.Println("Cover link:", kitsuAnime.Attributes.PosterImage.Original)

							resp, err := http.Get(kitsuAnime.Attributes.PosterImage.Original)
							if err != nil {
								panic(err)
							}

							bytes, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								panic(err)
							}

							err = resp.Body.Close()
							if err != nil {
								panic(err)
							}

							media, err := bot.UploadMedia(base64.StdEncoding.EncodeToString(bytes))
							if err != nil {
								panic(err)
							}

							// fmt.Println("Uploaded the anime cover to twitter.")

							anime := new(twist.AnimesFeedItem)
							for _, item := range animes.Items {
								if item.AnimetwistID == episode.AnimetwistID {
									// fmt.Println("Anime is:", item.AnimeTitle)
									anime = &item
									break
								}
							}

							if anime.AnimeOngoing == 1 {
								_, err := bot.PostTweet(episode.Description,
									url.Values{
										"media_ids": {media.MediaIDString},
									})
								if err != nil {
									panic(err)
								}
							} else {
								if episode.EpisodeNumber == 1 {
									_, err := bot.PostTweet(episode.Description,
										url.Values{
											"media_ids": {media.MediaIDString},
										})
									if err != nil {
										panic(err)
									}
								}
							}
						}
					}

					break
				}
			}

			lastFetch = episodes
		}
	})
}
