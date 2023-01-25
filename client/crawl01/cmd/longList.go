/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	crawlerpb "wd1/gen/go/crawler/v1"

	"github.com/spf13/cobra"
)

// longListCmd represents the longList command
var longListCmd = &cobra.Command{
	Use:   "longList",
	Short: "longList will return all of the Crawled URLs and their returned urls that were crawled",
	Long: `longList will return either a specific URL that was crawled and its results, or it will default to 
returning all of the URLs crawled and all of their responses

if you do not inlcude a -RootURL as a parameter, then all crawled URL's will be returned'

-rootURL is URL to crawl...  defaults to vyoumans.com 
-serverURL Optional Server Address, defauts to localhost
-port  port of server,  defaults to 50051


	URLMonitorCrawls_URLCurrentStatus_name = map[int32]string{
		0: "URL_CURRENT_STATUS_UNSPECIFIED",
		1: "URL_CURRENT_STATUS_NO_RECORD",
		2: "URL_CURRENT_STATUS_PENDING",
		3: "URL_CURRENT_STATUS_START",
		4: "URL_CURRENT_STATUS_IN_PROGRESS",
		5: "URL_CURRENT_STATUS_FINISHED",
		6: "URL_CURRENT_STATUS_CANCELED",
	}
	URLMonitorCrawls_URLCurrentStatus_value = map[string]int32{
		"URL_CURRENT_STATUS_UNSPECIFIED": 0,
		"URL_CURRENT_STATUS_NO_RECORD":   1,
		"URL_CURRENT_STATUS_PENDING":     2,
		"URL_CURRENT_STATUS_START":       3,
		"URL_CURRENT_STATUS_IN_PROGRESS": 4,
		"URL_CURRENT_STATUS_FINISHED":    5,
		"URL_CURRENT_STATUS_CANCELED":    6,
	}
)

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("longList called")
		urlRoot, _ := cmd.Flags().GetString("rootlURL")
		ServerlURL, _ := cmd.Flags().GetString("serverlURL")
		Port, _ := cmd.Flags().GetString("port")

		runLongList(urlRoot, ServerlURL, Port)

	},
}

func init() {
	rootCmd.AddCommand(longListCmd)

	longListCmd.Flags().String("rootURL", "vyoumans.com", "The URL to be Crawled, defaults to vyoumans.com")
	longListCmd.Flags().String("serverURL", "localhost", "Optional Server Address, defauts to localhost")
	longListCmd.Flags().String("port", "50051", "port of server,  defaults to 50051")
}

func runLongList(urlRoot string, ServerlURL string, port string) {

	target := fmt.Sprintf("%s:%s", ServerlURL, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	log.Println("======== client longList trying to connect to :", target)

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := crawlerpb.NewCrawlerServiceClient(conn)

	// Prepare the request
	//listReq := &crawlerpb.ListCrawlRequest{

	//if example.Field != nil {
	//	// field is set
	//} else {
	//	// field is not set
	//}

	//urlRootb := &urlRoot
	//if urlRoot != nil {
	//	// field is set
	//	urlRootb = ""
	//} else {
	//	// field is not set
	//}

	LonglistReq := &crawlerpb.LongListCrawlRequest{
		RootUrl: urlRoot,
	}

	// Call the non-blocking StartCrawl method
	//listRes, err := client.ListCrawl(context.Background(), listReq)
	LongListRes, err := client.LongListCrawl(context.Background(), LonglistReq)

	if err != nil {
		log.Fatalf("Failed to list crawl: %v", err)
	}
	log.Printf("StartCrawl response: %s", LongListRes.UrlCrawled)

	for x, res := range LongListRes.UrlCrawled {
		log.Println("=== ", x, "===", res.UrlRoot, "===", res.UrlStatus)

		for xx, resUrlCrawld := range res.UrlCrawled {
			log.Println("============= ", xx, "===", resUrlCrawld)
		}
	}

	log.Println("===========   doing it the easy PrettyPrint way ===========")

	b, err := json.MarshalIndent(LongListRes, "", "  ")
	log.Println(string(b))

}
