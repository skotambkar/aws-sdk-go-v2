// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Synchronizes the specified MFA device with its IAM resource object on the AWS
// servers. For more information about creating and working with virtual MFA
// devices, see Using a virtual MFA device
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html) in the
// IAM User Guide.
func (c *Client) ResyncMFADevice(ctx context.Context, params *ResyncMFADeviceInput, optFns ...func(*Options)) (*ResyncMFADeviceOutput, error) {
	if params == nil {
		params = &ResyncMFADeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResyncMFADevice", params, optFns, c.addOperationResyncMFADeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResyncMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResyncMFADeviceInput struct {

	// An authentication code emitted by the device. The format for this parameter is a
	// sequence of six digits.
	//
	// This member is required.
	AuthenticationCode1 *string

	// A subsequent authentication code emitted by the device. The format for this
	// parameter is a sequence of six digits.
	//
	// This member is required.
	AuthenticationCode2 *string

	// Serial number that uniquely identifies the MFA device. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	SerialNumber *string

	// The name of the user whose MFA device you want to resynchronize. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

type ResyncMFADeviceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationResyncMFADeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResyncMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResyncMFADevice{}, middleware.After)
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
	if err = addOpResyncMFADeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResyncMFADevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResyncMFADevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ResyncMFADevice",
	}
}
