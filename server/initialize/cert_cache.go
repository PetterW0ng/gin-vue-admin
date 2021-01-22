package initialize

import (
	"context"
	"fmt"
	"gin-vue-admin/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const dbName, collection, UserCertKey = "certificate", "certificate", "cert:user"

func InitUserCertCache() {
	mongoClient := global.GVA_MONGO
	collection := mongoClient.Database(dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	match := bson.D{{"$match", bson.D{{"issueTime", primitive.Regex{Pattern: "^2020"}}}}} //bson.D{{"$regex", "/^2020/"}}
	group := bson.D{{"$group", bson.D{{"_id", "$phone"}, {"count", bson.D{{"$sum", 1}}}}}}
	sort := bson.D{{"$sort", bson.D{{"count", -1}}}}
	opts := options.Aggregate().SetMaxTime(30 * time.Second)
	cur, err := collection.Aggregate(ctx, mongo.Pipeline{match, group, sort}, opts)
	if err != nil {
		panic("mongoInit failure")
	}
	defer cur.Close(context.Background())
	//var result []userRank
	var results []bson.M
	err = cur.All(context.Background(), &results)
	if err != nil {
		panic("mongoInit bound data failure")
	}
	size := len(results)
	rMap := make(map[string]interface{}, 0)
	for i, item := range results {
		rMap[item["_id"].(string)] = fmt.Sprintf("%.2f", float64((size-i)*100)/float64(size))
	}
	global.GVA_REDIS.HMSet(UserCertKey, rMap)
}

type userRank struct {
	_id   string
	count int
}
