// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the topic detection jobs that you have submitted.
func (c *Client) ListTopicsDetectionJobs(ctx context.Context, params *ListTopicsDetectionJobsInput, optFns ...func(*Options)) (*ListTopicsDetectionJobsOutput, error) {
	if params == nil {
		params = &ListTopicsDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTopicsDetectionJobs", params, optFns, addOperationListTopicsDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTopicsDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTopicsDetectionJobsInput struct {

	// Filters the jobs that are returned. Jobs can be filtered on their name, status,
	// or the date and time that they were submitted. You can set only one filter at a
	// time.
	Filter *types.TopicsDetectionJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListTopicsDetectionJobsOutput struct {

	// Identifies the next page of results to return.
	NextToken *string

	// A list containing the properties of each job that is returned.
	TopicsDetectionJobPropertiesList []types.TopicsDetectionJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTopicsDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTopicsDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTopicsDetectionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTopicsDetectionJobs(options.Region), middleware.Before); err != nil {
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

// ListTopicsDetectionJobsAPIClient is a client that implements the
// ListTopicsDetectionJobs operation.
type ListTopicsDetectionJobsAPIClient interface {
	ListTopicsDetectionJobs(context.Context, *ListTopicsDetectionJobsInput, ...func(*Options)) (*ListTopicsDetectionJobsOutput, error)
}

var _ ListTopicsDetectionJobsAPIClient = (*Client)(nil)

// ListTopicsDetectionJobsPaginatorOptions is the paginator options for
// ListTopicsDetectionJobs
type ListTopicsDetectionJobsPaginatorOptions struct {
	// The maximum number of results to return in each page. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTopicsDetectionJobsPaginator is a paginator for ListTopicsDetectionJobs
type ListTopicsDetectionJobsPaginator struct {
	options   ListTopicsDetectionJobsPaginatorOptions
	client    ListTopicsDetectionJobsAPIClient
	params    *ListTopicsDetectionJobsInput
	nextToken *string
	firstPage bool
}

// NewListTopicsDetectionJobsPaginator returns a new
// ListTopicsDetectionJobsPaginator
func NewListTopicsDetectionJobsPaginator(client ListTopicsDetectionJobsAPIClient, params *ListTopicsDetectionJobsInput, optFns ...func(*ListTopicsDetectionJobsPaginatorOptions)) *ListTopicsDetectionJobsPaginator {
	options := ListTopicsDetectionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTopicsDetectionJobsInput{}
	}

	return &ListTopicsDetectionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTopicsDetectionJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTopicsDetectionJobs page.
func (p *ListTopicsDetectionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTopicsDetectionJobsOutput, error) {
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

	result, err := p.client.ListTopicsDetectionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTopicsDetectionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListTopicsDetectionJobs",
	}
}
