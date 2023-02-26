package main

import (
	"geo-search/cmd/grpc_api/di"
	"geo-search/internal/utils"
)

func main() {
	utils.BuildLogger("geo-search-grpc")
	server := di.InitializeGrpc()
	server.Run()
}
