// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates settings for a global table.
func (c *Client) UpdateGlobalTableSettings(ctx context.Context, params *UpdateGlobalTableSettingsInput, optFns ...func(*Options)) (*UpdateGlobalTableSettingsOutput, error) {
	if params == nil {
		params = &UpdateGlobalTableSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGlobalTableSettings", params, optFns, c.addOperationUpdateGlobalTableSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGlobalTableSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlobalTableSettingsInput struct {

	// The name of the global table
	//
	// This member is required.
	GlobalTableName *string

	// The billing mode of the global table. If GlobalTableBillingMode is not
	// specified, the global table defaults to PROVISIONED capacity billing mode.
	//
	// *
	// PROVISIONED - We recommend using PROVISIONED for predictable workloads.
	// PROVISIONED sets the billing mode to Provisioned Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
	//
	// *
	// PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable
	// workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
	GlobalTableBillingMode types.BillingMode

	// Represents the settings of a global secondary index for a global table that will
	// be modified.
	GlobalTableGlobalSecondaryIndexSettingsUpdate []types.GlobalTableGlobalSecondaryIndexSettingsUpdate

	// Auto scaling settings for managing provisioned write capacity for the global
	// table.
	GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate *types.AutoScalingSettingsUpdate

	// The maximum number of writes consumed per second before DynamoDB returns a
	// ThrottlingException.
	GlobalTableProvisionedWriteCapacityUnits *int64

	// Represents the settings for a global table in a Region that will be modified.
	ReplicaSettingsUpdate []types.ReplicaSettingsUpdate
}

type UpdateGlobalTableSettingsOutput struct {

	// The name of the global table.
	GlobalTableName *string

	// The Region-specific settings for the global table.
	ReplicaSettings []types.ReplicaSettingsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateGlobalTableSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateGlobalTableSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateGlobalTableSettings{}, middleware.After)
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
	if err = addOpUpdateGlobalTableSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlobalTableSettings(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateGlobalTableSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateGlobalTableSettings",
	}
}
