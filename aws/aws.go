package aws

import (
	"context"
	"log"

	awsDep "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const IAMRoleARN = "arn:aws:iam::742412930455:role/GolangRAG"

func AuthToAws(ctx *context.Context) (awsDep.Config, error) {
	// Load AWS config with IAM role
	baseCfg, err := config.LoadDefaultConfig(*ctx,
		config.WithSharedConfigProfile("LikardaBedrock"),
		config.WithRegion("us-east-1"), // or your region
	)

	if err != nil {
		log.Fatalf("Unable to load base AWS config: %v", err)
	}
	stsClient := sts.NewFromConfig(baseCfg)
	creds := stscreds.NewAssumeRoleProvider(stsClient, IAMRoleARN)
	cfg := baseCfg.Copy()
	cfg.Credentials = awsDep.NewCredentialsCache(creds)

	if _, err := sts.NewFromConfig(cfg).GetCallerIdentity(*ctx, &sts.GetCallerIdentityInput{}); err != nil {
		log.Fatalf("AWS authentication failed: %v", err)
	} else {
		log.Println("✅ AWS authentication successful")
	}
	return cfg, nil
}
