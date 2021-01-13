// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your replays. You can either list all the replays or you can provide a
// prefix to match to the replay names. Filter parameters are exclusive.
func (c *Client) ListReplays(ctx context.Context, params *ListReplaysInput, optFns ...func(*Options)) (*ListReplaysOutput, error) {
	if params == nil {
		params = &ListReplaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReplays", params, optFns, addOperationListReplaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReplaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReplaysInput struct {

	// The ARN of the event source associated with the replay.
	EventSourceArn *string

	// The maximum number of replays to retrieve.
	Limit *int32

	// A name prefix to filter the replays returned. Only replays with name that match
	// the prefix are returned.
	NamePrefix *string

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string

	// The state of the replay.
	State types.ReplayState
}

type ListReplaysOutput struct {

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string

	// An array of Replay objects that contain information about the replay.
	Replays []types.Replay

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListReplaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReplays{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReplays{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReplays(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListReplays(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListReplays",
	}
}
