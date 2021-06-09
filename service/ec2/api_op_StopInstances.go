// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops an Amazon EBS-backed instance. You can use the Stop action to hibernate an
// instance if the instance is enabled for hibernation
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#enabling-hibernation)
// and it meets the hibernation prerequisites
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#hibernating-prerequisites).
// For more information, see Hibernate your instance
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the
// Amazon EC2 User Guide. We don't charge usage for a stopped instance, or data
// transfer fees; however, your root partition Amazon EBS volume remains and
// continues to persist your data, and you are charged for Amazon EBS volume usage.
// Every time you start your Windows instance, Amazon EC2 charges you for a full
// instance hour. If you stop and restart your Windows instance, a new instance
// hour begins and Amazon EC2 charges you for another full instance hour even if
// you are still within the same 60-minute period when it was stopped. Every time
// you start your Linux instance, Amazon EC2 charges a one-minute minimum for
// instance usage, and thereafter charges per second for instance usage. You can't
// stop or hibernate instance store-backed instances. You can't use the Stop action
// to hibernate Spot Instances, but you can specify that Amazon EC2 should
// hibernate Spot Instances when they are interrupted. For more information, see
// Hibernating interrupted Spot Instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-interruptions.html#hibernate-spot-instances)
// in the Amazon EC2 User Guide. When you stop or hibernate an instance, we shut it
// down. You can restart your instance at any time. Before stopping or hibernating
// an instance, make sure it is in a state from which it can be restarted. Stopping
// an instance does not preserve data stored in RAM, but hibernating an instance
// does preserve data stored in RAM. If an instance cannot hibernate successfully,
// a normal shutdown occurs. Stopping and hibernating an instance is different to
// rebooting or terminating it. For example, when you stop or hibernate an
// instance, the root device and any other devices attached to the instance
// persist. When you terminate an instance, the root device and any other devices
// attached during the instance launch are automatically deleted. For more
// information about the differences between rebooting, stopping, hibernating, and
// terminating instances, see Instance lifecycle
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
// in the Amazon EC2 User Guide. When you stop an instance, we attempt to shut it
// down forcibly after a short while. If your instance appears stuck in the
// stopping state after a period of time, there may be an issue with the underlying
// host computer. For more information, see Troubleshooting stopping your instance
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesStopping.html)
// in the Amazon EC2 User Guide.
func (c *Client) StopInstances(ctx context.Context, params *StopInstancesInput, optFns ...func(*Options)) (*StopInstancesOutput, error) {
	if params == nil {
		params = &StopInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopInstances", params, optFns, c.addOperationStopInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopInstancesInput struct {

	// The IDs of the instances.
	//
	// This member is required.
	InstanceIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Forces the instances to stop. The instances do not have an opportunity to flush
	// file system caches or file system metadata. If you use this option, you must
	// perform file system check and repair procedures. This option is not recommended
	// for Windows instances. Default: false
	Force *bool

	// Hibernates the instance if the instance was enabled for hibernation at launch.
	// If the instance cannot hibernate successfully, a normal shutdown occurs. For
	// more information, see Hibernate your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the
	// Amazon EC2 User Guide. Default: false
	Hibernate *bool
}

type StopInstancesOutput struct {

	// Information about the stopped instances.
	StoppingInstances []types.InstanceStateChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationStopInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpStopInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpStopInstances{}, middleware.After)
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
	if err = addOpStopInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "StopInstances",
	}
}
