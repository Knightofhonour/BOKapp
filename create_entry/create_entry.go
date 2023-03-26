package main

import (
	"context"

	core "github.com/knightofhonour/core"

	"github.com/aws/aws-lambda-go/lambda"
)

type CreateEntryEvent struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}

func HandleRequest(ctx context.Context, event CreateEntryEvent) (bool, error) {
	client := core.Connect()
	success, err := core.CreateEntry(client, event.Text, event.Category, core.EntryCollection, core.CategoryListCollection, core.CategoryCollection)
	return success, err
}

func main() {
	lambda.Start(HandleRequest)
}
