// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the quota increase requests for the specified service.
func (c *Client) ListRequestedServiceQuotaChangeHistory(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryOutput, error) {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRequestedServiceQuotaChangeHistory", params, optFns, addOperationListRequestedServiceQuotaChangeHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRequestedServiceQuotaChangeHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRequestedServiceQuotaChangeHistoryInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The service identifier.
	ServiceCode *string

	// The status of the quota increase request.
	Status types.RequestStatus
}

type ListRequestedServiceQuotaChangeHistoryOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the quota increase requests.
	RequestedQuotas []types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRequestedServiceQuotaChangeHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRequestedServiceQuotaChangeHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRequestedServiceQuotaChangeHistory{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistory(options.Region), middleware.Before); err != nil {
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

// ListRequestedServiceQuotaChangeHistoryAPIClient is a client that implements the
// ListRequestedServiceQuotaChangeHistory operation.
type ListRequestedServiceQuotaChangeHistoryAPIClient interface {
	ListRequestedServiceQuotaChangeHistory(context.Context, *ListRequestedServiceQuotaChangeHistoryInput, ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryOutput, error)
}

var _ ListRequestedServiceQuotaChangeHistoryAPIClient = (*Client)(nil)

// ListRequestedServiceQuotaChangeHistoryPaginatorOptions is the paginator options
// for ListRequestedServiceQuotaChangeHistory
type ListRequestedServiceQuotaChangeHistoryPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRequestedServiceQuotaChangeHistoryPaginator is a paginator for
// ListRequestedServiceQuotaChangeHistory
type ListRequestedServiceQuotaChangeHistoryPaginator struct {
	options   ListRequestedServiceQuotaChangeHistoryPaginatorOptions
	client    ListRequestedServiceQuotaChangeHistoryAPIClient
	params    *ListRequestedServiceQuotaChangeHistoryInput
	nextToken *string
	firstPage bool
}

// NewListRequestedServiceQuotaChangeHistoryPaginator returns a new
// ListRequestedServiceQuotaChangeHistoryPaginator
func NewListRequestedServiceQuotaChangeHistoryPaginator(client ListRequestedServiceQuotaChangeHistoryAPIClient, params *ListRequestedServiceQuotaChangeHistoryInput, optFns ...func(*ListRequestedServiceQuotaChangeHistoryPaginatorOptions)) *ListRequestedServiceQuotaChangeHistoryPaginator {
	options := ListRequestedServiceQuotaChangeHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryInput{}
	}

	return &ListRequestedServiceQuotaChangeHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRequestedServiceQuotaChangeHistoryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListRequestedServiceQuotaChangeHistory page.
func (p *ListRequestedServiceQuotaChangeHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryOutput, error) {
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

	result, err := p.client.ListRequestedServiceQuotaChangeHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListRequestedServiceQuotaChangeHistory",
	}
}
