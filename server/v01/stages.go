package main

import (
	"log"
	crawlerpb "wd1/gen/go/crawler/v1"
)

// assorted helper functions

// ================================================================
func F10AddPendingUrlToCrawlList(url string) {
	log.Printf("--------- f10   adding to Pending jobs %s", url)
	var urlsTemp URLCrawl
	urlsTemp.URLRoot = url
	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_PENDING)
	URLsMap[url] = urlsTemp
}

// ================================================================
func F11ModifyPendingToStarting(url string) {
	log.Printf("--------- f11   changing Pending to Start - %s", url)
	var urlsTemp URLCrawl
	urlsTemp.URLRoot = url
	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_START)
	URLsMap[url] = urlsTemp
}

// ================================================================
func F20_cancel_URL_in_Crawl_List(url string) {
	var urlsTemp URLCrawl
	urlsTemp.URLRoot = url
	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_CANCELED)
	URLsMap[url] = urlsTemp
}

// ================================================================
func F21ModifyStartToInProgress(url string) {
	log.Printf("--------- f20   changing start to in progress - %s", url)
	var urlsTemp URLCrawl
	urlsTemp.URLRoot = url
	urlsTemp.URLStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_IN_PROGRESS)
	URLsMap[url] = urlsTemp
}
