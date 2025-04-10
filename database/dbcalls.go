package database

import (
	"context"

	"github.com/rishavkumar7/url-shortner/constants"
	"github.com/rishavkumar7/url-shortner/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	inst := mgr.connection.Database(constants.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

func (mgr *manager) GetUrlFromCode(code string, collectionName string) (resp types.UrlDb, err error) {
	inst := mgr.connection.Database(constants.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	return resp, err
}
