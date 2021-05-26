// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists exports as specified by ID. All continuous exports associated with your
// user account can be listed if you call DescribeContinuousExports as is without
// passing any parameters.
func (c *Client) DescribeContinuousExports(ctx context.Context, params *DescribeContinuousExportsInput, optFns ...func(*Options)) (*DescribeContinuousExportsOutput, error) {
	if params == nil {
		params = &DescribeContinuousExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeContinuousExports", params, optFns, c.addOperationDescribeContinuousExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeContinuousExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContinuousExportsInput struct {

	// The unique IDs assigned to the exports.
	ExportIds []string

	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	MaxResults *int32

	// The token from the previous call to DescribeExportTasks.
	NextToken *string
}

type DescribeContinuousExportsOutput struct {

	// A list of continuous export descriptions.
	Descriptions []types.ContinuousExportDescription

	// The token from the previous call to DescribeExportTasks.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeContinuousExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeContinuousExports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeContinuousExports{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeContinuousExports(options.Region), middleware.Before); err != nil {
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

// DescribeContinuousExportsAPIClient is a client that implements the
// DescribeContinuousExports operation.
type DescribeContinuousExportsAPIClient interface {
	DescribeContinuousExports(context.Context, *DescribeContinuousExportsInput, ...func(*Options)) (*DescribeContinuousExportsOutput, error)
}

var _ DescribeContinuousExportsAPIClient = (*Client)(nil)

// DescribeContinuousExportsPaginatorOptions is the paginator options for
// DescribeContinuousExports
type DescribeContinuousExportsPaginatorOptions struct {
	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeContinuousExportsPaginator is a paginator for DescribeContinuousExports
type DescribeContinuousExportsPaginator struct {
	options   DescribeContinuousExportsPaginatorOptions
	client    DescribeContinuousExportsAPIClient
	params    *DescribeContinuousExportsInput
	nextToken *string
	firstPage bool
}

// NewDescribeContinuousExportsPaginator returns a new
// DescribeContinuousExportsPaginator
func NewDescribeContinuousExportsPaginator(client DescribeContinuousExportsAPIClient, params *DescribeContinuousExportsInput, optFns ...func(*DescribeContinuousExportsPaginatorOptions)) *DescribeContinuousExportsPaginator {
	if params == nil {
		params = &DescribeContinuousExportsInput{}
	}

	options := DescribeContinuousExportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeContinuousExportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeContinuousExportsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeContinuousExports page.
func (p *DescribeContinuousExportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeContinuousExportsOutput, error) {
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

	result, err := p.client.DescribeContinuousExports(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeContinuousExports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeContinuousExports",
	}
}
