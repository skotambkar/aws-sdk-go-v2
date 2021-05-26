// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a warm pool for the specified Auto Scaling group. A warm pool
// is a pool of pre-initialized EC2 instances that sits alongside the Auto Scaling
// group. Whenever your application needs to scale out, the Auto Scaling group can
// draw on the warm pool to meet its new desired capacity. For more information and
// example configurations, see Warm pools for Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html)
// in the Amazon EC2 Auto Scaling User Guide. This operation must be called from
// the Region in which the Auto Scaling group was created. This operation cannot be
// called on an Auto Scaling group that has a mixed instances policy or a launch
// template or launch configuration that requests Spot Instances. You can view the
// instances in the warm pool using the DescribeWarmPool API call. If you are no
// longer using a warm pool, you can delete it by calling the DeleteWarmPool API.
func (c *Client) PutWarmPool(ctx context.Context, params *PutWarmPoolInput, optFns ...func(*Options)) (*PutWarmPoolOutput, error) {
	if params == nil {
		params = &PutWarmPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutWarmPool", params, optFns, c.addOperationPutWarmPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutWarmPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutWarmPoolInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Specifies the maximum number of instances that are allowed to be in the warm
	// pool or in any state except Terminated for the Auto Scaling group. This is an
	// optional property. Specify it only if you do not want the warm pool size to be
	// determined by the difference between the group's maximum capacity and its
	// desired capacity. If a value for MaxGroupPreparedCapacity is not specified,
	// Amazon EC2 Auto Scaling launches and maintains the difference between the
	// group's maximum capacity and its desired capacity. If you specify a value for
	// MaxGroupPreparedCapacity, Amazon EC2 Auto Scaling uses the difference between
	// the MaxGroupPreparedCapacity and the desired capacity instead. The size of the
	// warm pool is dynamic. Only when MaxGroupPreparedCapacity and MinSize are set to
	// the same value does the warm pool have an absolute size. If the desired capacity
	// of the Auto Scaling group is higher than the MaxGroupPreparedCapacity, the
	// capacity of the warm pool is 0, unless you specify a value for MinSize. To
	// remove a value that you previously set, include the property but specify -1 for
	// the value.
	MaxGroupPreparedCapacity *int32

	// Specifies the minimum number of instances to maintain in the warm pool. This
	// helps you to ensure that there is always a certain number of warmed instances
	// available to handle traffic spikes. Defaults to 0 if not specified.
	MinSize *int32

	// Sets the instance state to transition to after the lifecycle actions are
	// complete. Default is Stopped.
	PoolState types.WarmPoolState
}

type PutWarmPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutWarmPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutWarmPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutWarmPool{}, middleware.After)
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
	if err = addOpPutWarmPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutWarmPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutWarmPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutWarmPool",
	}
}
