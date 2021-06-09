// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves detailed Job log with paging.
func (c *Client) DescribeJobLogItems(ctx context.Context, params *DescribeJobLogItemsInput, optFns ...func(*Options)) (*DescribeJobLogItemsOutput, error) {
	if params == nil {
		params = &DescribeJobLogItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobLogItems", params, optFns, c.addOperationDescribeJobLogItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobLogItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobLogItemsInput struct {

	// Request to describe Job log job ID.
	//
	// This member is required.
	JobID *string

	// Request to describe Job log item maximum results.
	MaxResults int32

	// Request to describe Job log next token.
	NextToken *string
}

type DescribeJobLogItemsOutput struct {

	// Request to describe Job log response items.
	Items []types.JobLog

	// Request to describe Job log response next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeJobLogItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobLogItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobLogItems{}, middleware.After)
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
	if err = addOpDescribeJobLogItemsValidationMiddleware(stack); err != nil {
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

// DescribeJobLogItemsAPIClient is a client that implements the DescribeJobLogItems
// operation.
type DescribeJobLogItemsAPIClient interface {
	DescribeJobLogItems(context.Context, *DescribeJobLogItemsInput, ...func(*Options)) (*DescribeJobLogItemsOutput, error)
}

var _ DescribeJobLogItemsAPIClient = (*Client)(nil)

// DescribeJobLogItemsPaginatorOptions is the paginator options for
// DescribeJobLogItems
type DescribeJobLogItemsPaginatorOptions struct {
	// Request to describe Job log item maximum results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeJobLogItemsPaginator is a paginator for DescribeJobLogItems
type DescribeJobLogItemsPaginator struct {
	options   DescribeJobLogItemsPaginatorOptions
	client    DescribeJobLogItemsAPIClient
	params    *DescribeJobLogItemsInput
	nextToken *string
	firstPage bool
}

// NewDescribeJobLogItemsPaginator returns a new DescribeJobLogItemsPaginator
func NewDescribeJobLogItemsPaginator(client DescribeJobLogItemsAPIClient, params *DescribeJobLogItemsInput, optFns ...func(*DescribeJobLogItemsPaginatorOptions)) *DescribeJobLogItemsPaginator {
	if params == nil {
		params = &DescribeJobLogItemsInput{}
	}

	options := DescribeJobLogItemsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeJobLogItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeJobLogItemsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeJobLogItems page.
func (p *DescribeJobLogItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeJobLogItemsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeJobLogItems(ctx, &params, optFns...)
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
