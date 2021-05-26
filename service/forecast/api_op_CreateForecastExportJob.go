// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports a forecast created by the CreateForecast operation to your Amazon Simple
// Storage Service (Amazon S3) bucket. The forecast file name will match the
// following conventions: __ where the component is in Java SimpleDateFormat
// (yyyy-MM-ddTHH-mm-ssZ). You must specify a DataDestination object that includes
// an AWS Identity and Access Management (IAM) role that Amazon Forecast can assume
// to access the Amazon S3 bucket. For more information, see
// aws-forecast-iam-roles. For more information, see howitworks-forecast. To get a
// list of all your forecast export jobs, use the ListForecastExportJobs operation.
// The Status of the forecast export job must be ACTIVE before you can access the
// forecast in your Amazon S3 bucket. To get the status, use the
// DescribeForecastExportJob operation.
func (c *Client) CreateForecastExportJob(ctx context.Context, params *CreateForecastExportJobInput, optFns ...func(*Options)) (*CreateForecastExportJobOutput, error) {
	if params == nil {
		params = &CreateForecastExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateForecastExportJob", params, optFns, c.addOperationCreateForecastExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateForecastExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateForecastExportJobInput struct {

	// The location where you want to save the forecast and an AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the location.
	// The forecast must be exported to an Amazon S3 bucket. If encryption is used,
	// Destination must include an AWS Key Management Service (KMS) key. The IAM role
	// must allow Amazon Forecast permission to access the key.
	//
	// This member is required.
	Destination *types.DataDestination

	// The Amazon Resource Name (ARN) of the forecast that you want to export.
	//
	// This member is required.
	ForecastArn *string

	// The name for the forecast export job.
	//
	// This member is required.
	ForecastExportJobName *string

	// The optional metadata that you apply to the forecast export job to help you
	// categorize and organize them. Each tag consists of a key and an optional value,
	// both of which you define. The following basic restrictions apply to tags:
	//
	// *
	// Maximum number of tags per resource - 50.
	//
	// * For each resource, each tag key
	// must be unique, and each tag key can have only one value.
	//
	// * Maximum key length
	// - 128 Unicode characters in UTF-8.
	//
	// * Maximum value length - 256 Unicode
	// characters in UTF-8.
	//
	// * If your tagging schema is used across multiple services
	// and resources, remember that other services may have restrictions on allowed
	// characters. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	// * Tag
	// keys and values are case sensitive.
	//
	// * Do not use aws:, AWS:, or any upper or
	// lowercase combination of such as a prefix for keys as it is reserved for AWS
	// use. You cannot edit or delete tag keys with this prefix. Values can have this
	// prefix. If a tag value has aws as its prefix but the key does not, then Forecast
	// considers it to be a user tag and will count against the limit of 50 tags. Tags
	// with only the key prefix of aws do not count against your tags per resource
	// limit.
	Tags []types.Tag
}

type CreateForecastExportJobOutput struct {

	// The Amazon Resource Name (ARN) of the export job.
	ForecastExportJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateForecastExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateForecastExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateForecastExportJob{}, middleware.After)
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
	if err = addOpCreateForecastExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateForecastExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateForecastExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "CreateForecastExportJob",
	}
}
