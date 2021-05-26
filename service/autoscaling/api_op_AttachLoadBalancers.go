// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// To attach an Application Load Balancer, Network Load Balancer, or Gateway Load
// Balancer, use the AttachLoadBalancerTargetGroups API operation instead. Attaches
// one or more Classic Load Balancers to the specified Auto Scaling group. Amazon
// EC2 Auto Scaling registers the running instances with these Classic Load
// Balancers. To describe the load balancers for an Auto Scaling group, call the
// DescribeLoadBalancers API. To detach the load balancer from the Auto Scaling
// group, call the DetachLoadBalancers API. For more information, see Elastic Load
// Balancing and Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) AttachLoadBalancers(ctx context.Context, params *AttachLoadBalancersInput, optFns ...func(*Options)) (*AttachLoadBalancersOutput, error) {
	if params == nil {
		params = &AttachLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachLoadBalancers", params, optFns, c.addOperationAttachLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachLoadBalancersInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The names of the load balancers. You can specify up to 10 load balancers.
	//
	// This member is required.
	LoadBalancerNames []string
}

type AttachLoadBalancersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAttachLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAttachLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachLoadBalancers{}, middleware.After)
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
	if err = addOpAttachLoadBalancersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachLoadBalancers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "AttachLoadBalancers",
	}
}
