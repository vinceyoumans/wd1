/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"

	"github.com/spf13/cobra"
	crawlerpb "wd1/gen/go/crawler/v1"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Will monitor a specific or all urlRoots",
	Long: `monitor stream a specific URL and its status or it will default to 
returning all of the URLs crawled and all of their responses

if you do not inlcude a -RootURL as a parameter, then all crawled URL's will be returned'

-rootURL is URL to crawl...  defaults to vyoumans.com 
-serverURL Optional Server Address, defauts to localhosts
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


`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor called")
		URLRootCmd, _ := cmd.Flags().GetString("rootURL")
		ServerURL, _ := cmd.Flags().GetString("serverURL")
		Port, _ := cmd.Flags().GetString("port")

		log.Println("====   calling startMonitor URLRootCMD :", URLRootCmd)
		startMonitor(URLRootCmd, ServerURL, Port)

	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	//monitorCmd.Flags().String("rootURL", "vyoumans.com", "The URL to be Crawled, defaults to vyoumans.com")
	monitorCmd.Flags().String("rootURL", "", "The URL to be Crawled, defaults to all jobs")
	monitorCmd.Flags().String("serverURL", "localhost", "Optional Server Address, defauts to localhost")
	monitorCmd.Flags().String("port", "50051", "port of server,  defaults to 50051")
}

func startMonitor(urlRoot string, ServerURL string, port string) {
	target := fmt.Sprintf("%s:%s", ServerURL, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	log.Println("======== client monitor trying to connect to :", target)

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	log.Println("======== client connected to :", target)

	// create stream
	client := crawlerpb.NewCrawlerServiceClient(conn)

	log.Println("***  urlRoot :", urlRoot)
	in := &crawlerpb.MonitorCrawlRequest{
		RootUrl: urlRoot,
		//RootUrl: "urlRoot",
	}

	stream, err := client.MonitorCrawl(context.Background(), in)

	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		log.Println("=======    in client monitor func")
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("===========   hit eof ===========")
				//done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Println("========    serverMonitor =================")
			log.Println("Resp rootURL - received: ", resp.RootUrl)
			log.Println("Resp rootURL - received: ", resp)

		}
	}()

	<-done //we will wait until all response is received
	log.Printf("xxxx  monitor finished")
}
