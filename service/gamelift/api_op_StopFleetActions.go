// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Suspends certain types of activity in a fleet location. Currently, this
// operation is used to stop auto-scaling activity. For multi-location fleets,
// fleet actions are managed separately for each location. Stopping fleet actions
// has several potential purposes. It allows you to temporarily stop auto-scaling
// activity but retain your scaling policies for use in the future. For
// multi-location fleets, you can set up fleet-wide auto-scaling, and then opt out
// of it for certain locations. This operation can be used in the following
// ways:
//
// * To stop actions on instances in the fleet's home Region, provide a
// fleet ID and the type of actions to suspend.
//
// * To stop actions on instances in
// one of the fleet's remote locations, provide a fleet ID, a location name, and
// the type of actions to suspend.
//
// If successful, GameLift no longer initiates
// scaling events except in response to manual changes using UpdateFleetCapacity.
// You can view a fleet's stopped actions using DescribeFleetAttributes or
// DescribeFleetLocationAttributes. Suspended activity can be restarted using
// StartFleetActions. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related actions CreateFleet | UpdateFleetCapacity | PutScalingPolicy |
// DescribeEC2InstanceLimits | DescribeFleetAttributes |
// DescribeFleetLocationAttributes | UpdateFleetAttributes | StopFleetActions |
// DeleteFleet | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) StopFleetActions(ctx context.Context, params *StopFleetActionsInput, optFns ...func(*Options)) (*StopFleetActionsOutput, error) {
	if params == nil {
		params = &StopFleetActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopFleetActions", params, optFns, c.addOperationStopFleetActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopFleetActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type StopFleetActionsInput struct {

	// List of actions to suspend on the fleet.
	//
	// This member is required.
	Actions []types.FleetAction

	// A unique identifier for the fleet to stop actions on. You can use either the
	// fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The fleet location to stop fleet actions for. Specify a location in the form of
	// an AWS Region code, such as us-west-2.
	Location *string
}

// Represents the input for a request operation.
type StopFleetActionsOutput struct {

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is
	// assigned to a GameLift fleet resource and uniquely identifies it. ARNs are
	// unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912.
	FleetArn *string

	// A unique identifier for the fleet to stop actions on.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationStopFleetActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopFleetActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopFleetActions{}, middleware.After)
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
	if err = addOpStopFleetActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopFleetActions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopFleetActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StopFleetActions",
	}
}
