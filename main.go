package main

import (
	"github.com/hasssanezzz/rest-workers/api"
)

func main() {
	listenAddr := ":3000"
	server := api.NewServer(listenAddr, 5)
	server.Start()
}
