// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the attributes for an accelerator.
func (c *Client) UpdateAcceleratorAttributes(ctx context.Context, params *UpdateAcceleratorAttributesInput, optFns ...func(*Options)) (*UpdateAcceleratorAttributesOutput, error) {
	if params == nil {
		params = &UpdateAcceleratorAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAcceleratorAttributes", params, optFns, addOperationUpdateAcceleratorAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAcceleratorAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAcceleratorAttributesInput struct {

	// The Amazon Resource Name (ARN) of the accelerator that you want to update.
	//
	// This member is required.
	AcceleratorArn *string

	// Update whether flow logs are enabled. The default value is false. If the value
	// is true, FlowLogsS3Bucket and FlowLogsS3Prefix must be specified. For more
	// information, see Flow Logs
	// (https://docs.aws.amazon.com/global-accelerator/latest/dg/monitoring-global-accelerator.flow-logs.html)
	// in the AWS Global Accelerator Developer Guide.
	FlowLogsEnabled *bool

	// The name of the Amazon S3 bucket for the flow logs. Attribute is required if
	// FlowLogsEnabled is true. The bucket must exist and have a bucket policy that
	// grants AWS Global Accelerator permission to write to the bucket.
	FlowLogsS3Bucket *string

	// Update the prefix for the location in the Amazon S3 bucket for the flow logs.
	// Attribute is required if FlowLogsEnabled is true. If you don’t specify a prefix,
	// the flow logs are stored in the root of the bucket. If you specify slash (/) for
	// the S3 bucket prefix, the log file bucket folder structure will include a double
	// slash (//), like the following: s3-bucket_name//AWSLogs/aws_account_id
	FlowLogsS3Prefix *string
}

type UpdateAcceleratorAttributesOutput struct {

	// Updated attributes for the accelerator.
	AcceleratorAttributes *types.AcceleratorAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAcceleratorAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAcceleratorAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAcceleratorAttributes{}, middleware.After)
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
	if err = addOpUpdateAcceleratorAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAcceleratorAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAcceleratorAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "UpdateAcceleratorAttributes",
	}
}
