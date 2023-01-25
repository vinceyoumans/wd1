# logic to the server...


Although the server starts, the Crawling is not started until client turns on crawling.
This is so I can add multiple URL's to be crawled before the process starts.
The crawling can also be stopped with out effecting the job que.  The server just goes into sleep until a start command is sent from client.

There are two nonBlocking clients 
- list, which only returns the que and status.
- longList which returns above plus the ChildURL's

There are two client blocking streams
- Monitor, which streams every 4 seconds, a...
  - (default) a full display of all jobs with status and links.
  - (rootURL) displays a specific RootURL job
- Monitor2, which streams ( every 4 seconds), a... 
  - ( default )a full json to client .
  - (rootURL) a specific JSON to client



# logic to start fetching


main -> StartCrawl -> Fetch01(Fetch01.go) ->
func fetchupdateUrlsmap(url string, urls []string) { (Fetch01.go)




# unused files
scan.go
scanb.go
