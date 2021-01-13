// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the available products that you can subscribe to and
// integrate with Security Hub in order to consolidate findings.
func (c *Client) DescribeProducts(ctx context.Context, params *DescribeProductsInput, optFns ...func(*Options)) (*DescribeProductsOutput, error) {
	if params == nil {
		params = &DescribeProductsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProducts", params, optFns, addOperationDescribeProductsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProductsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProductsInput struct {

	// The maximum number of results to return.
	MaxResults int32

	// The token that is required for pagination. On your first call to the
	// DescribeProducts operation, set the value of this parameter to NULL. For
	// subsequent calls to the operation, to continue listing data, set the value of
	// this parameter to the value returned from the previous response.
	NextToken *string
}

type DescribeProductsOutput struct {

	// A list of products, including details for each product.
	//
	// This member is required.
	Products []types.Product

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeProductsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProducts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProducts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProducts(options.Region), middleware.Before); err != nil {
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

// DescribeProductsAPIClient is a client that implements the DescribeProducts
// operation.
type DescribeProductsAPIClient interface {
	DescribeProducts(context.Context, *DescribeProductsInput, ...func(*Options)) (*DescribeProductsOutput, error)
}

var _ DescribeProductsAPIClient = (*Client)(nil)

// DescribeProductsPaginatorOptions is the paginator options for DescribeProducts
type DescribeProductsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeProductsPaginator is a paginator for DescribeProducts
type DescribeProductsPaginator struct {
	options   DescribeProductsPaginatorOptions
	client    DescribeProductsAPIClient
	params    *DescribeProductsInput
	nextToken *string
	firstPage bool
}

// NewDescribeProductsPaginator returns a new DescribeProductsPaginator
func NewDescribeProductsPaginator(client DescribeProductsAPIClient, params *DescribeProductsInput, optFns ...func(*DescribeProductsPaginatorOptions)) *DescribeProductsPaginator {
	options := DescribeProductsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeProductsInput{}
	}

	return &DescribeProductsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeProductsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeProducts page.
func (p *DescribeProductsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeProductsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeProducts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeProducts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DescribeProducts",
	}
}
