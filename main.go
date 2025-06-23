package main

import (
	"context"
	"log"

	"golang-bedrock-rag/aws"
	"golang-bedrock-rag/chunk"
	"golang-bedrock-rag/cli"
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

	_, err = chunk.ChunkDoc(args.Filename)
	if err != nil {
		log.Fatalf("could not chunk document: %v", err)
	}

	aws.BedrockInit(&ctx, &cfg)
}
