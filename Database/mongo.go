package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"time"
)

type MongoDBStore struct {
	Db     *mongo.Database
	Client *mongo.Client
}



func NewDataStore() *MongoDBStore {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		ErrLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		ErrLog.Print(err)
	} else {
		InfoLog.Print("connect database success")
	}
	return &MongoDBStore{
		Db:     client.Database(DbName),
		Client: client,
	}
}