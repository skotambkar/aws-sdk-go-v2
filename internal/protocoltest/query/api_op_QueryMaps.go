// Code generated by smithy-go-codegen DO NOT EDIT.
package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This test serializes simple and complex maps.
func (c *Client) QueryMaps(ctx context.Context, params *QueryMapsInput, optFns ...func(*Options)) (*QueryMapsOutput, error) {
	stack := middleware.NewStack("QueryMaps", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "QueryMaps",
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
			OperationName: "QueryMaps",
			Err:           err,
		}
	}
	out := result.(*QueryMapsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryMapsInput struct {
	ComplexMapArg           map[string]*types.GreetingStruct
	FlattenedMap            map[string]*string
	FlattenedMapWithXmlName map[string]*string
	MapArg                  map[string]*string
	MapOfLists              map[string][]*string
	MapWithXmlMemberName    map[string]*string
	RenamedMapArg           map[string]*string
}

type QueryMapsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
