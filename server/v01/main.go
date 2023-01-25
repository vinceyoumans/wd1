package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
	crawlerpb "wd1/gen/go/crawler/v1"
)

//var crawls = make(map[string][]string) // map[id] URL

// used by Monitors
var LUrlRoot string
var LUrlStatus crawlerpb.URLCurrentCrawls_URLCurrentStatus
var LUrlCrawled []string

type server struct {
	crawlerpb.UnimplementedCrawlerServiceServer
}

// URLsMap - slice of jobs to be crawled
var URLsMap = make(map[string]URLCrawl)
var RunInStartState bool // switches server fetch on and off

func init() {
	//urlCrawls = append(urlCrawls, &crawlerpb.URLCurrentCrawls{
	//	UrlRoot:   "https://www.example0.com",
	//	UrlStatus: crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_IN_PROGRESS,
	//})
	//
	//urlCrawls = append(urlCrawls, &crawlerpb.URLCurrentCrawls{
	//	UrlRoot:   "https://www.example1.com",
	//	UrlStatus: crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_UNSPECIFIED,
	//})
	//
	//urlCrawls = append(urlCrawls, &crawlerpb.URLCurrentCrawls{
	//	UrlRoot:   "https://www.example2.com",
	//	UrlStatus: crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_START,
	//})
}

func main() {
	host := flag.String("host", "localhost", "server host")
	port := flag.String("port", "50051", "server port")
	flag.Parse()
	target := fmt.Sprintf("%s:%s", *host, *port)

	lis, err := net.Listen("tcp", target)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	RunInStartState = false

	log.Println("**********   server01 listening on ", target)

	// todo add secure connection
	var opts []grpc.ServerOption
	//if *tls {
	//	if *certFile == "" {
	//		*certFile = data.Path("x509/server_cert.pem")
	//	}
	//	if *keyFile == "" {
	//		*keyFile = data.Path("x509/server_key.pem")
	//	}
	//	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	//	if err != nil {
	//		log.Fatalf("Failed to generate credentials %v", err)
	//	}
	//	opts = []grpc.ServerOption{grpc.Creds(creds)}
	//}

	grpcServer := grpc.NewServer(opts...)
	background := &backgroundService{}
	go background.Run(context.Background())

	crawlerpb.RegisterCrawlerServiceServer(grpcServer, &server{})

	log.Println("-------- AAA Starting Serve")

	grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("[x] serve: %v", err)
	}

}

//=====================================================
//type CrawlerServiceServer interface {
//	AddCrawl(ctx context.Context, in *AddCrawlRequest, opts ...grpc.CallOption) (*AddCrawlResponse, error)
//	ListCrawl(ctx context.Context, in *ListCrawlRequest, opts ...grpc.CallOption) (*ListCrawlResponse, error)
//	LongListCrawl(ctx context.Context, in *LongListCrawlRequest, opts ...grpc.CallOption) (*LongListCrawlResponse, error)
//	MonitorCrawl(ctx context.Context, in *MonitorCrawlRequest, opts ...grpc.CallOption) (CrawlerService_MonitorCrawlClient, error)
//	MonitorCrawl01(ctx context.Context, in *MonitorCrawl01Request, opts ...grpc.CallOption) (CrawlerService_MonitorCrawl01Client, error)
//}

func (s *server) AddCrawl(ctx context.Context, in *crawlerpb.AddCrawlRequest) (*crawlerpb.AddCrawlResponse, error) {
	// will add a URL to jobs
	F10AddPendingUrlToCrawlList(in.RootUrl)
	out := crawlerpb.AddCrawlResponse{
		Message: "added URL to job" + in.RootUrl,
	}
	return &out, nil
}

//
//func (s *server) StartCrawl(ctx context.Context, in *crawlerpb.StartCrawlRequest) (*crawlerpb.StartCrawlResponse, error) {
//	//  will start running jobs on the server, but is non blocking to client
//	RunInStartState = true
//
//	for {
//		for key, val := range URLsMap {
//			if val.URLStatus == crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_PENDING {
//				log.Printf("============   starting Fetch for %s", key)
//				F11ModifyPendingToStarting(key)
//				Fetch01(key) // blocking until all url children are done.  use go routine to unblock
//			} else {
//				log.Printf("=====   skipping URL %s", key)
//			}
//			time.Sleep(4 * time.Second)
//		}
//
//		RunInStartState = false
//
//		if !RunInStartState {
//			break
//		}
//	}
//	return nil, nil
//}

//// StopCrawl - the server is still running but the crawlFetch loop is stopped
//func (s *server) StopCrawl(ctx context.Context, in *crawlerpb.StopCrawlRequest) (*crawlerpb.StopCrawlResponse, error) {
//	RunInStartState = false
//	out := crawlerpb.StopCrawlResponse{
//		Message: "crawl fetch is stopped",
//	}
//	return &out, nil
//}

//// CancelCrawl - just cancels a RootURL JOB so that it is not run as a job
//func (s *server) CancelCrawl(ctx context.Context, in *crawlerpb.CancelCrawlRequest) (*crawlerpb.CancelCrawlResponse, error) {
//	log.Println("----   In server - calling ", in.RootUrl)
//	F20_cancel_URL_in_Crawl_List(in.RootUrl)
//	out := crawlerpb.CancelCrawlResponse{
//		Message: fmt.Sprintf("RootURL job Canceled : %s", in.RootUrl),
//	}
//	return &out, nil
//}

// ListCrawl - Short List of all url jobs.  Defaults to all jobs
func (s *server) ListCrawl(context.Context, *crawlerpb.ListCrawlRequest) (*crawlerpb.ListCrawlResponse, error) {

	// Implement your logic here
	var urlCrawls []*crawlerpb.URLCurrentCrawls

	//urlCrawls = append(urlCrawls, &crawlerpb.URLCurrentCrawls{
	//	UrlRoot:   "https://www.example0.com",
	//	UrlStatus: crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_IN_PROGRESS,
	//})
	//
	//urlCrawls = append(urlCrawls, &crawlerpb.URLCurrentCrawls{
	//	UrlRoot:   "https://www.example1.com",
	//	UrlStatus: crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_UNSPECIFIED,
	//})
	//
	//urlCrawls = append(urlCrawls, &crawlerpb.URLCurrentCrawls{
	//	UrlRoot:   "https://www.example2.com",
	//	UrlStatus: crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_START,
	//})

	for _, r := range URLsMap {
		urlCrawls = append(urlCrawls, &crawlerpb.URLCurrentCrawls{
			UrlRoot:   r.URLRoot,
			UrlStatus: r.URLStatus, //  crawlerpb.URLCurrentCrawls_URL_CURRENT_STATUS_START,
		})
	}

	out := &crawlerpb.ListCrawlResponse{
		UrlCrawled: urlCrawls,
	}
	return out, nil
}

//var (
//	URLMonitorCrawls_URLCurrentStatus_name = map[int32]string{
//		0: "URL_CURRENT_STATUS_UNSPECIFIED",
//		1: "URL_CURRENT_STATUS_NO_RECORD",
//		2: "URL_CURRENT_STATUS_PENDING",
//		3: "URL_CURRENT_STATUS_START",
//		4: "URL_CURRENT_STATUS_IN_PROGRESS",
//		5: "URL_CURRENT_STATUS_FINISHED",
//		6: "URL_CURRENT_STATUS_CANCELED",
//	}
//	URLMonitorCrawls_URLCurrentStatus_value = map[string]int32{
//		"URL_CURRENT_STATUS_UNSPECIFIED": 0,
//		"URL_CURRENT_STATUS_NO_RECORD":   1,
//		"URL_CURRENT_STATUS_PENDING":     2,
//		"URL_CURRENT_STATUS_START":       3,
//		"URL_CURRENT_STATUS_IN_PROGRESS": 4,
//		"URL_CURRENT_STATUS_FINISHED":    5,
//		"URL_CURRENT_STATUS_CANCELED":    6,
//	}
//)

// LongListCrawl - Returns LongList of JOBS, which includes the child URLS
func (s *server) LongListCrawl(context.Context, *crawlerpb.LongListCrawlRequest) (*crawlerpb.LongListCrawlResponse, error) {
	//return &crawlerpb.LongListCrawlResponse{UrlCrawled: []string{"Blah Blah", "bingo"}}, nil

	var urlCrawls []*crawlerpb.URLLongListCurrentCrawls

	//urlCrawls = append(urlCrawls, &crawlerpb.URLLongListCurrentCrawls{
	//	UrlRoot:    "https://www.example0.com",
	//	UrlStatus:  crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_FINISHED,
	//	UrlCrawled: []string{"https://www.example0.com/page1", "https://www.example0.com/page2"},
	//})
	//
	//urlCrawls = append(urlCrawls, &crawlerpb.URLLongListCurrentCrawls{
	//	UrlRoot:    "https://www.example1.com",
	//	UrlStatus:  crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_CANCELED,
	//	UrlCrawled: []string{"https://www.example1.com/page1", "https://www.example1.com/page2"},
	//})
	//
	//urlCrawls = append(urlCrawls, &crawlerpb.URLLongListCurrentCrawls{
	//	UrlRoot:    "https://www.example2.com",
	//	UrlStatus:  crawlerpb.URLLongListCurrentCrawls_URL_CURRENT_STATUS_START,
	//	UrlCrawled: []string{"https://www.example2.com/page1", "https://www.example2.com/page2"},
	//})

	log.Printf("=============   calling form LongListCrawl")
	log.Println(URLsMap)

	for _, r := range URLsMap {
		urlCrawls = append(urlCrawls, &crawlerpb.URLLongListCurrentCrawls{
			UrlRoot:    r.URLRoot,
			UrlStatus:  crawlerpb.URLLongListCurrentCrawls_URLCurrentStatus(r.URLStatus),
			UrlCrawled: r.URLCrawled,
		},
		)
	}
	out := &crawlerpb.LongListCrawlResponse{
		UrlCrawled: urlCrawls,
	}
	return out, nil
}

var ReturnThisJob bool = false

// MonitorCrawl - Streams jobs every 4 seconds to client
func (s *server) MonitorCrawl(req *crawlerpb.MonitorCrawlRequest, stream crawlerpb.CrawlerService_MonitorCrawlServer) error {
	rootToMonitor := req.RootUrl
	for {
		time.Sleep(4 * time.Second)
		log.Println("=========   in server Monitor url ", rootToMonitor)

		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "The stream has ended")
		default:
			LUrlRoot = ""
			LUrlStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(URLCurrentCrawls_URL_CURRENT_STATUS_UNSPECIFIED)
			LUrlCrawled = []string{}

			for _, val := range URLsMap {
				ReturnThisJob = false
				if req.RootUrl == "" {
					log.Println("return all jobs")
					ReturnThisJob = true
				} else if req.RootUrl == val.URLRoot {
					log.Println("return just this job")
					ReturnThisJob = true
				} else if req.RootUrl != val.URLRoot {
					ReturnThisJob = false
				}

				if ReturnThisJob {
					LUrlRoot = val.URLRoot
					LUrlStatus = val.URLStatus
					LUrlCrawled = val.URLCrawled

					err := stream.SendMsg(
						&crawlerpb.MonitorCrawlResponse{
							RootUrl: &crawlerpb.URLMonitorCrawls{
								UrlRoot:    val.URLRoot,
								UrlStatus:  crawlerpb.URLMonitorCrawls_URLCurrentStatus(val.URLStatus),
								UrlCrawled: val.URLCrawled,
							},
						},
					)
					if err != nil {
						return status.Error(codes.Canceled, "stream not being good for some reason")
					}
				}
			}
		}
	}
	return nil
}

// //MonitorCrawl01(*MonitorCrawl01Request, CrawlerService_MonitorCrawl01Server) error
func (s *server) MonitorCrawl01(req *crawlerpb.MonitorCrawl01Request, stream crawlerpb.CrawlerService_MonitorCrawl01Server) error {
	log.Println("----   in server... calling Monitor01")

	//var URLCrawlS []URLCrawl

	rootToMonitor := req.RootUrl
	for {
		time.Sleep(4 * time.Second)
		log.Println("=========   in server Monitor url ", rootToMonitor)

		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "The stream has ended")
		default:
			LUrlRoot = ""
			LUrlStatus = crawlerpb.URLCurrentCrawls_URLCurrentStatus(URLCurrentCrawls_URL_CURRENT_STATUS_UNSPECIFIED)
			LUrlCrawled = []string{}

			for _, val := range URLsMap {
				ReturnThisJob = false
				if req.RootUrl == "" {
					log.Println("return all jobs")
					ReturnThisJob = true
				} else if req.RootUrl == val.URLRoot {
					log.Println("return just this job")
					ReturnThisJob = true
				} else if req.RootUrl != val.URLRoot {
					ReturnThisJob = false
				}

				if ReturnThisJob {
					LUrlRoot = val.URLRoot
					LUrlStatus = val.URLStatus
					LUrlCrawled = val.URLCrawled

					err := stream.SendMsg(
						&crawlerpb.MonitorCrawlResponse{
							RootUrl: &crawlerpb.URLMonitorCrawls{
								UrlRoot:    val.URLRoot,
								UrlStatus:  crawlerpb.URLMonitorCrawls_URLCurrentStatus(val.URLStatus),
								UrlCrawled: val.URLCrawled,
							},
						},
					)
					if err != nil {
						return status.Error(codes.Canceled, "stream not being good for some reason")
					}
				}
			}
		}
	}
	return nil
}
