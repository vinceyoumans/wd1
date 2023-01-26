package scrap

// for crawling

//var visited = make(map[string]bool)
//var lock sync.Mutex

//func StartScan() {
//	// loop through the URLsMap
//	log.Println("--------  Begining Scan AAA")
//
//	for key, val := range URLsMap {
//		log.Printf("About to test %s  out of %d", key, len(URLsMap))
//		if val.URLStatus == crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_PENDING) {
//			fmt.Println("-----------AAAA  About to Crawl ---- key ", key)
//
//			F03_modify_Pending_URL_to_Starting(key)
//			urlChildren, urlChildrenSlice, err := Fetch01(val.URLRoot)
//			if err != nil {
//				//F04
//			}
//
//			//URLsMap[key].URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_IN_PROGRESS)
//			//
//			//var wg sync.WaitGroup
//			//wg.Add(1)
//			////urls := Crawl("https://golang.org", &wg)
//			////urls := Crawl(key, &wg)
//			//urls := Crawl(val.URLRoot, &wg)
//			//
//			//wg.Wait()
//			//fmt.Println(urls)
//			//
//			////   there has to be a better way...   but I am tired, and I can't figure it out.
//			//
//			////u := val.URLRoot
//			////URLsMap[u].URLCrawled = urls
//			//
//			////for _, s:= range urls{
//			////	append(URLsMap[val.URLRoot].URLCrawled, s)
//			////}
//
//			val, err := ScanbURL(key)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			fmt.Println("Child URLs found on", key, ":")
//			for _, url := range val {
//				fmt.Println(url)
//			}
//
//			E := URLsMap[key]
//
//			E.URLRoot = key
//			E.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_FINISHED)
//			E.URLCrawled = val
//
//			URLsMap[key] = E
//
//			//URLsMap[key].URLCrawled = val
//			//URLsMap[key].URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_FINISHED)
//			//
//
//			fmt.Println("----- Value ", val)
//		} else {
//			fmt.Println("-----------BBBB  skipping Crawl ---- key ", key)
//		}
//
//	}
//	//time.Sleep(3 * time.Second)
//	//fmt.Println("-----  finished sleepting.. starting scan again")
//
//}

//func Crawl(url string, wg *sync.WaitGroup) []string {
//	defer wg.Done()
//	urls := []string{}
//	resp, err := http.Get(url)
//	if err != nil {
//		return urls
//	}
//	defer resp.Body.Close()
//
//	doc, err := html.Parse(resp.Body)
//	if err != nil {
//		return urls
//	}
//
//	var f func(*html.Node)
//	f = func(n *html.Node) {
//		if n.Type == html.ElementNode && n.Data == "a" {
//			for _, a := range n.Attr {
//				if a.Key == "href" {
//					u, err := resp.Request.URL.Parse(a.Val)
//					if err == nil && u.Host == resp.Request.URL.Host {
//						wg.Add(1)
//						go Crawl(u.String(), wg)
//						urls = append(urls, u.String())
//					}
//				}
//			}
//		}
//		for c := n.FirstChild; c != nil; c = c.NextSibling {
//			f(c)
//		}
//	}
//
//	f(doc)
//
//	return urls
//}
