package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	expected_entry_array := []classic_entry{expected_entry1, expected_entry2, expected_entry3}
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
