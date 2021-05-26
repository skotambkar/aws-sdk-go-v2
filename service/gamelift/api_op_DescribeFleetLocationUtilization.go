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

// Retrieves current usage data for a fleet location. Utilization data provides a
// snapshot of current game hosting activity at the requested location. Use this
// operation to retrieve utilization information for a fleet's remote location or
// home Region (you can also retrieve home Region utilization by calling
// DescribeFleetUtilization). To retrieve utilization data, identify a fleet and
// location. If successful, a FleetUtilization object is returned for the requested
// fleet location. Learn more Setting up GameLift fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)GameLift
// metrics for fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet)
// Related actions CreateFleetLocations | DescribeFleetLocationAttributes |
// DescribeFleetLocationCapacity | DescribeFleetLocationUtilization |
// DescribeFleetAttributes | DescribeFleetCapacity | DescribeFleetUtilization |
// UpdateFleetCapacity | StopFleetActions | DeleteFleetLocations | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribeFleetLocationUtilization(ctx context.Context, params *DescribeFleetLocationUtilizationInput, optFns ...func(*Options)) (*DescribeFleetLocationUtilizationOutput, error) {
	if params == nil {
		params = &DescribeFleetLocationUtilizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetLocationUtilization", params, optFns, c.addOperationDescribeFleetLocationUtilizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetLocationUtilizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeFleetLocationUtilizationInput struct {

	// A unique identifier for the fleet to request location utilization for. You can
	// use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The fleet location to retrieve utilization information for. Specify a location
	// in the form of an AWS Region code, such as us-west-2.
	//
	// This member is required.
	Location *string
}

// Represents the returned data in response to a request operation.
type DescribeFleetLocationUtilizationOutput struct {

	// Utilization information for the requested fleet location. Utilization objects
	// are returned only for fleets and locations that currently exist.
	FleetUtilization *types.FleetUtilization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeFleetLocationUtilizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetLocationUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetLocationUtilization{}, middleware.After)
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
	if err = addOpDescribeFleetLocationUtilizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetLocationUtilization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFleetLocationUtilization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetLocationUtilization",
	}
}
