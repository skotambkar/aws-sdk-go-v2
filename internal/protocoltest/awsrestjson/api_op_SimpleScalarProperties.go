// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) SimpleScalarProperties(ctx context.Context, params *SimpleScalarPropertiesInput, optFns ...func(*Options)) (*SimpleScalarPropertiesOutput, error) {
	stack := middleware.NewStack("SimpleScalarProperties", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSimpleScalarPropertiesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSimpleScalarProperties(options.Region), middleware.Before)

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
			OperationName: "SimpleScalarProperties",
			Err:           err,
		}
	}
	out := result.(*SimpleScalarPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimpleScalarPropertiesInput struct {
	Foo               *string
	StringValue       *string
	TrueBooleanValue  *bool
	FalseBooleanValue *bool
	ByteValue         *int8
	ShortValue        *int16
	IntegerValue      *int32
	LongValue         *int64
	FloatValue        *float32
	DoubleValue       *float64
}

type SimpleScalarPropertiesOutput struct {
	Foo               *string
	StringValue       *string
	TrueBooleanValue  *bool
	FalseBooleanValue *bool
	ByteValue         *int8
	ShortValue        *int16
	IntegerValue      *int32
	LongValue         *int64
	FloatValue        *float32
	DoubleValue       *float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSimpleScalarPropertiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSimpleScalarProperties{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSimpleScalarProperties{}, middleware.After)
}

func newServiceMetadataMiddleware_opSimpleScalarProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "SimpleScalarProperties",
	}
}
