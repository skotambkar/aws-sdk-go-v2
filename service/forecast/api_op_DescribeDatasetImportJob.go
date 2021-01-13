// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a dataset import job created using the CreateDatasetImportJob
// operation. In addition to listing the parameters provided in the
// CreateDatasetImportJob request, this operation includes the following
// properties:
//
// * CreationTime
//
// * LastModificationTime
//
// * DataSize
//
// *
// FieldStatistics
//
// * Status
//
// * Message - If an error occurred, information about
// the error.
func (c *Client) DescribeDatasetImportJob(ctx context.Context, params *DescribeDatasetImportJobInput, optFns ...func(*Options)) (*DescribeDatasetImportJobOutput, error) {
	if params == nil {
		params = &DescribeDatasetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDatasetImportJob", params, optFns, addOperationDescribeDatasetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetImportJobInput struct {

	// The Amazon Resource Name (ARN) of the dataset import job.
	//
	// This member is required.
	DatasetImportJobArn *string
}

type DescribeDatasetImportJobOutput struct {

	// When the dataset import job was created.
	CreationTime *time.Time

	// The size of the dataset in gigabytes (GB) after the import job has finished.
	DataSize *float64

	// The location of the training data to import and an AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the data. If
	// encryption is used, DataSource includes an AWS Key Management Service (KMS) key.
	DataSource *types.DataSource

	// The Amazon Resource Name (ARN) of the dataset that the training data was
	// imported to.
	DatasetArn *string

	// The ARN of the dataset import job.
	DatasetImportJobArn *string

	// The name of the dataset import job.
	DatasetImportJobName *string

	// Statistical information about each field in the input data.
	FieldStatistics map[string]types.Statistics

	GeolocationFormat *string

	// The last time that the dataset was modified. The time depends on the status of
	// the job, as follows:
	//
	// * CREATE_PENDING - The same time as CreationTime.
	//
	// *
	// CREATE_IN_PROGRESS - The current timestamp.
	//
	// * ACTIVE or CREATE_FAILED - When
	// the job finished or failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the dataset import job. The status is reflected in the status of
	// the dataset. For example, when the import job status is CREATE_IN_PROGRESS, the
	// status of the dataset is UPDATE_IN_PROGRESS. States include:
	//
	// * ACTIVE
	//
	// *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	// * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	Status *string

	TimeZone *string

	// The format of timestamps in the dataset. The format that you specify depends on
	// the DataFrequency specified when the dataset was created. The following formats
	// are supported
	//
	// * "yyyy-MM-dd" For the following data frequencies: Y, M, W, and
	// D
	//
	// * "yyyy-MM-dd HH:mm:ss" For the following data frequencies: H, 30min, 15min,
	// and 1min; and optionally, for: Y, M, W, and D
	TimestampFormat *string

	UseGeolocationForTimeZone bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDatasetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDatasetImportJob{}, middleware.After)
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
	if err = addOpDescribeDatasetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatasetImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDatasetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeDatasetImportJob",
	}
}
