package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func TestGetEntryByEntryIDwithMock(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer mt.Close()

// 	mt.Run("success", func(mt *mtest.T) {
// 		expectedEntry := classic_entry{
// 			ID:   1,
// 			Text: "this is the first text",
// 			Tag:  "basic",
// 		}

// 		mCollection := mtest.Collection{Name: "entry"}
// 		mt.CreateCollection(mCollection, true)
// 		mt.AddMockResponses(mtest.CreateCursorResponse(1, "main.readEntry", mtest.FirstBatch, bson.D{
// 			primitive.E{Key: "Text", Value: expectedEntry.Text},
// 			primitive.E{Key: "ID", Value: expectedEntry.ID},
// 			primitive.E{Key: "Tag", Value: expectedEntry.Tag},
// 		}))
// 		userResponse := getEntryByEntryID(mt.Client, 1)
// 		assert.Equal(t, expectedEntry, userResponse)
// 	})
// }

func TestGetEntryByEntryID(t *testing.T) {
	client := connect()
	entry := getEntryByEntryID(client, 1)
	expected_entry := classic_entry{
		ID:   1,
		Text: "this is the first text",
		Tag:  "basic",
	}
	assert.Equal(t, entry, expected_entry)
}

func TestGetRandomEntry(t *testing.T) {
	client := connect()
	entry := getRandomEntry(client)
	expected_entry1 := classic_entry{
		ID:   1,
		Text: "this is the first text",
		Tag:  "basic",
	}
	expected_entry2 := classic_entry{
		ID:   2,
		Text: "this is the second text",
		Tag:  "basic",
	}
	expected_entry3 := classic_entry{
		ID:   3,
		Text: "this is the third text",
		Tag:  "basic",
	}
	expected_entry4 := classic_entry{
		ID:   entry.ID,
		Text: "test string",
		Tag:  "basic",
	}
	expected_entry_array := []classic_entry{expected_entry1, expected_entry2, expected_entry3, expected_entry4}
	assert.Contains(t, expected_entry_array, entry)
}

func TestGetAllEntriesWithTag(t *testing.T) {
	client := connect()
	entry := getAllEntriesWithCategory(client, "test1")
	expected_entry1 := classic_entry{
		ID:   1,
		Text: "this is the first text",
		Tag:  "basic",
	}
	expected_entry2 := classic_entry{
		ID:   2,
		Text: "this is the second text",
		Tag:  "basic",
	}
	expected_entry3 := classic_entry{
		ID:   3,
		Text: "this is the third text",
		Tag:  "basic",
	}
	expected_entry_array := []classic_entry{expected_entry1, expected_entry2, expected_entry3}
	assert.Equal(t, expected_entry_array, entry)
}

func TestInsertIntoEntry(t *testing.T) {
	test := "test string"
	client := connect()
	result := insertEntry(client, test)
	assert.Equal(t, true, result)
	coll := client.Database(dbName).Collection(entryCollection)
	filter := primitive.E{Key: "text", Value: test}
	deleteResult, err := coll.DeleteMany(context.TODO(), bson.D{filter})
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, 0, int(deleteResult.DeletedCount))
}

func TestInsertIntoCategoryAndUpdateCategory(t *testing.T) {
	test := "test"
	client := connect()
	result := insertCategory(client, test, 1)
	assert.Equal(t, true, result)
	result = updateCategory(client, test, 2)
	assert.Equal(t, true, result)
	coll := client.Database(dbName).Collection(categoryCollection)
	filter := primitive.E{Key: "Category", Value: test}
	deleteResult, err := coll.DeleteMany(context.TODO(), bson.D{filter})
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, 0, int(deleteResult.DeletedCount))
}
