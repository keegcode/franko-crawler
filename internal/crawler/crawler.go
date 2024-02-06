package crawler

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

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

func Crawl(url string) []*html.Node {
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

	return filter(root.FirstChild, "performanceevents_item_info_date", &[]*html.Node{})
}
