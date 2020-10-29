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

// Returns the access point policy associated with the specified access point. The
// following actions are related to GetAccessPointPolicy:
//
//     *
// PutAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
//
//
// * DeleteAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html)
func (c *Client) GetAccessPointPolicy(ctx context.Context, params *GetAccessPointPolicyInput, optFns ...func(*Options)) (*GetAccessPointPolicyOutput, error) {
	if params == nil {
		params = &GetAccessPointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessPointPolicy", params, optFns, addOperationGetAccessPointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointPolicyInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose policy you want to retrieve. For Amazon S3 on
	// Outposts specify the ARN of the access point accessed in the format
	// arn:aws:s3-outposts:::outpost//accesspoint/. For example, to access the access
	// point reports-ap through outpost my-outpost owned by account 123456789012 in
	// Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap.
	// The value must be URL encoded.
	//
	// This member is required.
	Name *string
}

type GetAccessPointPolicyOutput struct {

	// The access point policy associated with the specified access point.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAccessPointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPointPolicy{}, middleware.After)
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
	addEndpointPrefix_opGetAccessPointPolicyMiddleware(stack)
	addOpGetAccessPointPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPointPolicy(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

type endpointPrefix_opGetAccessPointPolicyMiddleware struct {
}

func (*endpointPrefix_opGetAccessPointPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessPointPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetAccessPointPolicyInput)
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
func addEndpointPrefix_opGetAccessPointPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetAccessPointPolicyMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opGetAccessPointPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetAccessPointPolicy",
	}
}

func (in GetAccessPointPolicyInput) getARNMemberValue() (*string, bool) {
	if in.Name == nil {
		return nil, false
	}
	return in.Name, false
}
func (in GetAccessPointPolicyInput) updateARNMemberValue(v string) GetAccessPointPolicyInput {
	in.Name = &v
	return in
}
func (in GetAccessPointPolicyInput) backfillAccountID(v string) (GetAccessPointPolicyInput, error) {
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return in, fmt.Errorf("error backfilling account id")
		}
		return in, nil
	}
	in.AccountId = &v
	return in, nil
}
