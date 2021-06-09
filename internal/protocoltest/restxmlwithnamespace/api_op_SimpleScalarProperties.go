// Code generated by smithy-go-codegen DO NOT EDIT.

package restxmlwithnamespace

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxmlwithnamespace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) SimpleScalarProperties(ctx context.Context, params *SimpleScalarPropertiesInput, optFns ...func(*Options)) (*SimpleScalarPropertiesOutput, error) {
	if params == nil {
		params = &SimpleScalarPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SimpleScalarProperties", params, optFns, c.addOperationSimpleScalarPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SimpleScalarPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimpleScalarPropertiesInput struct {
	ByteValue *int8

	DoubleValue *float64

	FalseBooleanValue *bool

	FloatValue *float32

	Foo *string

	IntegerValue *int32

	LongValue *int64

	Nested *types.NestedWithNamespace

	ShortValue *int16

	StringValue *string

	TrueBooleanValue *bool
}

type SimpleScalarPropertiesOutput struct {
	ByteValue *int8

	DoubleValue *float64

	FalseBooleanValue *bool

	FloatValue *float32

	Foo *string

	IntegerValue *int32

	LongValue *int64

	Nested *types.NestedWithNamespace

	ShortValue *int16

	StringValue *string

	TrueBooleanValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationSimpleScalarPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpSimpleScalarProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpSimpleScalarProperties{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSimpleScalarProperties(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSimpleScalarProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SimpleScalarProperties",
	}
}
