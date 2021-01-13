// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a labeling job.
func (c *Client) DescribeLabelingJob(ctx context.Context, params *DescribeLabelingJobInput, optFns ...func(*Options)) (*DescribeLabelingJobOutput, error) {
	if params == nil {
		params = &DescribeLabelingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLabelingJob", params, optFns, addOperationDescribeLabelingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLabelingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLabelingJobInput struct {

	// The name of the labeling job to return information for.
	//
	// This member is required.
	LabelingJobName *string
}

type DescribeLabelingJobOutput struct {

	// The date and time that the labeling job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Configuration information required for human workers to complete a labeling
	// task.
	//
	// This member is required.
	HumanTaskConfig *types.HumanTaskConfig

	// Input configuration information for the labeling job, such as the Amazon S3
	// location of the data objects and the location of the manifest file that
	// describes the data objects.
	//
	// This member is required.
	InputConfig *types.LabelingJobInputConfig

	// A unique identifier for work done as part of a labeling job.
	//
	// This member is required.
	JobReferenceCode *string

	// Provides a breakdown of the number of data objects labeled by humans, the number
	// of objects labeled by machine, the number of objects than couldn't be labeled,
	// and the total number of objects labeled.
	//
	// This member is required.
	LabelCounters *types.LabelCounters

	// The Amazon Resource Name (ARN) of the labeling job.
	//
	// This member is required.
	LabelingJobArn *string

	// The name assigned to the labeling job when it was created.
	//
	// This member is required.
	LabelingJobName *string

	// The processing status of the labeling job.
	//
	// This member is required.
	LabelingJobStatus types.LabelingJobStatus

	// The date and time that the labeling job was last updated.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The location of the job's output data and the AWS Key Management Service key ID
	// for the key used to encrypt the output data, if any.
	//
	// This member is required.
	OutputConfig *types.LabelingJobOutputConfig

	// The Amazon Resource Name (ARN) that Amazon SageMaker assumes to perform tasks on
	// your behalf during data labeling.
	//
	// This member is required.
	RoleArn *string

	// If the job failed, the reason that it failed.
	FailureReason *string

	// The attribute used as the label in the output manifest file.
	LabelAttributeName *string

	// The S3 location of the JSON file that defines the categories used to label data
	// objects. Please note the following label-category limits:
	//
	// * Semantic
	// segmentation labeling jobs using automated labeling: 20 labels
	//
	// * Box bounding
	// labeling jobs (all): 10 labels
	//
	// The file is a JSON structure in the following
	// format: {
	//     "document-version": "2018-11-28"
	//
	//     "labels": [
	//
	//     {
	//
	//
	// "label": "label 1"
	//
	//     },
	//
	//     {
	//
	//     "label": "label 2"
	//
	//     },
	//
	//     ...
	//
	//
	// {
	//
	//     "label": "label n"
	//
	//     }
	//
	//     ]
	//
	//     }
	LabelCategoryConfigS3Uri *string

	// Configuration information for automated data labeling.
	LabelingJobAlgorithmsConfig *types.LabelingJobAlgorithmsConfig

	// The location of the output produced by the labeling job.
	LabelingJobOutput *types.LabelingJobOutput

	// A set of conditions for stopping a labeling job. If any of the conditions are
	// met, the job is automatically stopped.
	StoppingConditions *types.LabelingJobStoppingConditions

	// An array of key-value pairs. You can use tags to categorize your AWS resources
	// in different ways, for example, by purpose, owner, or environment. For more
	// information, see Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLabelingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLabelingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLabelingJob{}, middleware.After)
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
	if err = addOpDescribeLabelingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLabelingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLabelingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeLabelingJob",
	}
}
