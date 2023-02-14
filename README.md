# Wd1 -

## Assignment

The following test is to be implemented in Go and while you can take as much time as you need,
it's not expected that you spend more than 2 or 3 hours on it.

The test consists of implementing a "Web Crawler as a gRPC service".
The application consists of a command line client and a local service which runs the actual web crawling.
The communication between client and server should be defined as a gRPC service (*).
For each URL, the Web Crawler, creates a "site tree",
which is a tree of links with the root of the tree being the root URL.
The crawler should only follow links on the domain of the provided URL and not follow external links.
Bonus points for making it as fast as possible.



The command line client should provide the following operations:



$ crawl -start www.example.com  # signals the service to start crawling www.example.com

$ crawl -stop www.example.com  # signals the service to stop crawling www.example.com

$ crawl -list                                        # shows the current "site tree" for all crawled URLs.



## VY Optional

I am adding my own optional items.

1. Using BUF CLI to build gRPC library
2. might write the actual scan code as a RUST service
3. Add github action for cicd


# Problems
I just had 4 hours to do this, and I had a lot of ideas in my head on how.  Plus, I was working on some other projects that
I wanted to experiment with.

## V1
I had this idea that I could have a background process that would continually scan a Map of Pending jobs and run the jobs.
I begin the scan with a ```go StartScan()``` in server main, just before calling serve.   I use a go process because both the
serve() and StartScan() are both blocking...  
But in doing this, I think I am messing up  the WorkGroups, which the Scan function is using.


## v2
from Add,
I will add the rootURL to JOBS map then run StartScan.

I think this is a better Idea anyway... as there is no reason to scan if no jobs were added.

## ScanURL
As I was developing the ScanFunction, I tended to blacklist my IP address, including my own website vyoumans.com
It took me some time to realise this....   which consumed some time.







# notes

go get -u golang.org/x/lint/golint


# MAKE

```make server```
will build to ./build/go/server.exe



```make serverTest```
TODO: make server Testable
TODO:  add cicd github actions to for integration testing.
TODO: GHA deploy server to some cloud 
TODO: create a DART gRPC client
TODO:  do a RUST server

TODO: do video

TODO: implement fetch test that will only return childURL of RootURL.  Just forgot to do it.
TODO: Figure out a cancel strategy because by the time the app is added, it's to late to cancel as the job is already finished.
- perhaps some sort of interrupt

TODO: video of how things work

TODO: clean up the code a bit.


# NOTE






