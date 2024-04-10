package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/keegcode/franko-crawler/internal/crawler"
	"github.com/keegcode/franko-crawler/internal/telegram"
	"golang.org/x/net/html"
)

var urls []string = []string{
	"https://ft.org.ua/ua/performance/konotopska-vidma",
	"https://ft.org.ua/ua/performance/kaligula",
}

var dates sync.Map = sync.Map{}

func main() {
	apiKey := flag.String("api", "", "Telegram Bot API Key")
	channelId := flag.String("id", "", "Telegram Channel ID")

	flag.Parse()

	if *apiKey == "" || *channelId == "" {
		fmt.Println("Missing API Key or Channel ID!")
		os.Exit(1)
	}

	tg := telegram.Telegram{ApiKey: *apiKey, ChannelId: *channelId}
	wg := sync.WaitGroup{}

	for {
		for _, url := range urls {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				nodes, err := crawler.Crawl(u)
				if err != nil {
					fmt.Print(err)
					tg.SendMessage("Мені Пагано: " + err.Error())
					return
				}

				for _, n := range nodes {
					go func(nd *html.Node) {
						date := nd.FirstChild.Data
						if _, ok := dates.Load(u + date); ok {
							return
						}

						err := tg.SendMessage(u + "\n" + date)
						if err != nil {
							fmt.Print(err)
							tg.SendMessage("Мені Пагано: " + err.Error())
							return
						}

						dates.Store(u+date, true)
					}(n)
				}
			}(url)
		}
		wg.Wait()
		time.Sleep(time.Minute)
	}
}
