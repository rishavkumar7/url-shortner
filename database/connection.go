package database

import (
	"context"
	"os"
	"time"

	"github.com/rishavkumar7/url-shortner/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetUrlFromCode(string, string) (types.UrlDb, error)
}

func ConnectDb() {
	uri := os.Getenv("DB_HOST")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	Mgr = &manager{
		connection: client,
		ctx:        ctx,
		cancel:     cancel,
	}
}
