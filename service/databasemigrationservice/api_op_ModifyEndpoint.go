// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified endpoint.
func (c *Client) ModifyEndpoint(ctx context.Context, params *ModifyEndpointInput, optFns ...func(*Options)) (*ModifyEndpointOutput, error) {
	if params == nil {
		params = &ModifyEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyEndpoint", params, optFns, addOperationModifyEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyEndpointInput struct {

	// The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.
	//
	// This member is required.
	EndpointArn *string

	// The Amazon Resource Name (ARN) of the certificate used for SSL connection.
	CertificateArn *string

	// The name of the endpoint database.
	DatabaseName *string

	// The settings in JSON format for the DMS transfer type of source endpoint.
	// Attributes include the following:
	//
	// * serviceAccessRoleArn - The AWS Identity and
	// Access Management (IAM) role that has permission to access the Amazon S3
	// bucket.
	//
	// * BucketName - The name of the S3 bucket to use.
	//
	// * compressionType -
	// An optional parameter to use GZIP to compress the target files. Either set this
	// parameter to NONE (the default) or don't use it to leave the files
	// uncompressed.
	//
	// Shorthand syntax for these settings is as follows:
	// ServiceAccessRoleArn=string ,BucketName=string,CompressionType=string JSON
	// syntax for these settings is as follows: { "ServiceAccessRoleArn": "string",
	// "BucketName": "string", "CompressionType": "none"|"gzip" }
	DmsTransferSettings *types.DmsTransferSettings

	// Settings in JSON format for the source DocumentDB endpoint. For more information
	// about the available settings, see the configuration properties section in  Using
	// DocumentDB as a Target for AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html)
	// in the AWS Database Migration Service User Guide.
	DocDbSettings *types.DocDbSettings

	// Settings in JSON format for the target Amazon DynamoDB endpoint. For information
	// about other available settings, see Using Object Mapping to Migrate Data to
	// DynamoDB
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html) in
	// the AWS Database Migration Service User Guide.
	DynamoDbSettings *types.DynamoDbSettings

	// Settings in JSON format for the target Elasticsearch endpoint. For more
	// information about the available settings, see Extra Connection Attributes When
	// Using Elasticsearch as a Target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration)
	// in the AWS Database Migration Service User Guide.
	ElasticsearchSettings *types.ElasticsearchSettings

	// The database endpoint identifier. Identifiers must begin with a letter and must
	// contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or
	// contain two consecutive hyphens.
	EndpointIdentifier *string

	// The type of endpoint. Valid values are source and target.
	EndpointType types.ReplicationEndpointTypeValue

	// The type of engine for the endpoint. Valid values, depending on the
	// EndpointType, include "mysql", "oracle", "postgres", "mariadb", "aurora",
	// "aurora-postgresql", "redshift", "s3", "db2", "azuredb", "sybase", "dynamodb",
	// "mongodb", "kinesis", "kafka", "elasticsearch", "documentdb", "sqlserver", and
	// "neptune".
	EngineName *string

	// The external table definition.
	ExternalTableDefinition *string

	// Additional attributes associated with the connection. To reset this parameter,
	// pass the empty string ("") as an argument.
	ExtraConnectionAttributes *string

	// Settings in JSON format for the source IBM Db2 LUW endpoint. For information
	// about other available settings, see Extra connection attributes when using Db2
	// LUW as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	IBMDb2Settings *types.IBMDb2Settings

	// Settings in JSON format for the target Apache Kafka endpoint. For more
	// information about the available settings, see Using Apache Kafka as a Target for
	// AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html) in the
	// AWS Database Migration Service User Guide.
	KafkaSettings *types.KafkaSettings

	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	// For more information about the available settings, see Using Amazon Kinesis Data
	// Streams as a Target for AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html) in
	// the AWS Database Migration Service User Guide.
	KinesisSettings *types.KinesisSettings

	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	// For information about other available settings, see Extra connection attributes
	// when using SQL Server as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.ConnectionAttrib)
	// and  Extra connection attributes when using SQL Server as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	MicrosoftSQLServerSettings *types.MicrosoftSQLServerSettings

	// Settings in JSON format for the source MongoDB endpoint. For more information
	// about the available settings, see the configuration properties section in  Using
	// MongoDB as a Target for AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html) in
	// the AWS Database Migration Service User Guide.
	MongoDbSettings *types.MongoDbSettings

	// Settings in JSON format for the source and target MySQL endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using MySQL as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.ConnectionAttrib)
	// and Extra connection attributes when using a MySQL-compatible database as a
	// target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	MySQLSettings *types.MySQLSettings

	// Settings in JSON format for the target Amazon Neptune endpoint. For more
	// information about the available settings, see Specifying Endpoint Settings for
	// Amazon Neptune as a Target
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings)
	// in the AWS Database Migration Service User Guide.
	NeptuneSettings *types.NeptuneSettings

	// Settings in JSON format for the source and target Oracle endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using Oracle as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.ConnectionAttrib)
	// and  Extra connection attributes when using Oracle as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	OracleSettings *types.OracleSettings

	// The password to be used to login to the endpoint database.
	Password *string

	// The port used by the endpoint database.
	Port *int32

	// Settings in JSON format for the source and target PostgreSQL endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using PostgreSQL as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.ConnectionAttrib)
	// and  Extra connection attributes when using PostgreSQL as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	PostgreSQLSettings *types.PostgreSQLSettings

	// Provides information that defines an Amazon Redshift endpoint.
	RedshiftSettings *types.RedshiftSettings

	// Settings in JSON format for the target Amazon S3 endpoint. For more information
	// about the available settings, see Extra Connection Attributes When Using Amazon
	// S3 as a Target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring)
	// in the AWS Database Migration Service User Guide.
	S3Settings *types.S3Settings

	// The name of the server where the endpoint database resides.
	ServerName *string

	// The Amazon Resource Name (ARN) for the service access role you want to use to
	// modify the endpoint.
	ServiceAccessRoleArn *string

	// The SSL mode used to connect to the endpoint. The default value is none.
	SslMode types.DmsSslModeValue

	// Settings in JSON format for the source and target SAP ASE endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using SAP ASE as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.ConnectionAttrib)
	// and Extra connection attributes when using SAP ASE as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	SybaseSettings *types.SybaseSettings

	// The user name to be used to login to the endpoint database.
	Username *string
}

//
type ModifyEndpointOutput struct {

	// The modified endpoint.
	Endpoint *types.Endpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyEndpoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ModifyEndpoint",
	}
}
