package initialize

import (
	"context"
	"gin-vue-admin/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Mongo() {
	mgoConfig := global.GVA_CONFIG.Mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mgoConfig.Username + ":" + mgoConfig.Password + "@" + mgoConfig.Host + ":" + mgoConfig.Port + "/" + mgoConfig.DbName))

	if err != nil {
		global.GVA_LOG.Fatal(err)
	}

	ctx, cancelConnect := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelConnect()
	err = client.Connect(ctx)
	if err != nil {
		global.GVA_LOG.Fatal(err)
	}
	ctx, cancelPing := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelPing()
	err = client.Ping(ctx, nil)
	if err != nil {
		global.GVA_LOG.Fatal(err)
	} else {
		global.GVA_LOG.Info("ping to mongo server OK")
		global.GVA_MONGO = client
	}

}
