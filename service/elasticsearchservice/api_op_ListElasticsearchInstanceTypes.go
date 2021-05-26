// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all Elasticsearch instance types that are supported for given
// ElasticsearchVersion
func (c *Client) ListElasticsearchInstanceTypes(ctx context.Context, params *ListElasticsearchInstanceTypesInput, optFns ...func(*Options)) (*ListElasticsearchInstanceTypesOutput, error) {
	if params == nil {
		params = &ListElasticsearchInstanceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListElasticsearchInstanceTypes", params, optFns, c.addOperationListElasticsearchInstanceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListElasticsearchInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the ListElasticsearchInstanceTypes operation.
type ListElasticsearchInstanceTypesInput struct {

	// Version of Elasticsearch for which list of supported elasticsearch instance
	// types are needed.
	//
	// This member is required.
	ElasticsearchVersion *string

	// DomainName represents the name of the Domain that we are trying to modify. This
	// should be present only if we are querying for list of available Elasticsearch
	// instance types when modifying existing domain.
	DomainName *string

	// Set this value to limit the number of results returned. Value provided must be
	// greater than 30 else it wont be honored.
	MaxResults int32

	// NextToken should be sent in case if earlier API call produced result containing
	// NextToken. It is used for pagination.
	NextToken *string
}

// Container for the parameters returned by ListElasticsearchInstanceTypes
// operation.
type ListElasticsearchInstanceTypesOutput struct {

	// List of instance types supported by Amazon Elasticsearch service for given
	// ElasticsearchVersion
	ElasticsearchInstanceTypes []types.ESPartitionInstanceType

	// In case if there are more results available NextToken would be present, make
	// further request to the same API with received NextToken to paginate remaining
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListElasticsearchInstanceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListElasticsearchInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListElasticsearchInstanceTypes{}, middleware.After)
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
	if err = addOpListElasticsearchInstanceTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListElasticsearchInstanceTypes(options.Region), middleware.Before); err != nil {
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

// ListElasticsearchInstanceTypesAPIClient is a client that implements the
// ListElasticsearchInstanceTypes operation.
type ListElasticsearchInstanceTypesAPIClient interface {
	ListElasticsearchInstanceTypes(context.Context, *ListElasticsearchInstanceTypesInput, ...func(*Options)) (*ListElasticsearchInstanceTypesOutput, error)
}

var _ ListElasticsearchInstanceTypesAPIClient = (*Client)(nil)

// ListElasticsearchInstanceTypesPaginatorOptions is the paginator options for
// ListElasticsearchInstanceTypes
type ListElasticsearchInstanceTypesPaginatorOptions struct {
	// Set this value to limit the number of results returned. Value provided must be
	// greater than 30 else it wont be honored.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListElasticsearchInstanceTypesPaginator is a paginator for
// ListElasticsearchInstanceTypes
type ListElasticsearchInstanceTypesPaginator struct {
	options   ListElasticsearchInstanceTypesPaginatorOptions
	client    ListElasticsearchInstanceTypesAPIClient
	params    *ListElasticsearchInstanceTypesInput
	nextToken *string
	firstPage bool
}

// NewListElasticsearchInstanceTypesPaginator returns a new
// ListElasticsearchInstanceTypesPaginator
func NewListElasticsearchInstanceTypesPaginator(client ListElasticsearchInstanceTypesAPIClient, params *ListElasticsearchInstanceTypesInput, optFns ...func(*ListElasticsearchInstanceTypesPaginatorOptions)) *ListElasticsearchInstanceTypesPaginator {
	if params == nil {
		params = &ListElasticsearchInstanceTypesInput{}
	}

	options := ListElasticsearchInstanceTypesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListElasticsearchInstanceTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListElasticsearchInstanceTypesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListElasticsearchInstanceTypes page.
func (p *ListElasticsearchInstanceTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListElasticsearchInstanceTypesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListElasticsearchInstanceTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListElasticsearchInstanceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListElasticsearchInstanceTypes",
	}
}
