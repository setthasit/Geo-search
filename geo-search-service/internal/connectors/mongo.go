package connectors

import (
	"context"
	"fmt"
	"geo-search/internal/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnector struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func WireMongoConnector(mongoConfig *config.MongoConfig) *MongoConnector {
	credential := options.Credential{
		Username: mongoConfig.Username,
		Password: mongoConfig.Password,
	}

	uri := fmt.Sprintf("mongodb://%s", mongoConfig.URI)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		logrus.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}

	return &MongoConnector{
		Client:   client,
		Database: client.Database(mongoConfig.Database),
	}
}
