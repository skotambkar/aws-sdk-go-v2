// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds device(s) to your account (i.e., claim one or more devices) if and only if
// you received a claim code with the device(s).
func (c *Client) ClaimDevicesByClaimCode(ctx context.Context, params *ClaimDevicesByClaimCodeInput, optFns ...func(*Options)) (*ClaimDevicesByClaimCodeOutput, error) {
	if params == nil {
		params = &ClaimDevicesByClaimCodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ClaimDevicesByClaimCode", params, optFns, addOperationClaimDevicesByClaimCodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ClaimDevicesByClaimCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ClaimDevicesByClaimCodeInput struct {

	// The claim code, starting with "C-", as provided by the device manufacturer.
	//
	// This member is required.
	ClaimCode *string
}

type ClaimDevicesByClaimCodeOutput struct {

	// The claim code provided by the device manufacturer.
	ClaimCode *string

	// The total number of devices associated with the claim code that has been
	// processed in the claim request.
	Total int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationClaimDevicesByClaimCodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpClaimDevicesByClaimCode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpClaimDevicesByClaimCode{}, middleware.After)
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
	if err = addOpClaimDevicesByClaimCodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opClaimDevicesByClaimCode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opClaimDevicesByClaimCode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "ClaimDevicesByClaimCode",
	}
}
