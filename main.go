package main

import (
	"eztravel-wh/internals/db"
	"eztravel-wh/internals/server"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
)

func main() {
	apiServer := server.NewServer(mongoClient)
	apiServer.StartServer()
}

func init() {
	mongoClient = db.ConnectDB()
}
