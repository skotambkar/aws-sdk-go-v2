// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists job runs based on a set of parameters. A job run is a unit of work, such
// as a Spark jar, PySpark script, or SparkSQL query, that you submit to Amazon EMR
// on EKS.
func (c *Client) ListJobRuns(ctx context.Context, params *ListJobRunsInput, optFns ...func(*Options)) (*ListJobRunsOutput, error) {
	if params == nil {
		params = &ListJobRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobRuns", params, optFns, addOperationListJobRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobRunsInput struct {

	// The ID of the virtual cluster for which to list the job run.
	//
	// This member is required.
	VirtualClusterId *string

	// The date and time after which the job runs were submitted.
	CreatedAfter *time.Time

	// The date and time before which the job runs were submitted.
	CreatedBefore *time.Time

	// The maximum number of job runs that can be listed.
	MaxResults *int32

	// The name of the job run.
	Name *string

	// The token for the next set of job runs to return.
	NextToken *string

	// The states of the job run.
	States []types.JobRunState
}

type ListJobRunsOutput struct {

	// This output lists information about the specified job runs.
	JobRuns []types.JobRun

	// This output displays the token for the next set of job runs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobRuns{}, middleware.After)
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
	if err = addOpListJobRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobRuns(options.Region), middleware.Before); err != nil {
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

// ListJobRunsAPIClient is a client that implements the ListJobRuns operation.
type ListJobRunsAPIClient interface {
	ListJobRuns(context.Context, *ListJobRunsInput, ...func(*Options)) (*ListJobRunsOutput, error)
}

var _ ListJobRunsAPIClient = (*Client)(nil)

// ListJobRunsPaginatorOptions is the paginator options for ListJobRuns
type ListJobRunsPaginatorOptions struct {
	// The maximum number of job runs that can be listed.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobRunsPaginator is a paginator for ListJobRuns
type ListJobRunsPaginator struct {
	options   ListJobRunsPaginatorOptions
	client    ListJobRunsAPIClient
	params    *ListJobRunsInput
	nextToken *string
	firstPage bool
}

// NewListJobRunsPaginator returns a new ListJobRunsPaginator
func NewListJobRunsPaginator(client ListJobRunsAPIClient, params *ListJobRunsInput, optFns ...func(*ListJobRunsPaginatorOptions)) *ListJobRunsPaginator {
	options := ListJobRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListJobRunsInput{}
	}

	return &ListJobRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobRunsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListJobRuns page.
func (p *ListJobRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobRunsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListJobRuns(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListJobRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "ListJobRuns",
	}
}
