// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists endpoints.
func (c *Client) ListEndpoints(ctx context.Context, params *ListEndpointsInput, optFns ...func(*Options)) (*ListEndpointsOutput, error) {
	if params == nil {
		params = &ListEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpoints", params, optFns, addOperationListEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointsInput struct {

	// A filter that returns only endpoints with a creation time greater than or equal
	// to the specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only endpoints that were created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only endpoints that were modified after the specified
	// timestamp.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only endpoints that were modified before the specified
	// timestamp.
	LastModifiedTimeBefore *time.Time

	// The maximum number of endpoints to return in the response.
	MaxResults *int32

	// A string in endpoint names. This filter returns only endpoints whose name
	// contains the specified string.
	NameContains *string

	// If the result of a ListEndpoints request was truncated, the response includes a
	// NextToken. To retrieve the next set of endpoints, use the token in the next
	// request.
	NextToken *string

	// Sorts the list of results. The default is CreationTime.
	SortBy types.EndpointSortKey

	// The sort order for results. The default is Descending.
	SortOrder types.OrderKey

	// A filter that returns only endpoints with the specified status.
	StatusEquals types.EndpointStatus
}

type ListEndpointsOutput struct {

	// An array or endpoint objects.
	//
	// This member is required.
	Endpoints []types.EndpointSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of training jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpoints(options.Region), middleware.Before); err != nil {
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

// ListEndpointsAPIClient is a client that implements the ListEndpoints operation.
type ListEndpointsAPIClient interface {
	ListEndpoints(context.Context, *ListEndpointsInput, ...func(*Options)) (*ListEndpointsOutput, error)
}

var _ ListEndpointsAPIClient = (*Client)(nil)

// ListEndpointsPaginatorOptions is the paginator options for ListEndpoints
type ListEndpointsPaginatorOptions struct {
	// The maximum number of endpoints to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEndpointsPaginator is a paginator for ListEndpoints
type ListEndpointsPaginator struct {
	options   ListEndpointsPaginatorOptions
	client    ListEndpointsAPIClient
	params    *ListEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListEndpointsPaginator returns a new ListEndpointsPaginator
func NewListEndpointsPaginator(client ListEndpointsAPIClient, params *ListEndpointsInput, optFns ...func(*ListEndpointsPaginatorOptions)) *ListEndpointsPaginator {
	options := ListEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListEndpointsInput{}
	}

	return &ListEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListEndpoints page.
func (p *ListEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEndpointsOutput, error) {
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

	result, err := p.client.ListEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListEndpoints",
	}
}
