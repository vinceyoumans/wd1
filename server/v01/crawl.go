package main

//const (
//	URLCurrentCrawls_URL_CURRENT_STATUS_UNSPECIFIED URLCurrentCrawls_URLCurrentStatus = 0
//	URLCurrentCrawls_URL_CURRENT_STATUS_NO_RECORD   URLCurrentCrawls_URLCurrentStatus = 1
//	URLCurrentCrawls_URL_CURRENT_STATUS_PENDING     URLCurrentCrawls_URLCurrentStatus = 2 // In List waiting for Start Que
//	URLCurrentCrawls_URL_CURRENT_STATUS_START       URLCurrentCrawls_URLCurrentStatus = 3 // Tagged for Crawling
//	URLCurrentCrawls_URL_CURRENT_STATUS_IN_PROGRESS URLCurrentCrawls_URLCurrentStatus = 4
//	URLCurrentCrawls_URL_CURRENT_STATUS_FINISHED    URLCurrentCrawls_URLCurrentStatus = 5
//	URLCurrentCrawls_URL_CURRENT_STATUS_CANCELED    URLCurrentCrawls_URLCurrentStatus = 6
//)

//var URLsMap = make(map[string]URLCrawl)

// var visited = make(map[string]bool)
//var initialURL *url.URL
//var wg sync.WaitGroup

//func init() {
//
//}

//// ================================================================
//func F03_modify_Pending_URL_to_Starting(url string) {
//	var urlsTemp URLCrawl
//	urlsTemp.URLRoot = url
//	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_START)
//
//	// debating how to keep track of the jobs
//	//URLCrawls.URLCrawl = append(URLCrawls.URLCrawl, urlsTemp)
//	URLsMap[url] = urlsTemp
//}

//// ================================================================
//func F04_begin_URL_Crawl(url string) string {
//	if URLsMap[url].URLStatus == crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_CANCELED) {
//		return "*** URL is in Cancel mode...   pleaase restart if you want to crawl"
//	}
//
//	var urlsTemp URLCrawl
//	urlsTemp.URLRoot = url
//	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_IN_PROGRESS)
//
//	// debating how to keep track of the jobs
//	//URLCrawls.URLCrawl = append(URLCrawls.URLCrawl, urlsTemp)
//	URLsMap[url] = urlsTemp
//
//	//urlsTemp.URLCrawled = doCrawlB(url)
//	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_FINISHED)
//
//	return "*** URL CrawlFinished"
//}

// ===========================
//var ll []string

//func doCrawlB(Surl string) []string {
//	ll := []string{}
//
//	fmt.Println("----   starting doCrawlb - ", Surl)
//	if visited[Surl] {
//		ll = append(ll, Surl)
//		return ll
//	}
//	visited[Surl] = true
//	//fmt.Println(Surl)
//
//	response, err := http.Get(Surl)
//	if err != nil {
//		ll = append(ll, Surl)
//		return ll
//	}
//	defer response.Body.Close()
//
//	tokenizer := html.NewTokenizer(response.Body)
//	for {
//		tokenType := tokenizer.Next()
//		if tokenType == html.ErrorToken {
//			ll = append(ll, Surl)
//			return ll
//		}
//		if tokenType == html.StartTagToken || tokenType == html.EndTagToken {
//			token := tokenizer.Token()
//			if token.Data == "a" {
//				for _, attribute := range token.Attr {
//					if attribute.Key == "href" {
//						link := attribute.Val
//						if strings.HasPrefix(link, "http") {
//							u, err := url.Parse(link)
//							if err != nil {
//								continue
//							}
//							if u.Hostname() == initialURL.Hostname() {
//								ll = append(ll, link)
//								go doCrawlB(link) //  NOT recursive and missing some URL's
//								//go doCrawl(link)
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//}

//==================================
