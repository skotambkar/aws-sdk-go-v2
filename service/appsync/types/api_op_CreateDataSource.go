// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/appsync/enums"
)

type CreateDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The API ID for the GraphQL API for the DataSource.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// A description of the DataSource.
	Description *string `locationName:"description" type:"string"`

	// Amazon DynamoDB settings.
	DynamodbConfig *DynamodbDataSourceConfig `locationName:"dynamodbConfig" type:"structure"`

	// Amazon Elasticsearch Service settings.
	ElasticsearchConfig *ElasticsearchDataSourceConfig `locationName:"elasticsearchConfig" type:"structure"`

	// HTTP endpoint settings.
	HttpConfig *HttpDataSourceConfig `locationName:"httpConfig" type:"structure"`

	// AWS Lambda settings.
	LambdaConfig *LambdaDataSourceConfig `locationName:"lambdaConfig" type:"structure"`

	// A user-supplied name for the DataSource.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// Relational database settings.
	RelationalDatabaseConfig *RelationalDatabaseDataSourceConfig `locationName:"relationalDatabaseConfig" type:"structure"`

	// The AWS IAM service role ARN for the data source. The system assumes this
	// role when accessing the data source.
	ServiceRoleArn *string `locationName:"serviceRoleArn" type:"string"`

	// The type of the DataSource.
	//
	// Type is a required field
	Type enums.DataSourceType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSourceInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.DynamodbConfig != nil {
		if err := s.DynamodbConfig.Validate(); err != nil {
			invalidParams.AddNested("DynamodbConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.ElasticsearchConfig != nil {
		if err := s.ElasticsearchConfig.Validate(); err != nil {
			invalidParams.AddNested("ElasticsearchConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.HttpConfig != nil {
		if err := s.HttpConfig.Validate(); err != nil {
			invalidParams.AddNested("HttpConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.LambdaConfig != nil {
		if err := s.LambdaConfig.Validate(); err != nil {
			invalidParams.AddNested("LambdaConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The DataSource object.
	DataSource *DataSource `locationName:"dataSource" type:"structure"`
}

// String returns the string representation
func (s CreateDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}
