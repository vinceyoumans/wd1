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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List returns all of the URL's that have been crawled",
	Long: `List returns short list of URLs that have been crawled and their status.
another command, longList will display all of the returned values

No Root URL is required because all URL's will be returned'

-serverURL Optional Server Address, defauts to localhost
-port  port of server,  defaults to 50051
=================

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
		fmt.Println("list called")

		//urlRoot, _ := cmd.Flags().GetString("rootlURL")
		ServerlURL, _ := cmd.Flags().GetString("serverlURL")
		Port, _ := cmd.Flags().GetString("port")

		runList(ServerlURL, Port)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().String("serverURL", "localhost", "Optional Server Address, defauts to localhost")
	listCmd.Flags().String("port", "50051", "port of server,  defaults to 50051")
}

func runList(ServerlURL string, port string) {
	fmt.Println("========  from client ---  doing ListCrawl Event")
	target := fmt.Sprintf("%s:%s", ServerlURL, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	log.Println("======== client trying to connect to :", target)

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := crawlerpb.NewCrawlerServiceClient(conn)

	// Prepare the request
	listReq := &crawlerpb.ListCrawlRequest{
		//RootUrl: RootUrl,
	}

	// Call the non-blocking StartCrawl method
	listRes, err := client.ListCrawl(context.Background(), listReq)

	if err != nil {
		log.Fatalf("Failed to list crawl: %v", err)
	}
	log.Printf("StartCrawl response: %s", listRes.UrlCrawled)

	for x, res := range listRes.UrlCrawled {
		log.Println("=== ", x, "===", res)
	}

	log.Println("===========   doing it the easy PrettyPrint way ===========")

	b, err := json.MarshalIndent(listRes, "", "  ")
	log.Println(string(b))

}
