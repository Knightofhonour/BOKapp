package main

import (
	"context"

	core "github.com/knightofhonour/core"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	client := core.Connect()
	text := core.GetTextFromRandomEntry(client, core.EntryCollection)
	return text, nil
}

func main() {
	lambda.Start(HandleRequest)
}
