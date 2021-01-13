// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom tape pool. You can use custom tape pool to enable tape
// retention lock on tapes that are archived in the custom pool.
func (c *Client) CreateTapePool(ctx context.Context, params *CreateTapePoolInput, optFns ...func(*Options)) (*CreateTapePoolOutput, error) {
	if params == nil {
		params = &CreateTapePoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTapePool", params, optFns, addOperationCreateTapePoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTapePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTapePoolInput struct {

	// The name of the new custom tape pool.
	//
	// This member is required.
	PoolName *string

	// The storage class that is associated with the new custom pool. When you use your
	// backup application to eject the tape, the tape is archived directly into the
	// storage class (S3 Glacier or S3 Glacier Deep Archive) that corresponds to the
	// pool.
	//
	// This member is required.
	StorageClass types.TapeStorageClass

	// Tape retention lock time is set in days. Tape retention lock can be enabled for
	// up to 100 years (36,500 days).
	RetentionLockTimeInDays *int32

	// Tape retention lock can be configured in two modes. When configured in
	// governance mode, AWS accounts with specific IAM permissions are authorized to
	// remove the tape retention lock from archived virtual tapes. When configured in
	// compliance mode, the tape retention lock cannot be removed by any user,
	// including the root AWS account.
	RetentionLockType types.RetentionLockType

	// A list of up to 50 tags that can be assigned to tape pool. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag
}

type CreateTapePoolOutput struct {

	// The unique Amazon Resource Name (ARN) that represents the custom tape pool. Use
	// the ListTapePools operation to return a list of tape pools for your account and
	// AWS Region.
	PoolARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTapePoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTapePool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTapePool{}, middleware.After)
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
	if err = addOpCreateTapePoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTapePool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTapePool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateTapePool",
	}
}
