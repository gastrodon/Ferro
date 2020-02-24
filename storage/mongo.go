package storage

import (
	"monke-cdn/log"
	"monke-cdn/util"

	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"path/filepath"
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
 * login    -> database login username
 * password    -> database login password
 * uri      -> database location
 * name     -> database to load from uri
 *
 * returns
 * err      -> error while connecting
 */
func ConnectTo(login, password, uri, name string) (err error) {
	log.Tracef("Connecting to %s as %s", uri, login)
	client, err = mongo.Connect(timeout_ctx(2), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:27017", login, password, uri)))

	if err == nil {
		log.Tracef("Probing mongod server at %s as %s", uri, login)
		err = client.Ping(timeout_ctx(20), readpref.Primary())
	}

	if err == nil {
		database = client.Database(name)
		media = database.Collection("media")
		return
	}

	log.Fatalf("Failed to connect to %s: %s", uri, err)

	return
}

func DropDB(name string) {
	log.Tracef("Dropping database %s", name)
	client.Database(name).Drop(context.TODO())
}

func NewUUID() (id string, err error) {
	log.Tracef("Generating a new UUIDv4")
	var exists bool

	for true {
		id = uuid.NewV4().String()
		_, exists, err = GetUnique(bson.D{{"id", id}})

		if err != nil || !exists {

			break
		}
	}

	log.Tracef("Generated %s", id)
	return
}

func GetUnique(filt bson.D) (result bson.M, exists bool, err error) {
	log.Traceln("Getting some unique object")
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

	log.Tracef("Unique was found: %t", exists)
	return
}

func DeleteUnique(filt bson.D) (deleted bool, err error) {
	log.Traceln("Deleting some unique object")
	var result *mongo.DeleteResult
	result, err = media.DeleteOne(timeout_ctx(5), filt)
	if err == nil {
		deleted = result.DeletedCount == 1
	}

	log.Tracef("Unique was deleted: %t", deleted)
	return
}

func CreateReference(id, mime string, md5 []byte) (conflicts bool, err error) {
	log.Tracef("Creating a reference of %s (%s) -> file", id, mime)
	_, conflicts, err = GetUnique(bson.D{{"id", id}})
	if conflicts || err != nil {
		return
	}

	var absolute string
	absolute, err = filepath.Abs(fmt.Sprintf("%s%s", file_root, id))
	if err != nil {
		return
	}

	var writable map[string]interface{} = map[string]interface{}{
		"path":    absolute,
		"id":      id,
		"mime":    mime,
		"md5":     md5,
		"created": time.Now().Unix(),
	}

	_, err = media.InsertOne(context.Background(), writable)
	return
}
