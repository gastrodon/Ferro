package storage

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"monke-cdn/util"
	"time"
)

var client *mongo.Client
var database *mongo.Database
var media *mongo.Collection

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
		media = database.Collection("media")
	}

	return
}

func DropDB(name string) {
	client.Database(name).Drop(context.TODO())
}

func NewUUID() (id string, err error) {
	var exists bool

	for true {
		id = uuid.NewV4().String()
		_, exists, err = GetUnique(bson.D{{"id", id}})

		if err != nil || !exists {
			break
		}
	}

	return
}

func GetUnique(filt bson.D) (result bson.M, exists bool, err error) {
	exists = false

	var cursor *mongo.Cursor
	cursor, err = media.Find(timeout_ctx(5), filt)

	if err != nil {
		return
	}

	var results []bson.M
	err = cursor.All(timeout_ctx(5), &results)

	if len(results) > 1 {
		err = util.TooManyResults(len(results))
	}

	if len(results) != 0 {
		exists = true
		result = results[0]
	}

	return
}

func CreateReference(id, mime string, md5 []byte) {
	var writable map[string]interface{} = map[string]interface{}{
		"id":      id,
		"mime":    mime,
		"md5":     md5,
		"created": time.Now().Unix(),
	}

	media.InsertOne(context.Background(), writable)
}
