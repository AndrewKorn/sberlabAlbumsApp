package main

import (
	"errors"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
	"time"
)

func httpGetWithTimeout(url string) (*http.Response, error) {
	timeout := time.Duration(15 * time.Second)
	c := http.Client{Timeout: timeout}
	return c.Get(url)
}

func matchTag(node *html.Node, t string) (h *html.Node, err error) {
	for h = node.FirstChild; h != nil; h = h.NextSibling {
		if h.Type == html.ElementNode && h.Data == t {
			return
		}
	}
	err = errors.New("tag not found: " + t)
	return
}

func GetSongID(url string) (songID string, err error) {
	response, err := httpGetWithTimeout(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := html.Parse(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var currentNode *html.Node
	if currentNode, err = matchTag(doc, "html"); err != nil {
		log.Fatal("Cannot find html tag")
	}
	if currentNode, err = matchTag(currentNode, "head"); err != nil {
		log.Fatal("Cannot find head tag")
	}

	for h := currentNode.FirstChild; h != nil; h = h.NextSibling {
		if h.Data == "meta" {
			var found bool
			var v string
			for _, a := range h.Attr {
				if a.Key == "property" && a.Val == "al:ios:url" {
					found = true
				}
				if a.Key == "content" {
					v = a.Val
				}
			}

			if found {
				i := strings.LastIndex(v, ":")
				songID = v[i+1:]
				break
			}
		}
	}
	return
}
