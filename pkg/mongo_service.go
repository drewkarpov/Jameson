package pkg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoImageService struct {
	Collection *mongo.Collection
	Database   *mongo.Database
}

func (d MongoImageService) Init() MongoImageService {
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	var cred = options.Credential{Username: "mongoadmin", Password: "mongoadmin"}
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred))
	d.Database = client.Database("jameson")
	d.Collection = client.Database("jameson").Collection("projects")
	return MongoImageService{Collection: d.Collection, Database: d.Database}
}
