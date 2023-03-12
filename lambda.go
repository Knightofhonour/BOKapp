package main

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

func createEntry(client *mongo.Client, text string, category string) (bool, error) {
	success, new_id := insertEntry(client, text)
	if !success {
		return false, errors.New("failed to create entry")
	}
	categories := getAllCategories(client)
	categoryAlreadyPresent := inStringArray(categories, category)
	if categoryAlreadyPresent {
		success = updateCategory(client, category, new_id)
		if !success {
			return false, errors.New("failed to update category")
		}
	} else {
		success = insertCategory(client, category, new_id)
		if !success {
			return false, errors.New("failed to insert into category")
		}
	}
	return true, nil
}

func getTextFromAllEntriesWithCategory(client *mongo.Client, category string, start int, numberToCall int) []string {
	allEntriesWithCategory := getAllEntriesWithCategory(client, category)
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

func getTextFromRandomEntry(client *mongo.Client) string {
	randomEntry := getRandomEntry(client)
	return randomEntry.Text
}
