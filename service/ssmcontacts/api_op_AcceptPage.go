// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to acknowledge an engagement to a contact channel during an incident.
func (c *Client) AcceptPage(ctx context.Context, params *AcceptPageInput, optFns ...func(*Options)) (*AcceptPageOutput, error) {
	if params == nil {
		params = &AcceptPageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptPage", params, optFns, c.addOperationAcceptPageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptPageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptPageInput struct {

	// The accept code is a 6-digit code used to acknowledge the page.
	//
	// This member is required.
	AcceptCode *string

	// The type indicates if the page was DELIVERED or READ.
	//
	// This member is required.
	AcceptType types.AcceptType

	// The Amazon Resource Name (ARN) of the engagement to a contact channel.
	//
	// This member is required.
	PageId *string

	// The ARN of the contact channel.
	ContactChannelId *string

	// Information provided by the user when the user acknowledges the page.
	Note *string
}

type AcceptPageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAcceptPageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptPage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptPage{}, middleware.After)
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
	if err = addOpAcceptPageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptPage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptPage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "AcceptPage",
	}
}
