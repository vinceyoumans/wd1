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
	crawlerpb "wd1/gen/go/crawler/v1"

	"github.com/spf13/cobra"
)

// monitor01Cmd represents the monitor01 command
var monitor01Cmd = &cobra.Command{
	Use:   "monitor01",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor01 called")

		urlRoot, _ := cmd.Flags().GetString("rootlURL")
		ServerlURL, _ := cmd.Flags().GetString("serverlURL")
		Port, _ := cmd.Flags().GetString("port")

		runMonitor01(urlRoot, ServerlURL, Port)
	},
}

func init() {
	rootCmd.AddCommand(monitor01Cmd)

	monitor01Cmd.Flags().String("rootURL", "vyoumans.com", "The URL to be Crawled, defaults to vyoumans.com")
	monitor01Cmd.Flags().String("serverURL", "localhost", "Optional Server Address, defauts to localhost")
	monitor01Cmd.Flags().String("port", "50051", "port of server,  defaults to 50051")
}

func runMonitor01(rootURL string, urlServer string, port string) {

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

	log.Println("***  urlRoot :", rootURL)
	in := &crawlerpb.MonitorCrawl01Request{
		RootUrl: rootURL,
	}

	stream, err := client.MonitorCrawl01(context.Background(), in)

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
