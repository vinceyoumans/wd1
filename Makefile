.PHONY: server serverTest client runServer rustgen rustServer rustClient

server:
	@echo "Linting Go files...   except my goLint is not working for some reason...  need to check that out"
#	@echo "Linting Go files..."
#	golint ./server/v01/*.go
	@echo "Building Go files..."
	go build -o ./build/go/crawlserver ./server/v01/*.go
	@echo "Build completed"

serverTest:
	@echo "Running Go tests..."
	go test ./server/v01/*.go
	@echo "Tests completed"

runServer:
	@echo "Running server..."
	./build/go/crawlserver
	@echo "Server stopped"

client:
	@echo "Linting Go Client files...   except my goLint is not working for some reason...  need to check that out"
#	@echo "Linting Go files..."
#	golint ./client/crawl01/*.go
	@echo "Building CLEINT Go files..."
	go build -o ./build/go/crawl ./client/crawl01/*.go
	@echo "client Build completed - look in build dir"


rustgen:
	@echo "RUST GEN gRPC/Proto library.   Not sure how to lint, but the goland linter should be good enough"
	@echo "Building CLEINT Go files..."
	#go build -o ./build/go/crawl ./client/crawl01/*.go
	protoc -I=. crawler.proto --rust_out=. --grpc_out=. --plugin=protoc-gen-grpc=path/to/grpc_rust_plugin

	@echo "client Build completed - look in build dir"
