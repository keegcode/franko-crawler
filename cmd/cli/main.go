package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/keegcode/franko-crawler/internal/crawler"
	"github.com/keegcode/franko-crawler/internal/telegram"
)

var urls []string = []string{
	"http://ft.org.ua/ua/performance/konotopska-vidma",
	"http://ft.org.ua/ua/performance/kaligula",
}

var dates map[string]bool = map[string]bool{}

func main() {
	apiKey := flag.String("api", "", "Telegram Bot API Key")
	channelId := flag.String("id", "", "Telegram Channel ID")

	flag.Parse()

	tg := telegram.Telegram{ApiKey: *apiKey, ChannelId: *channelId}

	for {
		for _, url := range urls {
			nodes := crawler.Crawl(url)

			for _, n := range nodes {
				date := n.FirstChild.Data
				if dates[url+date] {
					continue
				}

				err := tg.SendMessage(url + "\n" + date)
				if err != nil {
					fmt.Print(err)
					os.Exit(1)
				}

				dates[url+date] = true
			}
		}
		time.Sleep(time.Minute)
	}
}
