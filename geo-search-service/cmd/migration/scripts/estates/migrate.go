package estates

import (
	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

func MigrateElasticsearch(client *elasticsearch.Client) {
	createEstateIndex(client)
}

func SeedEstates(mongoClient *mongo.Database, esClient *elasticsearch.Client) {
	seedRandomEstates(mongoClient, esClient)
}
