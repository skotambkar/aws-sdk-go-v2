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

// Creates a new event bus within your account. This can be a custom event bus
// which you can use to receive events from your custom applications and services,
// or it can be a partner event bus which can be matched to a partner event source.
func (c *Client) CreateEventBus(ctx context.Context, params *CreateEventBusInput, optFns ...func(*Options)) (*CreateEventBusOutput, error) {
	if params == nil {
		params = &CreateEventBusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventBus", params, optFns, addOperationCreateEventBusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventBusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventBusInput struct {

	// The name of the new event bus. Event bus names cannot contain the / character.
	// You can't use the name default for a custom event bus, as this name is already
	// used for your account's default event bus. If this is a partner event bus, the
	// name must exactly match the name of the partner event source that this event bus
	// is matched to.
	//
	// This member is required.
	Name *string

	// If you are creating a partner event bus, this specifies the partner event source
	// that the new event bus will be matched with.
	EventSourceName *string

	// Tags to associate with the event bus.
	Tags []types.Tag
}

type CreateEventBusOutput struct {

	// The ARN of the new event bus.
	EventBusArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEventBusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventBus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventBus{}, middleware.After)
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
	if err = addOpCreateEventBusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventBus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEventBus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "CreateEventBus",
	}
}
