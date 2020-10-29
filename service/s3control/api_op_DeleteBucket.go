// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// This API operation deletes an Amazon S3 on Outposts bucket. To delete an S3
// bucket, see DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html) in the
// Amazon Simple Storage Service API. Deletes the Amazon S3 on Outposts bucket. All
// objects (including all object versions and delete markers) in the bucket must be
// deleted before the bucket itself can be deleted. For more information, see Using
// Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in Amazon
// Simple Storage Service Developer Guide. All Amazon S3 on Outposts REST API
// requests for this action require an additional parameter of outpost-id to be
// passed with the request and an S3 on Outposts endpoint hostname prefix instead
// of s3-control. For an example of the request syntax for Amazon S3 on Outposts
// that uses the S3 on Outposts endpoint hostname prefix and the outpost-id derived
// using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucket.html#API_control_DeleteBucket_Examples)
// section below. Related Resources
//
//     * CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html)
//
//
// * GetBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetBucket.html)
//
//
// * DeleteObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)
func (c *Client) DeleteBucket(ctx context.Context, params *DeleteBucketInput, optFns ...func(*Options)) (*DeleteBucketOutput, error) {
	if params == nil {
		params = &DeleteBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucket", params, optFns, addOperationDeleteBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketInput struct {

	// The account ID that owns the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// Specifies the bucket being deleted. For Amazon S3 on Outposts specify the ARN of
	// the bucket accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For
	// example, to access the bucket reports through outpost my-outpost owned by
	// account 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type DeleteBucketOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucket{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addEndpointPrefix_opDeleteBucketMiddleware(stack)
	addOpDeleteBucketValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucket(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

type endpointPrefix_opDeleteBucketMiddleware struct {
}

func (*endpointPrefix_opDeleteBucketMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteBucketMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteBucketMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteBucketMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteBucket(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucket",
	}
}

func (in DeleteBucketInput) getARNMemberValue() (*string, bool) {
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, false
}
func (in DeleteBucketInput) updateARNMemberValue(v string) DeleteBucketInput {
	in.Bucket = &v
	return in
}
func (in DeleteBucketInput) backfillAccountID(v string) (DeleteBucketInput, error) {
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return in, fmt.Errorf("error backfilling account id")
		}
		return in, nil
	}
	in.AccountId = &v
	return in, nil
}
