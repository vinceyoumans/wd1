package main

//func ScanbURL(rootURL string) ([]string, error) {
//	// Create a timeout value for the http.Get request
//	//timeout := time.Duration(5 * time.Second)
//	client := http.Client{
//		//Timeout: timeout,
//	}
//
//	// Make the GET request to the rootURL
//	resp, err := client.Get(rootURL)
//	if err != nil {
//		return nil, fmt.Errorf("Error making GET request to %s: %s", rootURL, err)
//	}
//	defer resp.Body.Close()
//
//	// Parse the HTML from the GET request
//	doc, err := html.Parse(resp.Body)
//	if err != nil {
//		return nil, fmt.Errorf("Error parsing HTML from %s: %s", rootURL, err)
//	}
//
//	// Initialize a slice to hold the child URLs
//	childURLs := []string{}
//
//	// Define a function to recursively search the HTML tree for child URLs
//	var findLinks func(*html.Node)
//	findLinks = func(n *html.Node) {
//		if n.Type == html.ElementNode && n.Data == "a" {
//			for _, a := range n.Attr {
//				if a.Key == "href" {
//					// add code here to test if this is a child under the rootURL
//					childURLs = append(childURLs, a.Val)
//				}
//			}
//		}
//		for c := n.FirstChild; c != nil; c = c.NextSibling {
//			findLinks(c)
//		}
//	}
//
//	// Call the findLinks function to begin the search
//	findLinks(doc)
//
//	return childURLs, nil
//}
