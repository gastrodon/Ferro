package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

var id string

func TestMain(main *testing.M) {
	ConnectTo("foobar:foobar2000", "localhost", "test_storage")
	DropDB("test_storage")

	id, _ = NewUUID()
	CreateReference(id, "image/jpg", "f")

	os.Exit(main.Run())
	DropDB("test_storage")
}

func Test_GetUnique(test *testing.T) {
	var result bson.M
	var exists bool
	var err error
	result, exists, err = GetUnique(bson.D{{"id", id}})

	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("Test_GetUnique: UUIDv4 %s does not exist", id)
		test.FailNow()
	}

	var retrieved string = result["id"].(string)
	if retrieved != id {
		test.Errorf("Test_GetUnique: UUIDv4 mismatch: have %s, want %s", retrieved, id)
	}
}

func Test_GetUnique_dupe(test *testing.T) {
	var dupe_id string = "foobar"
	var writable map[string]interface{} = map[string]interface{}{
		"id": dupe_id,
	}

	media.InsertOne(context.Background(), writable)
	media.InsertOne(context.Background(), writable)

	var exists bool
	var err error
	_, exists, err = GetUnique(bson.D{{"id", dupe_id}})

	if !exists {
		test.Errorf("Test_GetUnique_dupe: dupe key %s does not exist", dupe_id)
	}

	if err == nil {
		test.Error("Test_GetUnique_dupe: no err returned")
	}

}

func Test_NewUUID(test *testing.T) {
	var new_id string
	var err error
	var exists bool
	var count int = 1000

	for count != 0 {
		count -= 1
		new_id, err = NewUUID()

		_, exists, err = GetUnique(bson.D{{"id", new_id}})

		if err != nil {
			test.Fatal(err)
		}

		if exists {
			test.Errorf("Test_NewUUID: UUID %s already exists", new_id)
		}

		media.InsertOne(context.Background(), bson.D{{"id", new_id}})
	}
}
