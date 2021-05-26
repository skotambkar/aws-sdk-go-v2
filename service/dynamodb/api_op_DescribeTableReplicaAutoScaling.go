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

// Describes auto scaling settings across replicas of the global table at once.
// This operation only applies to Version 2019.11.21
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
// of global tables.
func (c *Client) DescribeTableReplicaAutoScaling(ctx context.Context, params *DescribeTableReplicaAutoScalingInput, optFns ...func(*Options)) (*DescribeTableReplicaAutoScalingOutput, error) {
	if params == nil {
		params = &DescribeTableReplicaAutoScalingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTableReplicaAutoScaling", params, optFns, c.addOperationDescribeTableReplicaAutoScalingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableReplicaAutoScalingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTableReplicaAutoScalingInput struct {

	// The name of the table.
	//
	// This member is required.
	TableName *string
}

type DescribeTableReplicaAutoScalingOutput struct {

	// Represents the auto scaling properties of the table.
	TableAutoScalingDescription *types.TableAutoScalingDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeTableReplicaAutoScalingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeTableReplicaAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeTableReplicaAutoScaling{}, middleware.After)
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
	if err = addOpDescribeTableReplicaAutoScalingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTableReplicaAutoScaling(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTableReplicaAutoScaling(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "DescribeTableReplicaAutoScaling",
	}
}
