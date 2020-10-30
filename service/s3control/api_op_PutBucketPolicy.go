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

// This API action puts a bucket policy to an Amazon S3 on Outposts bucket. To put
// a policy on an S3 bucket, see PutBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html) in
// the Amazon Simple Storage Service API. Applies an Amazon S3 bucket policy to an
// Outposts bucket. For more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
// Amazon Simple Storage Service Developer Guide. If you are using an identity
// other than the root user of the AWS account that owns the Outposts bucket, the
// calling identity must have the PutBucketPolicy permissions on the specified
// Outposts bucket and belong to the bucket owner's account in order to use this
// operation. If you don't have PutBucketPolicy permissions, Amazon S3 returns a
// 403 Access Denied error. If you have the correct permissions, but you're not
// using an identity that belongs to the bucket owner's account, Amazon S3 returns
// a 405 Method Not Allowed error. As a security precaution, the root user of the
// AWS account that owns a bucket can always use this operation, even if the policy
// explicitly denies the root user the ability to perform this action. For more
// information about bucket policies, see Using Bucket Policies and User Policies
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). All
// Amazon S3 on Outposts REST API requests for this action require an additional
// parameter of outpost-id to be passed with the request and an S3 on Outposts
// endpoint hostname prefix instead of s3-control. For an example of the request
// syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname
// prefix and the outpost-id derived using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_PutBucketPolicy.html#API_control_PutBucketPolicy_Examples)
// section below. The following actions are related to PutBucketPolicy:
//
//     *
// GetBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html)
//
//
// * DeleteBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html)
func (c *Client) PutBucketPolicy(ctx context.Context, params *PutBucketPolicyInput, optFns ...func(*Options)) (*PutBucketPolicyOutput, error) {
	if params == nil {
		params = &PutBucketPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketPolicy", params, optFns, addOperationPutBucketPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketPolicyInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The ARN of the bucket. For Amazon S3 on Outposts specify the ARN of the bucket
	// accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For example, to
	// access the bucket reports through outpost my-outpost owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	// The bucket policy as a JSON document.
	//
	// This member is required.
	Policy *string

	// Set this parameter to true to confirm that you want to remove your permissions
	// to change this bucket policy in the future. This is not supported by Amazon S3
	// on Outposts buckets.
	ConfirmRemoveSelfBucketAccess *bool
}

type PutBucketPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketPolicy{}, middleware.After)
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
	addEndpointPrefix_opPutBucketPolicyMiddleware(stack)
	addOpPutBucketPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketPolicy(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	smithyhttp.AddChecksumMiddleware(stack)
	return nil
}

type endpointPrefix_opPutBucketPolicyMiddleware struct {
}

func (*endpointPrefix_opPutBucketPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutBucketPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutBucketPolicyInput)
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
func addEndpointPrefix_opPutBucketPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutBucketPolicyMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opPutBucketPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketPolicy",
	}
}

func (in PutBucketPolicyInput) getARNMemberValue() (*string, bool) {
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func (in PutBucketPolicyInput) updateARNMemberValue(v string) interface{} {
	in.Bucket = &v
	return &in
}
func (in PutBucketPolicyInput) backfillAccountID(v string) (interface{}, error) {
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return &in, fmt.Errorf("error backfilling account id")
		}
		return &in, nil
	}
	in.AccountId = &v
	return &in, nil
}
