// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a DBProxyEndpoint. Doing so removes the ability to access the DB proxy
// using the endpoint that you defined. The endpoint that you delete might have
// provided capabilities such as read/write or read-only operations, or using a
// different VPC than the DB proxy's default VPC.
func (c *Client) DeleteDBProxyEndpoint(ctx context.Context, params *DeleteDBProxyEndpointInput, optFns ...func(*Options)) (*DeleteDBProxyEndpointOutput, error) {
	if params == nil {
		params = &DeleteDBProxyEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBProxyEndpoint", params, optFns, c.addOperationDeleteDBProxyEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBProxyEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDBProxyEndpointInput struct {

	// The name of the DB proxy endpoint to delete.
	//
	// This member is required.
	DBProxyEndpointName *string
}

type DeleteDBProxyEndpointOutput struct {

	// The data structure representing the details of the DB proxy endpoint that you
	// delete.
	DBProxyEndpoint *types.DBProxyEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteDBProxyEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBProxyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBProxyEndpoint{}, middleware.After)
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
	if err = addOpDeleteDBProxyEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBProxyEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBProxyEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBProxyEndpoint",
	}
}
