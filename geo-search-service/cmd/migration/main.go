package main

import (
	"geo-search/cmd/migration/di"
	"geo-search/cmd/migration/scripts/estates"
	"geo-search/internal/utils"
	"os"
)

func main() {
	utils.BuildLogger("geo-search-migration")
	client := di.InitializeMigration()

	command := os.Args[1]
	switch command {
	case "migrate:elasticsearch":
		estates.MigrateElasticsearch(client.Elasticsearch.Client)
	case "seed":
		estates.SeedEstates(client.Mongo.Database, client.Elasticsearch.Client)
	default:
		utils.LogFatal("command not found")
	}
}
