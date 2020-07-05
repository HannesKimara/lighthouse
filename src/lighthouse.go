package lighthouse

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Application struct {
	Config Config
	Client *mongo.Client
	Db *mongo.Database
}

func (app *Application) InitMongoConnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI(app.Config.Mongo.URL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	app.Client = client
	app.Db = app.Client.Database(app.Config.Mongo.Database)
}

func CreateApp() Application {
	app := Application{}

	app.Config = Config{}
	app.Config.LoadConf()

	app.InitMongoConnection()

	return app
}
