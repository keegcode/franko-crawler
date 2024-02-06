package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var urls []string = []string{
	"http://ft.org.ua/ua/performance/konotopska-vidma",
	"http://ft.org.ua/ua/performance/kaligula",
}

func filter(node *html.Node, class string, nodes []*html.Node) []*html.Node {
	if node == nil {
		return nodes
	}

	if node.FirstChild != nil {
		filter(node.FirstChild, class, nodes)
	}

	return filter(node.NextSibling, class, nodes)
}

func main() {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		node, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		filter(node.FirstChild, "d", []*html.Node{})
	}
}
