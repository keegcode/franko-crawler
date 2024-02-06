package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"golang.org/x/net/html"
)

var urls []string = []string{
	"http://ft.org.ua/ua/performance/konotopska-vidma",
	"http://ft.org.ua/ua/performance/kaligula",
}

type PerfomanceInfo struct {
	url  string
	date string
}

var dates map[string]bool = map[string]bool{}

func filter(node *html.Node, class string, nodes *[]*html.Node) []*html.Node {
	if node == nil {
		return *nodes
	}

	if node.Data == "div" {
		for _, a := range node.Attr {
			if a.Key == "class" && a.Val == class {
				*nodes = append(*nodes, node)
				break
			}
		}
	}

	if node.FirstChild != nil {
		filter(node.FirstChild, class, nodes)
	}

	return filter(node.NextSibling, class, nodes)
}

func crawl() []PerfomanceInfo {
	info := []PerfomanceInfo{}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		root, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		nodes := filter(root.FirstChild, "performanceevents_item_info_date", &[]*html.Node{})

		for _, n := range nodes {
			date := n.FirstChild.Data
			info = append(info, PerfomanceInfo{url, date})
		}
	}

	return info
}

func main() {
	for {
		perfomances := crawl()
		new := 0

		for _, p := range perfomances {
			if dates[p.url+p.date] {
				continue
			}
			new += 1
			fmt.Printf("New Anouncement:\nShow: %s\nDate: %s\n\n", path.Base(p.url), p.date)
			dates[p.url+p.date] = true
		}

		if new == 0 {
			fmt.Println("No new dates has arrived!")
		}

		time.Sleep(time.Second * 10)
	}
}
