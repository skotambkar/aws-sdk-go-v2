// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the status of the specified access key from Active to Inactive, or vice
// versa. This operation can be used to disable a user's key as part of a key
// rotation workflow. If the UserName is not specified, the user name is determined
// implicitly based on the AWS access key ID used to sign the request. This
// operation works for access keys under the AWS account. Consequently, you can use
// this operation to manage AWS account root user credentials even if the AWS
// account has no associated users. For information about rotating keys, see
// Managing Keys and Certificates
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingCredentials.html) in
// the IAM User Guide.
func (c *Client) UpdateAccessKey(ctx context.Context, params *UpdateAccessKeyInput, optFns ...func(*Options)) (*UpdateAccessKeyOutput, error) {
	if params == nil {
		params = &UpdateAccessKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccessKey", params, optFns, addOperationUpdateAccessKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccessKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccessKeyInput struct {

	// The access key ID of the secret access key you want to update. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters that can consist of any upper or lowercased letter or digit.
	//
	// This member is required.
	AccessKeyId *string

	// The status you want to assign to the secret access key. Active means that the
	// key can be used for API calls to AWS, while Inactive means that the key cannot
	// be used.
	//
	// This member is required.
	Status types.StatusType

	// The name of the user whose key you want to update. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	UserName *string
}

type UpdateAccessKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAccessKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateAccessKey{}, middleware.After)
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
	if err = addOpUpdateAccessKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccessKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAccessKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateAccessKey",
	}
}
