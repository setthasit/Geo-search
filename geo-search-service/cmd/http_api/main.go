package main

import (
	"geo-search/cmd/http_api/di"
	"geo-search/internal/utils"
)

func main() {
	utils.BuildLogger("geo-search-http")
	server := di.InitializeAPI()
	server.Run()
}
