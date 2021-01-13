// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This example uses fixed query string params and variable query string params.
// The fixed query string parameters and variable parameters must both be
// serialized (implementations may need to merge them together).
func (c *Client) ConstantAndVariableQueryString(ctx context.Context, params *ConstantAndVariableQueryStringInput, optFns ...func(*Options)) (*ConstantAndVariableQueryStringOutput, error) {
	if params == nil {
		params = &ConstantAndVariableQueryStringInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConstantAndVariableQueryString", params, optFns, addOperationConstantAndVariableQueryStringMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConstantAndVariableQueryStringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConstantAndVariableQueryStringInput struct {
	Baz *string

	MaybeSet *string
}

type ConstantAndVariableQueryStringOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationConstantAndVariableQueryStringMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConstantAndVariableQueryString{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConstantAndVariableQueryString{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConstantAndVariableQueryString(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConstantAndVariableQueryString(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConstantAndVariableQueryString",
	}
}
