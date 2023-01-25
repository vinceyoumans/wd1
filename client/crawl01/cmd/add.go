/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	crawlerpb "wd1/gen/go/crawler/v1"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Will add RootURL to QUE but does not start fetching",
	Long: `Add a rootURL for Que for processing.   This does not start the crawl.
the concept is to create a long list of rootURL in Pending state.  Once the Crawl Start is run,
the RootURL's will be run as go routines'
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		urlRoot, _ := cmd.Flags().GetString("rootURL")
		ServerlURL, _ := cmd.Flags().GetString("serverlURL")
		Port, _ := cmd.Flags().GetString("port")

		runAdd(urlRoot, ServerlURL, Port)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("rootURL", "vyoumans.com", "The URL to be Crawled, defaults to vyoumans.com")
	addCmd.Flags().String("serverURL", "localhost", "Optional Server Address, defauts to localhost")
	addCmd.Flags().String("port", "50051", "port of server,  defaults to 50051")

}

func runAdd(rootURL string, urlServer string, port string) {

	target := fmt.Sprintf("%s:%s", urlServer, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	log.Println("======== client add  connect to :", target)

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := crawlerpb.NewCrawlerServiceClient(conn)

	addReq := &crawlerpb.AddCrawlRequest{ //  LongListCrawlRequest{
		RootUrl: rootURL,
	}

	AddRes, err := client.AddCrawl(context.Background(), addReq)

	log.Printf("added %s  - getting message %s  ", rootURL, AddRes.Message)
	if err != nil {
		log.Printf(" client add Error %d", err)
	}

}
