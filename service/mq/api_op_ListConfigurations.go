// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all configurations.
func (c *Client) ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) {
	if params == nil {
		params = &ListConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurations", params, optFns, c.addOperationListConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationsInput struct {

	// The maximum number of configurations that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	MaxResults int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string
}

type ListConfigurationsOutput struct {

	// The list of all revisions for the specified configuration.
	Configurations []types.Configuration

	// The maximum number of configurations that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	MaxResults int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "ListConfigurations",
	}
}
