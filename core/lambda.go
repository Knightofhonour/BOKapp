package core

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

func inStringArray(arr []string, element string) bool {
	for _, v := range arr {
		if v == element {
			return true
		}
	}
	return false
}

func CreateEntry(client *mongo.Client, text string, category string, entryCollection string, categoryListCollection string, categoryCollection string) (bool, error) {
	success, new_id := insertEntry(client, text, entryCollection)
	if !success {
		return false, errors.New("failed to create entry")
	}
	categories := getAllCategories(client, categoryListCollection)
	categoryAlreadyPresent := inStringArray(categories, category)
	if categoryAlreadyPresent {
		success = updateCategory(client, category, new_id, categoryCollection)
		if !success {
			return false, errors.New("failed to update category")
		}
	} else {
		success = insertCategory(client, category, new_id, categoryCollection, categoryListCollection)
		if !success {
			return false, errors.New("failed to insert into category")
		}
	}
	return true, nil
}

func getTextFromAllEntriesWithCategory(client *mongo.Client, category string, start int, numberToCall int, categoryCollection string, entryCollection string) []string {
	allEntriesWithCategory := getAllEntriesWithCategory(client, category, categoryCollection, entryCollection)
	end := start + numberToCall
	if end > len(allEntriesWithCategory) {
		end = len(allEntriesWithCategory)
	}
	if start > len(allEntriesWithCategory) {
		return []string{}
	}
	allEntriesWithCategory = allEntriesWithCategory[start:end]
	var AllEntriesText []string
	for _, v := range allEntriesWithCategory {
		AllEntriesText = append(AllEntriesText, v.Text)
	}
	return AllEntriesText
}

func GetTextFromRandomEntry(client *mongo.Client, entryCollection string) string {
	randomEntry := getRandomEntry(client, entryCollection)
	return randomEntry.Text
}
