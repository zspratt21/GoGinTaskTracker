package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func NewSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_TOKEN")),
		Endpoint:    aws.String(os.Getenv("DYNAMO_ENDPOINT")),
		Region:      aws.String(os.Getenv("AWS_REGION")),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
