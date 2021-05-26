// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a snapshot schedule for a cluster.
func (c *Client) ModifyClusterSnapshotSchedule(ctx context.Context, params *ModifyClusterSnapshotScheduleInput, optFns ...func(*Options)) (*ModifyClusterSnapshotScheduleOutput, error) {
	if params == nil {
		params = &ModifyClusterSnapshotScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClusterSnapshotSchedule", params, optFns, c.addOperationModifyClusterSnapshotScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterSnapshotScheduleInput struct {

	// A unique identifier for the cluster whose snapshot schedule you want to modify.
	//
	// This member is required.
	ClusterIdentifier *string

	// A boolean to indicate whether to remove the assoiciation between the cluster and
	// the schedule.
	DisassociateSchedule *bool

	// A unique alphanumeric identifier for the schedule that you want to associate
	// with the cluster.
	ScheduleIdentifier *string
}

type ModifyClusterSnapshotScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationModifyClusterSnapshotScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterSnapshotSchedule{}, middleware.After)
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
	if err = addOpModifyClusterSnapshotScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterSnapshotSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyClusterSnapshotSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterSnapshotSchedule",
	}
}
