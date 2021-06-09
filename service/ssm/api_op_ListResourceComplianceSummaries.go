// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a resource-level summary count. The summary includes information about
// compliant and non-compliant statuses and detailed compliance-item severity
// counts, according to the filter criteria you specify.
func (c *Client) ListResourceComplianceSummaries(ctx context.Context, params *ListResourceComplianceSummariesInput, optFns ...func(*Options)) (*ListResourceComplianceSummariesOutput, error) {
	if params == nil {
		params = &ListResourceComplianceSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceComplianceSummaries", params, optFns, c.addOperationListResourceComplianceSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceComplianceSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceComplianceSummariesInput struct {

	// One or more filters. Use a filter to return a more specific list of results.
	Filters []types.ComplianceStringFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type ListResourceComplianceSummariesOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A summary count for specified or targeted managed instances. Summary count
	// includes information about compliant and non-compliant State Manager
	// associations, patch status, or custom items according to the filter criteria
	// that you specify.
	ResourceComplianceSummaryItems []types.ResourceComplianceSummaryItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListResourceComplianceSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceComplianceSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceComplianceSummaries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceComplianceSummaries(options.Region), middleware.Before); err != nil {
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

// ListResourceComplianceSummariesAPIClient is a client that implements the
// ListResourceComplianceSummaries operation.
type ListResourceComplianceSummariesAPIClient interface {
	ListResourceComplianceSummaries(context.Context, *ListResourceComplianceSummariesInput, ...func(*Options)) (*ListResourceComplianceSummariesOutput, error)
}

var _ ListResourceComplianceSummariesAPIClient = (*Client)(nil)

// ListResourceComplianceSummariesPaginatorOptions is the paginator options for
// ListResourceComplianceSummaries
type ListResourceComplianceSummariesPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceComplianceSummariesPaginator is a paginator for
// ListResourceComplianceSummaries
type ListResourceComplianceSummariesPaginator struct {
	options   ListResourceComplianceSummariesPaginatorOptions
	client    ListResourceComplianceSummariesAPIClient
	params    *ListResourceComplianceSummariesInput
	nextToken *string
	firstPage bool
}

// NewListResourceComplianceSummariesPaginator returns a new
// ListResourceComplianceSummariesPaginator
func NewListResourceComplianceSummariesPaginator(client ListResourceComplianceSummariesAPIClient, params *ListResourceComplianceSummariesInput, optFns ...func(*ListResourceComplianceSummariesPaginatorOptions)) *ListResourceComplianceSummariesPaginator {
	if params == nil {
		params = &ListResourceComplianceSummariesInput{}
	}

	options := ListResourceComplianceSummariesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceComplianceSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceComplianceSummariesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListResourceComplianceSummaries page.
func (p *ListResourceComplianceSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceComplianceSummariesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListResourceComplianceSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourceComplianceSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListResourceComplianceSummaries",
	}
}
