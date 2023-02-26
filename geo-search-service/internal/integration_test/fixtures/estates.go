//go:build integration
// +build integration

package fixtures

import (
	"bytes"
	"context"
	"encoding/json"
	"geo-search/internal/connectors"
	"geo-search/internal/entities"
	"geo-search/internal/utils"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"go.mongodb.org/mongo-driver/mongo"
)

// Modify here for new fixtures
var estates []entities.Estate = []entities.Estate{
	{Title: "integration_test_estate_1", Location: entities.Location{Lat: 51.8939742956985, Lon: 4.504141893127702}},
	{Title: "integration_test_estate_2", Location: entities.Location{Lat: 51.89351154372879, Lon: 4.505364310052456}},
	{Title: "integration_test_estate_3", Location: entities.Location{Lat: 51.89214546336856, Lon: 4.505871963015301}},
	{Title: "integration_test_estate_4", Location: entities.Location{Lat: 51.89272613716663, Lon: 4.503918468613102}},
	{Title: "integration_test_estate_5", Location: entities.Location{Lat: 51.89282880902377, Lon: 4.506271395002921}},
}

// topLeft := []float64{51.89563455208273, 4.49980791636062}
// bottomRight := []float64{51.888317570989095, 4.511376402730734}

type EstateFixtures struct {
	mongo         *mongo.Collection
	elasticsearch *elasticsearch.Client
}

func WireEstateFixtures(
	mongo *connectors.MongoConnector,
	elasticsearch *connectors.ElasticSearchConnector,
) *EstateFixtures {
	return &EstateFixtures{
		mongo:         mongo.Database.Collection(entities.ESTATES_COLLECTION),
		elasticsearch: elasticsearch.Client,
	}
}

func (fixtures *EstateFixtures) InsertFixtures(ctx context.Context) error {
	return fixtures.createBulkEstate(ctx, &estates)
}

func (fixtures *EstateFixtures) createBulkEstate(ctx context.Context, estates *[]entities.Estate) error {
	utils.LogInfo("inserting estates fixtures...")
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:   entities.ESTATES_INDEX,
		Client:  fixtures.elasticsearch,
		Refresh: "true",
	})
	if err != nil {
		return err
	}

	mongoBulkEstates := make([]interface{}, 0, len(*estates))
	for _, estate := range *estates {
		mongoBulkEstates = append(mongoBulkEstates, estate)

		json, err := json.Marshal(estate)
		if err != nil {
			return err
		}

		err = indexer.Add(ctx, esutil.BulkIndexerItem{
			Action: "index",
			Index:  entities.ESTATES_INDEX,
			Body:   bytes.NewReader(json),
		})
		if err != nil {
			return err
		}
	}

	// MongoDB
	_, err = fixtures.mongo.InsertMany(ctx, mongoBulkEstates)
	if err != nil {
		return err
	}

	// Elasticsearch
	err = indexer.Close(ctx)
	if err != nil {
		return err
	}

	utils.LogInfo("Insert estates fixtures completed!")
	return nil
}
