package storage

import (
	"context"
	"fmt"
	// "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client *mongo.Client
var database *mongo.Database

/*
 * seconds  -> context time
 * Get a context for a background task of n seconds
 *
 * returns
 * ctx      -> background context
 */
func timeout_ctx(seconds time.Duration) (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), seconds*time.Second)
	return
}

/*
 * Connect to some database. This should be called before database operations are done
 * login    -> database login info
 * uri      -> database location
 * name     -> database to load from uri
 *
 * returns
 * err      -> error while connecting
 */
func ConnectTo(login, uri, name string) (err error) {
	var ctx context.Context
	ctx = timeout_ctx(10)

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s@%s:27017", login, uri)))

	if err == nil {
		database = client.Database(name)
	}

	return
}
