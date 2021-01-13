// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops an asynchronous batch translation job that is in progress. If the job's
// state is IN_PROGRESS, the job will be marked for termination and put into the
// STOP_REQUESTED state. If the job completes before it can be stopped, it is put
// into the COMPLETED state. Otherwise, the job is put into the STOPPED state.
// Asynchronous batch translation jobs are started with the StartTextTranslationJob
// operation. You can use the DescribeTextTranslationJob or ListTextTranslationJobs
// operations to get a batch translation job's JobId.
func (c *Client) StopTextTranslationJob(ctx context.Context, params *StopTextTranslationJobInput, optFns ...func(*Options)) (*StopTextTranslationJobOutput, error) {
	if params == nil {
		params = &StopTextTranslationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopTextTranslationJob", params, optFns, addOperationStopTextTranslationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopTextTranslationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopTextTranslationJobInput struct {

	// The job ID of the job to be stopped.
	//
	// This member is required.
	JobId *string
}

type StopTextTranslationJobOutput struct {

	// The job ID of the stopped batch translation job.
	JobId *string

	// The status of the designated job. Upon successful completion, the job's status
	// will be STOPPED.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopTextTranslationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopTextTranslationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopTextTranslationJob{}, middleware.After)
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
	if err = addOpStopTextTranslationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopTextTranslationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopTextTranslationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "StopTextTranslationJob",
	}
}
