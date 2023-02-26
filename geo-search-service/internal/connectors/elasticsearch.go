package connectors

import (
	"geo-search/internal/config"
	"geo-search/internal/utils"
	"io/ioutil"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
)

type Esq map[string]interface{}

type EsResponse[T any] struct {
	Hits struct {
		Hits  []EsResponseHit[T] `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}

type EsResponseHit[T any] struct {
	Source T `json:"_source"`
}

// Error type
type EsSearchError struct {
	Error  EsSearchErrorError `json:"error"`
	Status int                `json:"status"`
}

type EsSearchErrorError struct {
	RootCause []EsSearchErrorCause `json:"root_cause"`
	EsSearchErrorCause
}

type EsSearchErrorCause struct {
	Type         string `json:"type"`
	Reason       string `json:"reason"`
	ResourceType string `json:"resource.type"`
	ResourceId   string `json:"resource.id"`
	IndexUUID    string `json:"index_uuid"`
	Index        string `json:"index"`
}

// Connector type
type ElasticSearchConnector struct {
	Client *elasticsearch.Client
}

func WireElasticsearchConnector(elasticConfig *config.ElasticsearchConfig) *ElasticSearchConnector {
	cert, err := ioutil.ReadFile(elasticConfig.CertPath)
	if err != nil {
		logrus.Fatal(err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{elasticConfig.URI},
		Username:  elasticConfig.Username,
		Password:  elasticConfig.Password,
		CACert:    cert,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		utils.LogFatal(err)
	}

	res, err := client.Info()
	if err != nil {
		utils.LogFatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		utils.LogFatalf("Error: %s", res.String())
	}

	return &ElasticSearchConnector{
		Client: client,
	}
}
