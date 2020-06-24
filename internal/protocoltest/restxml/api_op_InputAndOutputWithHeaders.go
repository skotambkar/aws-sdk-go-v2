// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests and responses are serialized when there is no
// input or output payload but there are HTTP header bindings.
func (c *Client) InputAndOutputWithHeaders(ctx context.Context, params *InputAndOutputWithHeadersInput, optFns ...func(*Options)) (*InputAndOutputWithHeadersOutput, error) {
	stack := middleware.NewStack("InputAndOutputWithHeaders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "InputAndOutputWithHeaders",
	}, middleware.Before)
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

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
			OperationName: "InputAndOutputWithHeaders",
			Err:           err,
		}
	}
	out := result.(*InputAndOutputWithHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InputAndOutputWithHeadersInput struct {
	HeaderBooleanList   []*bool
	HeaderByte          *int8
	HeaderDouble        *float64
	HeaderEnum          types.FooEnum
	HeaderEnumList      []types.FooEnum
	HeaderFalseBool     *bool
	HeaderFloat         *float32
	HeaderInteger       *int32
	HeaderIntegerList   []*int32
	HeaderLong          *int64
	HeaderShort         *int16
	HeaderString        *string
	HeaderStringList    []*string
	HeaderStringSet     []*string
	HeaderTimestampList []*time.Time
	HeaderTrueBool      *bool
}

type InputAndOutputWithHeadersOutput struct {
	HeaderBooleanList   []*bool
	HeaderByte          *int8
	HeaderDouble        *float64
	HeaderEnum          types.FooEnum
	HeaderEnumList      []types.FooEnum
	HeaderFalseBool     *bool
	HeaderFloat         *float32
	HeaderInteger       *int32
	HeaderIntegerList   []*int32
	HeaderLong          *int64
	HeaderShort         *int16
	HeaderString        *string
	HeaderStringList    []*string
	HeaderStringSet     []*string
	HeaderTimestampList []*time.Time
	HeaderTrueBool      *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
