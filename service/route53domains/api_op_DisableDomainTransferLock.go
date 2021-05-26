// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation removes the transfer lock on the domain (specifically the
// clientTransferProhibited status) to allow domain transfers. We recommend you
// refrain from performing this action unless you intend to transfer the domain to
// a different registrar. Successful submission returns an operation ID that you
// can use to track the progress and completion of the action. If the request is
// not completed successfully, the domain registrant will be notified by email.
func (c *Client) DisableDomainTransferLock(ctx context.Context, params *DisableDomainTransferLockInput, optFns ...func(*Options)) (*DisableDomainTransferLockOutput, error) {
	if params == nil {
		params = &DisableDomainTransferLockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableDomainTransferLock", params, optFns, c.addOperationDisableDomainTransferLockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableDomainTransferLockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The DisableDomainTransferLock request includes the following element.
type DisableDomainTransferLockInput struct {

	// The name of the domain that you want to remove the transfer lock for.
	//
	// This member is required.
	DomainName *string
}

// The DisableDomainTransferLock response includes the following element.
type DisableDomainTransferLockOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDisableDomainTransferLockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableDomainTransferLock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableDomainTransferLock{}, middleware.After)
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
	if err = addOpDisableDomainTransferLockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableDomainTransferLock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableDomainTransferLock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "DisableDomainTransferLock",
	}
}
