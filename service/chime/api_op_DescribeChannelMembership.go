// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the full details of a user's channel membership. The x-amz-chime-bearer
// request header is mandatory. Use the AppInstanceUserArn of the user that makes
// the API call as the value in the header.
func (c *Client) DescribeChannelMembership(ctx context.Context, params *DescribeChannelMembershipInput, optFns ...func(*Options)) (*DescribeChannelMembershipOutput, error) {
	if params == nil {
		params = &DescribeChannelMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChannelMembership", params, optFns, c.addOperationDescribeChannelMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChannelMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChannelMembershipInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the member.
	//
	// This member is required.
	MemberArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string
}

type DescribeChannelMembershipOutput struct {

	// The details of the membership.
	ChannelMembership *types.ChannelMembership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeChannelMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChannelMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChannelMembership{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeChannelMembershipMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeChannelMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChannelMembership(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeChannelMembershipMiddleware struct {
}

func (*endpointPrefix_opDescribeChannelMembershipMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeChannelMembershipMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeChannelMembershipMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeChannelMembershipMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeChannelMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DescribeChannelMembership",
	}
}
