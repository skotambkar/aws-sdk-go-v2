// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve status of one or more export tasks. You can retrieve the status of up
// to 100 export tasks.
func (c *Client) DescribeExportTasks(ctx context.Context, params *DescribeExportTasksInput, optFns ...func(*Options)) (*DescribeExportTasksOutput, error) {
	if params == nil {
		params = &DescribeExportTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExportTasks", params, optFns, c.addOperationDescribeExportTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportTasksInput struct {

	// One or more unique identifiers used to query the status of an export request.
	ExportIds []string

	// One or more filters.
	//
	// * AgentId - ID of the agent whose collected data will be
	// exported
	Filters []types.ExportFilter

	// The maximum number of volume results returned by DescribeExportTasks in
	// paginated output. When this parameter is used, DescribeExportTasks only returns
	// maxResults results in a single page along with a nextToken response element.
	MaxResults int32

	// The nextToken value returned from a previous paginated DescribeExportTasks
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string
}

type DescribeExportTasksOutput struct {

	// Contains one or more sets of export request details. When the status of a
	// request is SUCCEEDED, the response includes a URL for an Amazon S3 bucket where
	// you can view the data in a CSV file.
	ExportsInfo []types.ExportInfo

	// The nextToken value to include in a future DescribeExportTasks request. When the
	// results of a DescribeExportTasks request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeExportTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExportTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExportTasks{}, middleware.After)
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
	if err = addOpDescribeExportTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportTasks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExportTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeExportTasks",
	}
}
