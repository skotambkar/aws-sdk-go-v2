// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/enums"
)

type ModifyEndpointInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the certificate used for SSL connection.
	CertificateArn *string `type:"string"`

	// The name of the endpoint database.
	DatabaseName *string `type:"string"`

	// The settings in JSON format for the DMS transfer type of source endpoint.
	//
	// Attributes include the following:
	//
	//    * serviceAccessRoleArn - The IAM role that has permission to access the
	//    Amazon S3 bucket.
	//
	//    * BucketName - The name of the S3 bucket to use.
	//
	//    * compressionType - An optional parameter to use GZIP to compress the
	//    target files. Set to NONE (the default) or do not use to leave the files
	//    uncompressed.
	//
	// Shorthand syntax: ServiceAccessRoleArn=string ,BucketName=string,CompressionType=string
	//
	// JSON syntax:
	//
	// { "ServiceAccessRoleArn": "string", "BucketName": "string", "CompressionType":
	// "none"|"gzip" }
	DmsTransferSettings *DmsTransferSettings `type:"structure"`

	// Settings in JSON format for the target Amazon DynamoDB endpoint. For more
	// information about the available settings, see Using Object Mapping to Migrate
	// Data to DynamoDB (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html)
	// in the AWS Database Migration Service User Guide.
	DynamoDbSettings *DynamoDbSettings `type:"structure"`

	// Settings in JSON format for the target Elasticsearch endpoint. For more information
	// about the available settings, see Extra Connection Attributes When Using
	// Elasticsearch as a Target for AWS DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration)
	// in the AWS Database Migration User Guide.
	ElasticsearchSettings *ElasticsearchSettings `type:"structure"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`

	// The database endpoint identifier. Identifiers must begin with a letter; must
	// contain only ASCII letters, digits, and hyphens; and must not end with a
	// hyphen or contain two consecutive hyphens.
	EndpointIdentifier *string `type:"string"`

	// The type of endpoint. Valid values are source and target.
	EndpointType enums.ReplicationEndpointTypeValue `type:"string" enum:"true"`

	// The type of engine for the endpoint. Valid values, depending on the EndpointType,
	// include mysql, oracle, postgres, mariadb, aurora, aurora-postgresql, redshift,
	// s3, db2, azuredb, sybase, dynamodb, mongodb, and sqlserver.
	EngineName *string `type:"string"`

	// The external table definition.
	ExternalTableDefinition *string `type:"string"`

	// Additional attributes associated with the connection. To reset this parameter,
	// pass the empty string ("") as an argument.
	ExtraConnectionAttributes *string `type:"string"`

	// Settings in JSON format for the target Amazon Kinesis Data Streams endpoint.
	// For more information about the available settings, see Using Object Mapping
	// to Migrate Data to a Kinesis Data Stream (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping)
	// in the AWS Database Migration User Guide.
	KinesisSettings *KinesisSettings `type:"structure"`

	// Settings in JSON format for the source MongoDB endpoint. For more information
	// about the available settings, see the configuration properties section in
	// Using MongoDB as a Target for AWS Database Migration Service (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html)
	// in the AWS Database Migration Service User Guide.
	MongoDbSettings *MongoDbSettings `type:"structure"`

	// The password to be used to login to the endpoint database.
	Password *string `type:"string" sensitive:"true"`

	// The port used by the endpoint database.
	Port *int64 `type:"integer"`

	RedshiftSettings *RedshiftSettings `type:"structure"`

	// Settings in JSON format for the target Amazon S3 endpoint. For more information
	// about the available settings, see Extra Connection Attributes When Using
	// Amazon S3 as a Target for AWS DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring)
	// in the AWS Database Migration Service User Guide.
	S3Settings *S3Settings `type:"structure"`

	// The name of the server where the endpoint database resides.
	ServerName *string `type:"string"`

	// The Amazon Resource Name (ARN) for the service access role you want to use
	// to modify the endpoint.
	ServiceAccessRoleArn *string `type:"string"`

	// The SSL mode used to connect to the endpoint. The default value is none.
	SslMode enums.DmsSslModeValue `type:"string" enum:"true"`

	// The user name to be used to login to the endpoint database.
	Username *string `type:"string"`
}

// String returns the string representation
func (s ModifyEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyEndpointInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}
	if s.DynamoDbSettings != nil {
		if err := s.DynamoDbSettings.Validate(); err != nil {
			invalidParams.AddNested("DynamoDbSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.ElasticsearchSettings != nil {
		if err := s.ElasticsearchSettings.Validate(); err != nil {
			invalidParams.AddNested("ElasticsearchSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The modified endpoint.
	Endpoint *Endpoint `type:"structure"`
}

// String returns the string representation
func (s ModifyEndpointOutput) String() string {
	return awsutil.Prettify(s)
}
