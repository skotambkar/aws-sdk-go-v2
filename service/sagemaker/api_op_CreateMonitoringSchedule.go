// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a schedule that regularly starts Amazon SageMaker Processing Jobs to
// monitor the data captured for an Amazon SageMaker Endoint.
func (c *Client) CreateMonitoringSchedule(ctx context.Context, params *CreateMonitoringScheduleInput, optFns ...func(*Options)) (*CreateMonitoringScheduleOutput, error) {
	if params == nil {
		params = &CreateMonitoringScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMonitoringSchedule", params, optFns, c.addOperationCreateMonitoringScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMonitoringScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMonitoringScheduleInput struct {

	// The configuration object that specifies the monitoring schedule and defines the
	// monitoring job.
	//
	// This member is required.
	MonitoringScheduleConfig *types.MonitoringScheduleConfig

	// The name of the monitoring schedule. The name must be unique within an AWS
	// Region within an AWS account.
	//
	// This member is required.
	MonitoringScheduleName *string

	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)
	// in the AWS Billing and Cost Management User Guide.
	Tags []types.Tag
}

type CreateMonitoringScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateMonitoringScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMonitoringSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMonitoringSchedule{}, middleware.After)
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
	if err = addOpCreateMonitoringScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMonitoringSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMonitoringSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateMonitoringSchedule",
	}
}
