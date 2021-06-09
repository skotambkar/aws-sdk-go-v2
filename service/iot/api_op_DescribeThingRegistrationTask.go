// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a bulk thing provisioning task.
func (c *Client) DescribeThingRegistrationTask(ctx context.Context, params *DescribeThingRegistrationTaskInput, optFns ...func(*Options)) (*DescribeThingRegistrationTaskOutput, error) {
	if params == nil {
		params = &DescribeThingRegistrationTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeThingRegistrationTask", params, optFns, c.addOperationDescribeThingRegistrationTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeThingRegistrationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeThingRegistrationTaskInput struct {

	// The task ID.
	//
	// This member is required.
	TaskId *string
}

type DescribeThingRegistrationTaskOutput struct {

	// The task creation date.
	CreationDate *time.Time

	// The number of things that failed to be provisioned.
	FailureCount int32

	// The S3 bucket that contains the input file.
	InputFileBucket *string

	// The input file key.
	InputFileKey *string

	// The date when the task was last modified.
	LastModifiedDate *time.Time

	// The message.
	Message *string

	// The progress of the bulk provisioning task expressed as a percentage.
	PercentageProgress int32

	// The role ARN that grants access to the input file bucket.
	RoleArn *string

	// The status of the bulk thing provisioning task.
	Status types.Status

	// The number of things successfully provisioned.
	SuccessCount int32

	// The task ID.
	TaskId *string

	// The task's template.
	TemplateBody *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeThingRegistrationTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeThingRegistrationTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeThingRegistrationTask{}, middleware.After)
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
	if err = addOpDescribeThingRegistrationTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThingRegistrationTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeThingRegistrationTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeThingRegistrationTask",
	}
}
