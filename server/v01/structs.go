package main

import crawlerpb "wd1/gen/go/crawler/v1"

const (
	URLCurrentCrawls_URL_CURRENT_STATUS_UNSPECIFIED crawlerpb.URLCurrentCrawls_URLCurrentStatus = 0
	URLCurrentCrawls_URL_CURRENT_STATUS_NO_RECORD   crawlerpb.URLCurrentCrawls_URLCurrentStatus = 1
	URLCurrentCrawls_URL_CURRENT_STATUS_PENDING     crawlerpb.URLCurrentCrawls_URLCurrentStatus = 2 // In List waiting for Start Que
	URLCurrentCrawls_URL_CURRENT_STATUS_START       crawlerpb.URLCurrentCrawls_URLCurrentStatus = 3 // Tagged for Crawling
	URLCurrentCrawls_URL_CURRENT_STATUS_IN_PROGRESS crawlerpb.URLCurrentCrawls_URLCurrentStatus = 4
	URLCurrentCrawls_URL_CURRENT_STATUS_FINISHED    crawlerpb.URLCurrentCrawls_URLCurrentStatus = 5
	URLCurrentCrawls_URL_CURRENT_STATUS_CANCELED    crawlerpb.URLCurrentCrawls_URLCurrentStatus = 6
)

type URLCrawl struct {
	URLRoot    string
	URLStatus  crawlerpb.URLCurrentCrawls_URLCurrentStatus
	URLCrawled []string
}
