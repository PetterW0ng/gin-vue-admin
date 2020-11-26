package service

import (
	"context"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const dbName, collection = "certificate", "certificate"

func GetCertByPhone(phone string) (error error, list []model.Holder) {
	mongoClient := global.GVA_MONGO
	collection := mongoClient.Database(dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	filter := bson.M{"phone": phone}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return err, list
	}
	defer cur.Close(context.Background())
	error = cur.All(context.Background(), &list)
	return
}

func GetCertHolderList(pgQuery request.PaginatedQuery) (error error, list []model.Holder, total int) {
	limit := pgQuery.PageSize
	offset := pgQuery.PageSize * (pgQuery.Page - 1)
	mongoClient := global.GVA_MONGO
	collection := mongoClient.Database(dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	filter := bson.M{}
	if len(pgQuery.Query) > 0 {
		filter = bson.M{"userName": pgQuery.Query}
	}
	var findOptions *options.FindOptions
	if limit > 0 {
		findOptions = &options.FindOptions{}
		findOptions.SetLimit(int64(limit))
		findOptions.SetSkip(int64(offset))
	}
	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return err, list, 0
	}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err, list, 0
	}
	defer cur.Close(context.Background())
	error = cur.All(context.Background(), &list)
	if error != nil {
		return error, list, 0
	}
	return nil, list, int(count)
}
