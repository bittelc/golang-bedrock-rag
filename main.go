package main

import (
	"context"
	"fmt"
	"log"

	"golang-bedrock-rag/aws"
	"golang-bedrock-rag/chunk"
	"golang-bedrock-rag/cli"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
)

// Configuration placeholders
const (
	EmbeddingModelID  = "cohere.english-small-v3"
	FoundationModelID = "anthropic.claude-sonnet-4-20250514-v1:0"
)

func main() {
	args, err := cli.GetUserArgs()
	if err != nil {
		log.Fatalf("could not get user args: %v", err)
	}

	ctx := context.Background()
	cfg, err := aws.AuthToAws(&ctx)
	if err != nil {
		log.Fatalf("could not auth to AWS: %v", err)
	}

	bedrockClient := bedrock.NewFromConfig(cfg)
	result, err := bedrockClient.ListFoundationModels(ctx, &bedrock.ListFoundationModelsInput{})
	if err != nil {
		fmt.Printf("Couldn't list foundation models. Here's why: %v\n", err)
		return
	}
	if len(result.ModelSummaries) == 0 {
		fmt.Println("There are no foundation models.")
	}
	for _, modelSummary := range result.ModelSummaries {
		fmt.Println(*modelSummary.ModelId)
	}
	// err = chunk.ChunkDoc("data/short-test-text.docx")
	err = chunk.ChunkDoc(args.Filename)
	if err != nil {
		log.Fatalf("could not chunk document: %v", err)
	}
}
