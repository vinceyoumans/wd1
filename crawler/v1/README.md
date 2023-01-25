# setting up grpc

go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc


## protoc

protoc --go_out=plugins=grpc:. crawlerpb/v1/crawler.proto


## buf
I am using the buf cli to build out golang and rust libraries


1. Install the gRPC and protobuf packages for Go by running "go get -u google.golang.org/grpc" and "go get -u github.com/golang/protobuf/protoc-gen-go" in your terminal.
2. Generate the Go code for your proto file by running "protoc -I crawlerpb/ crawlerpb/v1/crawler.proto --go_out=plugins=grpc:crawlerpb" in the same directory as your proto file.
3. Implement the StartCrawl, CancelCrawl and MonitorCrawl handlers in your server. Each handler should take in the appropriate request and response message types from the proto file, and should use Go routines to perform the non-blocking actions.
4. In the StartCrawl handler, use the root_url and id fields from the request message to start a new go routine that performs the recursive link search. Use a map to store the state of each crawl, identified by the id field, so you can cancel or monitor it later.
5. In the CancelCrawl handler, use the id field from the request message to cancel the corresponding crawl, if it exists.
6. In the MonitorCrawl handler, use the id field from the request message to stream the progress of the corresponding crawl, if it exists. If the id field is not provided, stream the progress of all active crawls.
7. Register your server with a gRPC server instance, and start listening for incoming requests.

