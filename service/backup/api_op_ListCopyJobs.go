// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata about your copy jobs.
func (c *Client) ListCopyJobs(ctx context.Context, params *ListCopyJobsInput, optFns ...func(*Options)) (*ListCopyJobsOutput, error) {
	if params == nil {
		params = &ListCopyJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCopyJobs", params, optFns, addOperationListCopyJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCopyJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCopyJobsInput struct {

	// The account ID to list the jobs from. Returns only copy jobs associated with the
	// specified account ID.
	ByAccountId *string

	// Returns only copy jobs that were created after the specified date.
	ByCreatedAfter *time.Time

	// Returns only copy jobs that were created before the specified date.
	ByCreatedBefore *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies a source backup vault to
	// copy from; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	ByDestinationVaultArn *string

	// Returns only copy jobs that match the specified resource Amazon Resource Name
	// (ARN).
	ByResourceArn *string

	// Returns only backup jobs for the specified resources:
	//
	// * DynamoDB for Amazon
	// DynamoDB
	//
	// * EBS for Amazon Elastic Block Store
	//
	// * EC2 for Amazon Elastic Compute
	// Cloud
	//
	// * EFS for Amazon Elastic File System
	//
	// * RDS for Amazon Relational
	// Database Service
	//
	// * Storage Gateway for AWS Storage Gateway
	ByResourceType *string

	// Returns only copy jobs that are in the specified state.
	ByState types.CopyJobState

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
}

type ListCopyJobsOutput struct {

	// An array of structures containing metadata about your copy jobs returned in JSON
	// format.
	CopyJobs []types.CopyJob

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCopyJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCopyJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCopyJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCopyJobs(options.Region), middleware.Before); err != nil {
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

// ListCopyJobsAPIClient is a client that implements the ListCopyJobs operation.
type ListCopyJobsAPIClient interface {
	ListCopyJobs(context.Context, *ListCopyJobsInput, ...func(*Options)) (*ListCopyJobsOutput, error)
}

var _ ListCopyJobsAPIClient = (*Client)(nil)

// ListCopyJobsPaginatorOptions is the paginator options for ListCopyJobs
type ListCopyJobsPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCopyJobsPaginator is a paginator for ListCopyJobs
type ListCopyJobsPaginator struct {
	options   ListCopyJobsPaginatorOptions
	client    ListCopyJobsAPIClient
	params    *ListCopyJobsInput
	nextToken *string
	firstPage bool
}

// NewListCopyJobsPaginator returns a new ListCopyJobsPaginator
func NewListCopyJobsPaginator(client ListCopyJobsAPIClient, params *ListCopyJobsInput, optFns ...func(*ListCopyJobsPaginatorOptions)) *ListCopyJobsPaginator {
	options := ListCopyJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListCopyJobsInput{}
	}

	return &ListCopyJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCopyJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCopyJobs page.
func (p *ListCopyJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCopyJobsOutput, error) {
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

	result, err := p.client.ListCopyJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCopyJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListCopyJobs",
	}
}
