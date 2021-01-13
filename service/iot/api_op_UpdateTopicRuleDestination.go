// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a topic rule destination. You use this to change the status, endpoint
// URL, or confirmation URL of the destination.
func (c *Client) UpdateTopicRuleDestination(ctx context.Context, params *UpdateTopicRuleDestinationInput, optFns ...func(*Options)) (*UpdateTopicRuleDestinationOutput, error) {
	if params == nil {
		params = &UpdateTopicRuleDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTopicRuleDestination", params, optFns, addOperationUpdateTopicRuleDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTopicRuleDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTopicRuleDestinationInput struct {

	// The ARN of the topic rule destination.
	//
	// This member is required.
	Arn *string

	// The status of the topic rule destination. Valid values are: IN_PROGRESS A topic
	// rule destination was created but has not been confirmed. You can set status to
	// IN_PROGRESS by calling UpdateTopicRuleDestination. Calling
	// UpdateTopicRuleDestination causes a new confirmation challenge to be sent to
	// your confirmation endpoint. ENABLED Confirmation was completed, and traffic to
	// this destination is allowed. You can set status to DISABLED by calling
	// UpdateTopicRuleDestination. DISABLED Confirmation was completed, and traffic to
	// this destination is not allowed. You can set status to ENABLED by calling
	// UpdateTopicRuleDestination. ERROR Confirmation could not be completed, for
	// example if the confirmation timed out. You can call GetTopicRuleDestination for
	// details about the error. You can set status to IN_PROGRESS by calling
	// UpdateTopicRuleDestination. Calling UpdateTopicRuleDestination causes a new
	// confirmation challenge to be sent to your confirmation endpoint.
	//
	// This member is required.
	Status types.TopicRuleDestinationStatus
}

type UpdateTopicRuleDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTopicRuleDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTopicRuleDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTopicRuleDestination{}, middleware.After)
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
	if err = addOpUpdateTopicRuleDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTopicRuleDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTopicRuleDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateTopicRuleDestination",
	}
}
