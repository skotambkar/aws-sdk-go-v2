// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example uses a constant query string parameters and a label. This simply
// tests that labels and query string parameters are compatible. The fixed query
// string parameter named "hello" should in no way conflict with the label,
// {hello}.
func (c *Client) ConstantQueryString(ctx context.Context, params *ConstantQueryStringInput, optFns ...func(*Options)) (*ConstantQueryStringOutput, error) {
	stack := middleware.NewStack("ConstantQueryString", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpConstantQueryStringMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpConstantQueryStringValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConstantQueryString(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "ConstantQueryString",
			Err:           err,
		}
	}
	out := result.(*ConstantQueryStringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConstantQueryStringInput struct {
	Hello *string
}

type ConstantQueryStringOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpConstantQueryStringMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpConstantQueryString{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpConstantQueryString{}, middleware.After)
}

func newServiceMetadataMiddleware_opConstantQueryString(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "ConstantQueryString",
	}
}
