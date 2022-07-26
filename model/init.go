package model

import (
	"context"
	"fmt"
	"github.com/codecodify/chat/vars"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mongo *mongo.Database

const MongoUri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20"

func init() {
	fmt.Println("初始化mongodb")
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoUri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	Mongo = client.Database("chat")

	go func() {
		<-vars.Quit
		fmt.Println("关闭mongodb连接句柄")
		_ = client.Disconnect(context.TODO())
	}()
}
