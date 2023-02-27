package estates

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"geo-search/internal/connectors"
	"geo-search/internal/entities"
	"geo-search/internal/utils"
	"math/rand"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"go.mongodb.org/mongo-driver/mongo"
)

func createEstateIndex(client *elasticsearch.Client) {
	utils.LogInfo("Start migrating estates index...")
	ctx := context.Background()

	existsRequest := esapi.IndicesExistsRequest{
		Index: []string{entities.ESTATES_INDEX},
	}
	existsRes, err := existsRequest.Do(ctx, client)
	if err != nil {
		utils.LogFatal(err)
	}

	if !existsRes.IsError() {
		utils.LogInfof("%s index already existed", entities.ESTATES_INDEX)
		return
	}

	mapping := connectors.Esq{
		"mappings": connectors.Esq{
			"properties": connectors.Esq{
				"location": connectors.Esq{
					"type": "geo_point",
				},
			},
		},
	}

	body, err := json.Marshal(mapping)
	if err != nil {
		utils.LogFatal(err)
	}

	request := esapi.IndicesCreateRequest{
		Index: entities.ESTATES_INDEX,
		Body:  bytes.NewBuffer(body),
	}
	res, err := request.Do(ctx, client)
	if err != nil {
		utils.LogFatal(err)
	}
	defer res.Body.Close()
	if res.IsError() {
		utils.LogFatal(res)
	}

	utils.LogInfo("Done migrating estates index!")
}

func seedRandomEstates(mongoClient *mongo.Database, esClient *elasticsearch.Client) {
	utils.LogInfo("Start seeding estates data...")
	size := 10_000
	limit := 10_000
	collection := mongoClient.Collection(entities.ESTATES_COLLECTION)

	TopLeftLat := 66.94107
	TopLeftLon := 57.34416

	BottomRightLat := 46.81660
	BottomRightLon := 130.20578

	count := 1
	estates := make([]entities.Estate, 0, limit)
	for i := 0; i < size; i++ {
		now := utils.NowUTC()
		lat := randFloat(BottomRightLat, TopLeftLat)
		lon := randFloat(BottomRightLon, TopLeftLon)
		estates = append(estates, entities.Estate{
			Title: fmt.Sprintf("seed estate %d", i+1),
			Location: entities.Location{
				Lat: lat,
				Lon: lon,
			},
			CreatedAt: now,
			UpdatedAt: now,
		})

		if count == limit {
			createBulkEstate(&estates, collection, esClient)
			estates = estates[:0]
			count = 1
			continue
		}

		count++
	}

	if len(estates) > 0 {
		createBulkEstate(&estates, collection, esClient)
	}

	utils.LogInfo("Done seeding estates data!")
}

func createBulkEstate(estates *[]entities.Estate, mongoCollection *mongo.Collection, esClient *elasticsearch.Client) {
	utils.LogInfo("Preparing estates data...")
	ctx := context.Background()
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:   entities.ESTATES_INDEX,
		Client:  esClient,
		Refresh: "true",
	})
	if err != nil {
		utils.LogFatal(err)
	}

	mongoBulkEstates := make([]interface{}, 0, len(*estates))
	for _, estate := range *estates {
		mongoBulkEstates = append(mongoBulkEstates, estate)

		json, err := json.Marshal(estate)
		if err != nil {
			utils.LogFatal(err)
		}

		err = indexer.Add(ctx, esutil.BulkIndexerItem{
			Action: "index",
			Index:  entities.ESTATES_INDEX,
			Body:   bytes.NewReader(json),
		})
		if err != nil {
			utils.LogFatal(err)
		}
	}
	utils.LogInfo("Done Preparing estates data!")

	// MongoDB
	utils.LogInfo("Inserting estates data to mongo...")
	_, err = mongoCollection.InsertMany(ctx, mongoBulkEstates)
	if err != nil {
		utils.LogFatal(err)
	}
	utils.LogInfo("Done Inserting estates data to mongo!")

	// Elasticsearch
	utils.LogInfo("Inserting estates data to elasticsearch...")
	err = indexer.Close(ctx)
	if err != nil {
		utils.LogFatal(err)
	}
	utils.LogInfo("Done Inserting estates data to elasticsearch!")
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
