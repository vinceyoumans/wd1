package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
	"time"
	crawlerpb "wd1/gen/go/crawler/v1"
)

type backgroundService struct{}

func (s *backgroundService) Run(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("background service stopping")
			return
		default:
			fmt.Println("background service running")

			for key, val := range URLsMap {
				if val.URLStatus == crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_PENDING {
					log.Printf("============   starting Fetch for %s", key)
					F11ModifyPendingToStarting(key)
					Fetch01(key) // blocking until all url children are done.  use go routine to unblock
				} else {
					log.Printf("=====   skipping URL %s", key)
				}
			}

			time.Sleep(4 * time.Second)
		}
	}
}

func Fetch01(urlRoot string) ([]byte, []string, error) {

	F21ModifyStartToInProgress(urlRoot)

	res, err := http.Get(urlRoot)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var childUrls []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					//childUrls = append(childUrls, a.Val)
					log.Printf(string("---------   in scan   a.Val = %s  -- urlRoot = %s"), a.Val, urlRoot)
					if strings.HasPrefix(a.Val, urlRoot) {
						childUrls = append(childUrls, a.Val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	fetchupdateUrlsmap(urlRoot, childUrls)

	b, _ := json.Marshal(childUrls)
	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	return out.Bytes(), childUrls, nil
}

func fetchupdateUrlsmap(url string, urls []string) {
	//URLCrawl
	log.Printf("--------- f20   changing start to in progress - %s", url)
	var urlsTemp URLCrawl
	urlsTemp.URLRoot = url
	urlsTemp.URLCrawled = urls
	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_FINISHED)
	URLsMap[url] = urlsTemp
}

//func main() {
//	jsonData, err := Fetch("https://www.wix.com")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(jsonData))
//}
