// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new FeatureGroup. A FeatureGroup is a group of Features defined in the
// FeatureStore to describe a Record. The FeatureGroup defines the schema and
// features contained in the FeatureGroup. A FeatureGroup definition is composed of
// a list of Features, a RecordIdentifierFeatureName, an EventTimeFeatureName and
// configurations for its OnlineStore and OfflineStore. Check AWS service quotas
// (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html) to see
// the FeatureGroups quota for your AWS account. You must include at least one of
// OnlineStoreConfig and OfflineStoreConfig to create a FeatureGroup.
func (c *Client) CreateFeatureGroup(ctx context.Context, params *CreateFeatureGroupInput, optFns ...func(*Options)) (*CreateFeatureGroupOutput, error) {
	if params == nil {
		params = &CreateFeatureGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFeatureGroup", params, optFns, c.addOperationCreateFeatureGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFeatureGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFeatureGroupInput struct {

	// The name of the feature that stores the EventTime of a Record in a FeatureGroup.
	// An EventTime is a point in time when a new event occurs that corresponds to the
	// creation or update of a Record in a FeatureGroup. All Records in the
	// FeatureGroup must have a corresponding EventTime. An EventTime can be a String
	// or Fractional.
	//
	// * Fractional: EventTime feature values must be a Unix timestamp
	// in seconds.
	//
	// * String: EventTime feature values must be an ISO-8601 string in
	// the format. The following formats are supported yyyy-MM-dd'T'HH:mm:ssZ and
	// yyyy-MM-dd'T'HH:mm:ss.SSSZ where yyyy, MM, and dd represent the year, month, and
	// day respectively and HH, mm, ss, and if applicable, SSS represent the hour,
	// month, second and milliseconds respsectively. 'T' and Z are constants.
	//
	// This member is required.
	EventTimeFeatureName *string

	// A list of Feature names and types. Name and Type is compulsory per Feature.
	// Valid feature FeatureTypes are Integral, Fractional and String. FeatureNames
	// cannot be any of the following: is_deleted, write_time, api_invocation_time You
	// can create up to 2,500 FeatureDefinitions per FeatureGroup.
	//
	// This member is required.
	FeatureDefinitions []types.FeatureDefinition

	// The name of the FeatureGroup. The name must be unique within an AWS Region in an
	// AWS account. The name:
	//
	// * Must start and end with an alphanumeric character.
	//
	// *
	// Can only contain alphanumeric character and hyphens. Spaces are not allowed.
	//
	// This member is required.
	FeatureGroupName *string

	// The name of the Feature whose value uniquely identifies a Record defined in the
	// FeatureStore. Only the latest record per identifier value will be stored in the
	// OnlineStore. RecordIdentifierFeatureName must be one of feature definitions'
	// names. You use the RecordIdentifierFeatureName to access data in a FeatureStore.
	// This name:
	//
	// * Must start and end with an alphanumeric character.
	//
	// * Can only
	// contains alphanumeric characters, hyphens, underscores. Spaces are not allowed.
	//
	// This member is required.
	RecordIdentifierFeatureName *string

	// A free-form description of a FeatureGroup.
	Description *string

	// Use this to configure an OfflineFeatureStore. This parameter allows you to
	// specify:
	//
	// * The Amazon Simple Storage Service (Amazon S3) location of an
	// OfflineStore.
	//
	// * A configuration for an AWS Glue or AWS Hive data cataolgue.
	//
	// *
	// An KMS encryption key to encrypt the Amazon S3 location used for
	// OfflineStore.
	//
	// To learn more about this parameter, see OfflineStoreConfig.
	OfflineStoreConfig *types.OfflineStoreConfig

	// You can turn the OnlineStore on or off by specifying True for the
	// EnableOnlineStore flag in OnlineStoreConfig; the default value is False. You can
	// also include an AWS KMS key ID (KMSKeyId) for at-rest encryption of the
	// OnlineStore.
	OnlineStoreConfig *types.OnlineStoreConfig

	// The Amazon Resource Name (ARN) of the IAM execution role used to persist data
	// into the OfflineStore if an OfflineStoreConfig is provided.
	RoleArn *string

	// Tags used to identify Features in each FeatureGroup.
	Tags []types.Tag
}

type CreateFeatureGroupOutput struct {

	// The Amazon Resource Name (ARN) of the FeatureGroup. This is a unique identifier
	// for the feature group.
	//
	// This member is required.
	FeatureGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateFeatureGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFeatureGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFeatureGroup{}, middleware.After)
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
	if err = addOpCreateFeatureGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFeatureGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFeatureGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateFeatureGroup",
	}
}
