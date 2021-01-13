// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an AWS DMS event notification subscription. You can specify the type of
// source (SourceType) you want to be notified of, provide a list of AWS DMS source
// IDs (SourceIds) that triggers the events, and provide a list of event categories
// (EventCategories) for events you want to be notified of. If you specify both the
// SourceType and SourceIds, such as SourceType = replication-instance and
// SourceIdentifier = my-replinstance, you will be notified of all the replication
// instance events for the specified source. If you specify a SourceType but don't
// specify a SourceIdentifier, you receive notice of the events for that source
// type for all your AWS DMS sources. If you don't specify either SourceType nor
// SourceIdentifier, you will be notified of events generated from all AWS DMS
// sources belonging to your customer account. For more information about AWS DMS
// events, see Working with Events and Notifications
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html) in the AWS
// Database Migration Service User Guide.
func (c *Client) CreateEventSubscription(ctx context.Context, params *CreateEventSubscriptionInput, optFns ...func(*Options)) (*CreateEventSubscriptionOutput, error) {
	if params == nil {
		params = &CreateEventSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventSubscription", params, optFns, addOperationCreateEventSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateEventSubscriptionInput struct {

	// The Amazon Resource Name (ARN) of the Amazon SNS topic created for event
	// notification. The ARN is created by Amazon SNS when you create a topic and
	// subscribe to it.
	//
	// This member is required.
	SnsTopicArn *string

	// The name of the AWS DMS event notification subscription. This name must be less
	// than 255 characters.
	//
	// This member is required.
	SubscriptionName *string

	// A Boolean value; set to true to activate the subscription, or set to false to
	// create the subscription but not activate it.
	Enabled *bool

	// A list of event categories for a source type that you want to subscribe to. For
	// more information, see Working with Events and Notifications
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html) in the AWS
	// Database Migration Service User Guide.
	EventCategories []string

	// A list of identifiers for which AWS DMS provides notification events. If you
	// don't specify a value, notifications are provided for all sources. If you
	// specify multiple values, they must be of the same type. For example, if you
	// specify a database instance ID, then all of the other values must be database
	// instance IDs.
	SourceIds []string

	// The type of AWS DMS resource that generates the events. For example, if you want
	// to be notified of events generated by a replication instance, you set this
	// parameter to replication-instance. If this value isn't specified, all events are
	// returned. Valid values: replication-instance | replication-task
	SourceType *string

	// One or more tags to be assigned to the event subscription.
	Tags []types.Tag
}

//
type CreateEventSubscriptionOutput struct {

	// The event subscription that was created.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEventSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventSubscription{}, middleware.After)
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
	if err = addOpCreateEventSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEventSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "CreateEventSubscription",
	}
}
