package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAll(t *testing.T) {
	client := connect()
	entryCollection := "test_entry"
	categoryCollection := "test_category"
	categorylistCollection := "test_categorylist"
	expected_text1 := "this is the first text"
	expected_text2 := "this is the second text"
	expected_text3 := "this is the third text"

	//test create entry
	success, err := createEntry(client, expected_text1, "test1", entryCollection, categorylistCollection, categoryCollection)
	assert.Equal(t, true, success)
	assert.Equal(t, nil, err)
	success, err = createEntry(client, expected_text2, "test1", entryCollection, categorylistCollection, categoryCollection)
	assert.Equal(t, true, success)
	assert.Equal(t, nil, err)
	success, err = createEntry(client, expected_text3, "test2", entryCollection, categorylistCollection, categoryCollection)
	assert.Equal(t, true, success)
	assert.Equal(t, nil, err)

	//test get entries with category
	texts := getTextFromAllEntriesWithCategory(client, "test1", 0, 10, categoryCollection, entryCollection) //test a larger length
	expected_entry_array := []string{expected_text1, expected_text2}
	assert.Equal(t, expected_entry_array, texts)
	texts2 := getTextFromAllEntriesWithCategory(client, "test1", 10, 12, categoryCollection, entryCollection) //test a bad start point
	expected_entry_array2 := []string{}
	assert.Equal(t, expected_entry_array2, texts2)
	texts3 := getTextFromAllEntriesWithCategory(client, "test2", 0, 1, categoryCollection, entryCollection) //test a bad start point
	expected_entry_array3 := []string{expected_text3}
	assert.Equal(t, expected_entry_array3, texts3)

	//test random entry
	text := getTextFromRandomEntry(client, entryCollection)
	expected_entry_array = []string{expected_text1, expected_text2, expected_text3}
	assert.Contains(t, expected_entry_array, text)

	//test all categories
	categories := getAllCategories(client, categorylistCollection)
	assert.Equal(t, 2, len(categories))
	assert.Equal(t, "test1", categories[0])
	assert.Equal(t, "test2", categories[1])

	//reset all test collections
	coll := client.Database(dbName).Collection(categoryCollection)
	deleteResult, err := coll.DeleteMany(context.TODO(), bson.D{})
	assert.Equal(t, 2, int(deleteResult.DeletedCount))
	assert.Equal(t, nil, err)
	coll = client.Database(dbName).Collection(entryCollection)
	deleteResult, err = coll.DeleteMany(context.TODO(), bson.D{})
	assert.Equal(t, 3, int(deleteResult.DeletedCount))
	assert.Equal(t, nil, err)
	coll = client.Database(dbName).Collection(categorylistCollection)
	categoryListUpdate := classic_category_list{Categories: []string{}}
	result, err := coll.ReplaceOne(context.TODO(), bson.D{}, categoryListUpdate)
	assert.Equal(t, 1, int(result.ModifiedCount))
	assert.Equal(t, nil, err)
}
