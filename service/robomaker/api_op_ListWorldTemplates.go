// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists world templates.
func (c *Client) ListWorldTemplates(ctx context.Context, params *ListWorldTemplatesInput, optFns ...func(*Options)) (*ListWorldTemplatesOutput, error) {
	if params == nil {
		params = &ListWorldTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorldTemplates", params, optFns, c.addOperationListWorldTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorldTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorldTemplatesInput struct {

	// When this parameter is used, ListWorldTemplates only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListWorldTemplates request
	// with the returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, then ListWorldTemplates returns up to 100 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldTemplates again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
}

type ListWorldTemplatesOutput struct {

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldTemplates again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	// Summary information for templates.
	TemplateSummaries []types.TemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListWorldTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorldTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorldTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorldTemplates(options.Region), middleware.Before); err != nil {
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

// ListWorldTemplatesAPIClient is a client that implements the ListWorldTemplates
// operation.
type ListWorldTemplatesAPIClient interface {
	ListWorldTemplates(context.Context, *ListWorldTemplatesInput, ...func(*Options)) (*ListWorldTemplatesOutput, error)
}

var _ ListWorldTemplatesAPIClient = (*Client)(nil)

// ListWorldTemplatesPaginatorOptions is the paginator options for
// ListWorldTemplates
type ListWorldTemplatesPaginatorOptions struct {
	// When this parameter is used, ListWorldTemplates only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListWorldTemplates request
	// with the returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, then ListWorldTemplates returns up to 100 results and a
	// nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorldTemplatesPaginator is a paginator for ListWorldTemplates
type ListWorldTemplatesPaginator struct {
	options   ListWorldTemplatesPaginatorOptions
	client    ListWorldTemplatesAPIClient
	params    *ListWorldTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListWorldTemplatesPaginator returns a new ListWorldTemplatesPaginator
func NewListWorldTemplatesPaginator(client ListWorldTemplatesAPIClient, params *ListWorldTemplatesInput, optFns ...func(*ListWorldTemplatesPaginatorOptions)) *ListWorldTemplatesPaginator {
	if params == nil {
		params = &ListWorldTemplatesInput{}
	}

	options := ListWorldTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorldTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorldTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListWorldTemplates page.
func (p *ListWorldTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorldTemplatesOutput, error) {
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

	result, err := p.client.ListWorldTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorldTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListWorldTemplates",
	}
}
