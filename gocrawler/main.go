package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	log "github.com/llimllib/loglevel"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type link struct {
	url   string
	text  string
	depth int
}

type httpError struct {
	original string
}

// MaxDepth const
const MaxDepth = 3

func linkReader(resp *http.Response, depth int) []link {
	page := html.NewTokenizer(resp.Body)
	links := []link{}

	var start *html.Token
	var text string

	for {
		_ = page.Next()
		token := page.Token()
		if token.Type == html.ErrorToken {
			break
		}

		if start != nil && token.Type == html.TextToken {
			text = fmt.Sprintf("%s%s", text, token.Data)
		}

		if token.DataAtom == atom.A {
			switch token.Type {
			case html.StartTagToken:
				if len(token.Attr) > 0 {
					start = &token
				}
			case html.EndTagToken:
				if start == nil {
					log.Warnf("Link end found without start: %s", text)
					continue
				}
				link := newLink(*start, text, depth)
				if link.Valid() {
					links = append(links, link)
					log.Debug("Link Found:%s", link)
				}

				start = nil
				text = ""
			}
		}
	}

	log.Debug(links)
	return links
}

func newLink(tag html.Token, text string, depth int) link {
	link := link{text: strings.TrimSpace(text), depth: depth}

	for i := range tag.Attr {
		if tag.Attr[i].Key == "href" {
			link.url = strings.TrimSpace(tag.Attr[i].Val)
		}
	}

	return link
}

func (s link) String() string {
	spacer := strings.Repeat("\t", s.depth)
	return fmt.Sprintf("%s%s (%d) - %s", spacer, s.text, s.depth, s.url)
}

func (s link) Valid() bool {
	if s.depth > MaxDepth {
		return false
	}

	if len(s.text) == 0 {
		return false
	}

	if len(s.url) == 0 || strings.Contains(strings.ToLower(s.url), "javascript") {
		return false
	}

	return true
}

func (s httpError) Error() string {
	return s.original
}

var wg = sync.WaitGroup{}

func recurDownloader(url string, depth int) {
	page, err := downloader(url)
	if err != nil {
		log.Error(err)
		return
	}
	links := linkReader(page, depth)

	for _, l := range links {
		fmt.Println(l)
		if depth+1 < MaxDepth {
			go func(l link) {
				wg.Add(1)
				recurDownloader(l.url, depth+1)
				wg.Done()
			}(l)
		}
	}
	wg.Wait()
}

func downloader(url string) (resp *http.Response, err error) {
	log.Debugf("Downloading %s", url)
	resp, err = http.Get(url)
	if err != nil {
		log.Debugf("Error: %s", err)
		return
	}

	if resp.StatusCode > 299 {
		err = httpError{fmt.Sprintf("Error (%d): %s", resp.StatusCode, url)}
		log.Debugf("Error: %s", err)
		return
	}

	return resp, nil
}

func main() {
	log.SetPriorityString("info")
	log.SetPrefix("crawler")

	log.Debug(os.Args)

	if len(os.Args) < 2 {
		log.Fatalln("Missing Url arg")
	}

	recurDownloader(os.Args[1], 0)
}
