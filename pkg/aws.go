package pkg

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Datasource - A source to DynamoDB
type Datasource struct {
	Region    string
	AccessKey string
	SecretKey string
	Endpoint  string
}

// AWSConfig - The config for AWS
type AWSConfig struct {
	Region    string
	AccessKey string
	SecretKey string
}

func newAWSClient(cfg *AWSConfig) Datasource {
	dynamoVendorSchedules := Datasource{
		Region:    cfg.Region,
		AccessKey: cfg.AccessKey,
		SecretKey: cfg.SecretKey,
	}
	return dynamoVendorSchedules
}

// NewDynamoDB - creates a new instance of DynamoDB
func NewDynamoDB(cfg Config) (*dynamodb.DynamoDB, error) {
	awscfg := &AWSConfig{
		Region:    cfg.GetRegion(),
		AccessKey: cfg.GetAccessKey(),
		SecretKey: cfg.GetSecretKey(),
	}
	c := newAWSClient(awscfg)
	sess, err := c.GetSession()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

// GetSession - get the session for DynamoDB
func (d *Datasource) GetSession() (*session.Session, error) {
	cfg := &aws.Config{Region: aws.String(d.Region)}
	if d.SecretKey != "" {
		cfg = cfg.WithCredentials(credentials.NewCredentials(d))
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}
	return sess, nil

}

// Retrieve - Get the credentials
func (d *Datasource) Retrieve() (credentials.Value, error) {
	return credentials.Value{AccessKeyID: d.AccessKey, SecretAccessKey: d.SecretKey}, nil
}

//IsExpired - to fullfill credentials interface
func (d *Datasource) IsExpired() bool {
	return false
}
