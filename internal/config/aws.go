package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
	"os"
)

type AwsConfig struct {
	awsRegion    string
	awsAccessKey string
	awsSecretKey string
}

func GetDefaultConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Panicln(err)
	}
	return cfg
}

func GetKeyFromSecretManager(cfg aws.Config) *secretsmanager.GetSecretValueOutput {
	client := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(os.Getenv("AWS_SECRET_NAME")),
		VersionStage: aws.String("AWSCURRENT"),
	}
	result, err := client.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Panicf("Failed to get secret value : %v", err)
	}
	return result
}
