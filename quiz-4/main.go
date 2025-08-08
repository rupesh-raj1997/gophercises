package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	file, err := os.Open("./ex1.html")

	if err != nil {
		fmt.Println("File reading error exiting program")
	}

	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println("File parsing error exiting program")
	}

	var links []Link
	dfs(doc, &links)
	fmt.Println(links)
}

func getHref(n *html.Node) string {
	var href string
	for _, attr := range n.Attr {
		key, val := attr.Key, attr.Val
		if key == "href" {
			href = val
		}
	}
	return href
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var result string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += getText(c)
	}
	return result
}

func dfs(n *html.Node, l *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		*l = append(*l, Link{Href: getHref(n), Text: getText(n)})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, l)
	}
}
