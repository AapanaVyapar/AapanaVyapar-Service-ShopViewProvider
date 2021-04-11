package data_base

import (
	"aapanavyapar-service-shopviewprovider/configurations/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type MongoDataBase struct {
	mutex sync.RWMutex
	Data  *mongo.Client
}

func NewDataBase() *MongoDataBase {

	client := mongodb.InitMongo()

	return &MongoDataBase{
		Data: client,
	}
}
