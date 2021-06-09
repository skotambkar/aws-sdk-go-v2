// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the health status of the specified instance. For more information, see
// Health checks for Auto Scaling instances
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html) in the
// Amazon EC2 Auto Scaling User Guide.
func (c *Client) SetInstanceHealth(ctx context.Context, params *SetInstanceHealthInput, optFns ...func(*Options)) (*SetInstanceHealthOutput, error) {
	if params == nil {
		params = &SetInstanceHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetInstanceHealth", params, optFns, c.addOperationSetInstanceHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetInstanceHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetInstanceHealthInput struct {

	// The health status of the instance. Set to Healthy to have the instance remain in
	// service. Set to Unhealthy to have the instance be out of service. Amazon EC2
	// Auto Scaling terminates and replaces the unhealthy instance.
	//
	// This member is required.
	HealthStatus *string

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// If the Auto Scaling group of the specified instance has a HealthCheckGracePeriod
	// specified for the group, by default, this call respects the grace period. Set
	// this to False, to have the call not respect the grace period associated with the
	// group. For more information about the health check grace period, see
	// CreateAutoScalingGroup
	// (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateAutoScalingGroup.html)
	// in the Amazon EC2 Auto Scaling API Reference.
	ShouldRespectGracePeriod *bool
}

type SetInstanceHealthOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationSetInstanceHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetInstanceHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetInstanceHealth{}, middleware.After)
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
	if err = addOpSetInstanceHealthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetInstanceHealth(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetInstanceHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "SetInstanceHealth",
	}
}
