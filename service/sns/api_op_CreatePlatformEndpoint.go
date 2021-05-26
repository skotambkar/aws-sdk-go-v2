// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS.
// CreatePlatformEndpoint requires the PlatformApplicationArn that is returned from
// CreatePlatformApplication. You can use the returned EndpointArn to send a
// message to a mobile app or by the Subscribe action for subscription to a topic.
// The CreatePlatformEndpoint action is idempotent, so if the requester already
// owns an endpoint with the same device token and attributes, that endpoint's ARN
// is returned without creating a new endpoint. For more information, see Using
// Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). When using
// CreatePlatformEndpoint with Baidu, two attributes must be provided: ChannelId
// and UserId. The token field must also contain the ChannelId. For more
// information, see Creating an Amazon SNS Endpoint for Baidu
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePushBaiduEndpoint.html).
func (c *Client) CreatePlatformEndpoint(ctx context.Context, params *CreatePlatformEndpointInput, optFns ...func(*Options)) (*CreatePlatformEndpointOutput, error) {
	if params == nil {
		params = &CreatePlatformEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlatformEndpoint", params, optFns, c.addOperationCreatePlatformEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlatformEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for CreatePlatformEndpoint action.
type CreatePlatformEndpointInput struct {

	// PlatformApplicationArn returned from CreatePlatformApplication is used to create
	// a an endpoint.
	//
	// This member is required.
	PlatformApplicationArn *string

	// Unique identifier created by the notification service for an app on a device.
	// The specific name for Token will vary, depending on which notification service
	// is being used. For example, when using APNS as the notification service, you
	// need the device token. Alternatively, when using GCM (Firebase Cloud Messaging)
	// or ADM, the device token equivalent is called the registration ID.
	//
	// This member is required.
	Token *string

	// For a list of attributes, see SetEndpointAttributes
	// (https://docs.aws.amazon.com/sns/latest/api/API_SetEndpointAttributes.html).
	Attributes map[string]string

	// Arbitrary user data to associate with the endpoint. Amazon SNS does not use this
	// data. The data must be in UTF-8 format and less than 2KB.
	CustomUserData *string
}

// Response from CreateEndpoint action.
type CreatePlatformEndpointOutput struct {

	// EndpointArn returned from CreateEndpoint action.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreatePlatformEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreatePlatformEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePlatformEndpoint{}, middleware.After)
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
	if err = addOpCreatePlatformEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlatformEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePlatformEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "CreatePlatformEndpoint",
	}
}
