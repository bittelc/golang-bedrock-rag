package aws

import (
	"context"
	"fmt"

	awsDep "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
)

func BedrockInit(ctx *context.Context, cfg *awsDep.Config) {
	bedrockClient := bedrock.NewFromConfig(*cfg)
	result, err := bedrockClient.ListFoundationModels(*ctx, &bedrock.ListFoundationModelsInput{})
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

}
