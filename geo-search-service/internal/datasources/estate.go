package datasources

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"geo-search/internal/common/internalerror"
	"geo-search/internal/connectors"
	"geo-search/internal/entities"
	"geo-search/internal/repositories"
	"geo-search/internal/utils"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EstateDatasource struct {
	elasticSearch *connectors.ElasticSearchConnector
	mongoClient   *mongo.Collection
}

var _ repositories.EstateRepository = &EstateDatasource{}

func WireEstateDatasource(elasticSearch *connectors.ElasticSearchConnector, mongoClient *connectors.MongoConnector) *EstateDatasource {
	ds := &EstateDatasource{
		elasticSearch: elasticSearch,
		mongoClient:   mongoClient.Database.Collection(entities.ESTATES_COLLECTION),
	}

	return ds
}

func (*EstateDatasource) GetAll(ctx context.Context) ([]entities.Estate, error) {
	return nil, nil
}

func (*EstateDatasource) GetByID(ctx context.Context, id uint) (*entities.Estate, error) {
	return nil, nil
}

func (ds *EstateDatasource) GetByBoundBox(ctx context.Context, topLeft *entities.Location, bottomRight *entities.Location) ([]entities.Estate, int, error) {
	query := connectors.Esq{
		"size": 1000,
		"query": connectors.Esq{
			"bool": connectors.Esq{
				"must": connectors.Esq{
					"match_all": connectors.Esq{},
				},
				"filter": connectors.Esq{
					"geo_bounding_box": connectors.Esq{
						"location": connectors.Esq{
							"top_left": connectors.Esq{ // 51.601034, -0.08782414703561299
								"lat": topLeft.Lat,
								"lon": topLeft.Lon,
							},
							"bottom_right": connectors.Esq{ // 51.5991415507311, -0.08261261508703863
								"lat": bottomRight.Lat,
								"lon": bottomRight.Lon,
							},
						},
					},
				},
			},
		},
	}
	queryJson, err := json.Marshal(query)
	if err != nil {
		utils.LogError(err)
		return nil, 0, internalerror.EstateRepositoryParsingRequest(err)
	}

	res, err := ds.elasticSearch.Client.Search(
		ds.elasticSearch.Client.Search.WithContext(ctx),
		ds.elasticSearch.Client.Search.WithIndex(entities.ESTATES_INDEX),
		ds.elasticSearch.Client.Search.WithBody(bytes.NewBuffer(queryJson)),
		ds.elasticSearch.Client.Search.WithPretty(),
	)
	if err != nil {
		utils.LogError(err)
		return nil, 0, internalerror.EstateRepositoryRequest(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		esErr := &connectors.EsSearchError{}
		err = json.NewDecoder(res.Body).Decode(esErr)
		if err != nil {
			utils.LogError(err)
			return nil, 0, internalerror.EstateRepositoryRequest(err)
		}
		utils.LogErrorf(esErr.Error.Reason)
		return nil, 0, internalerror.EstateRepositoryRequest(errors.New(esErr.Error.Reason))
	}

	result := &connectors.EsResponse[entities.Estate]{}
	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		utils.LogError(err)
		return nil, 0, internalerror.EstateRepositoryParsingResponse(err)
	}

	estates := make([]entities.Estate, len(result.Hits.Hits))
	for idx, estate := range result.Hits.Hits {
		estates[idx] = estate.Source
	}

	return estates, result.Hits.Total.Value, nil
}

func (ds *EstateDatasource) Create(ctx context.Context, estate *entities.Estate) error {
	now := utils.NowUTC()
	estate.CreatedAt = now
	estate.UpdatedAt = now

	dbResult, err := ds.mongoClient.InsertOne(ctx, estate)
	if err != nil {
		utils.LogError(err)
		return internalerror.EstateRepositoryCannotSave(err)
	}

	data, err := json.Marshal(estate)
	if err != nil {
		utils.LogError(err)
		return internalerror.EstateRepositoryParsingRequest(err)
	}

	esIndexReq := esapi.IndexRequest{
		Index:      entities.ESTATES_INDEX,
		DocumentID: dbResult.InsertedID.(primitive.ObjectID).Hex(),
		Body:       bytes.NewBuffer(data),
		Refresh:    "true",
	}

	reqCtx := context.Background()
	esIndexRes, err := esIndexReq.Do(reqCtx, ds.elasticSearch.Client)
	if err != nil {
		utils.LogError(err)
		return internalerror.EstateRepositoryRequest(err)
	}

	defer esIndexRes.Body.Close()
	if esIndexRes.IsError() {
		esErr := &connectors.EsSearchError{}
		err = json.NewDecoder(esIndexRes.Body).Decode(esErr)
		if err != nil {
			utils.LogError(err)
			return internalerror.EstateRepositoryIndexing(err)
		}
		utils.LogErrorf(esErr.Error.Reason)
		return internalerror.EstateRepositoryIndexing(errors.New(esErr.Error.Reason))
	}

	return nil
}
