package main

import (
	"flag"

	"github.com/hasssanezzz/rest-workers/api"
)

func main() {
	listenAddr := flag.String("a", ":3030", "the listen address in which the server will listen to")
	workerCount := flag.Int("w", 5, "number of workers")
	flag.Parse()

	server := api.NewServer(*listenAddr, *workerCount)
	server.Start()
}
