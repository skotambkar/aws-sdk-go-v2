// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of all Selenium testing projects in your account.
func (c *Client) ListTestGridProjects(ctx context.Context, params *ListTestGridProjectsInput, optFns ...func(*Options)) (*ListTestGridProjectsOutput, error) {
	if params == nil {
		params = &ListTestGridProjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestGridProjects", params, optFns, addOperationListTestGridProjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestGridProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestGridProjectsInput struct {

	// Return no more than this number of results.
	MaxResult *int32

	// From a response, used to continue a paginated listing.
	NextToken *string
}

type ListTestGridProjectsOutput struct {

	// Used for pagination. Pass into ListTestGridProjects to get more results in a
	// paginated request.
	NextToken *string

	// The list of TestGridProjects, based on a ListTestGridProjectsRequest.
	TestGridProjects []types.TestGridProject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTestGridProjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTestGridProjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTestGridProjects{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestGridProjects(options.Region), middleware.Before); err != nil {
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

// ListTestGridProjectsAPIClient is a client that implements the
// ListTestGridProjects operation.
type ListTestGridProjectsAPIClient interface {
	ListTestGridProjects(context.Context, *ListTestGridProjectsInput, ...func(*Options)) (*ListTestGridProjectsOutput, error)
}

var _ ListTestGridProjectsAPIClient = (*Client)(nil)

// ListTestGridProjectsPaginatorOptions is the paginator options for
// ListTestGridProjects
type ListTestGridProjectsPaginatorOptions struct {
	// Return no more than this number of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestGridProjectsPaginator is a paginator for ListTestGridProjects
type ListTestGridProjectsPaginator struct {
	options   ListTestGridProjectsPaginatorOptions
	client    ListTestGridProjectsAPIClient
	params    *ListTestGridProjectsInput
	nextToken *string
	firstPage bool
}

// NewListTestGridProjectsPaginator returns a new ListTestGridProjectsPaginator
func NewListTestGridProjectsPaginator(client ListTestGridProjectsAPIClient, params *ListTestGridProjectsInput, optFns ...func(*ListTestGridProjectsPaginatorOptions)) *ListTestGridProjectsPaginator {
	options := ListTestGridProjectsPaginatorOptions{}
	if params.MaxResult != nil {
		options.Limit = *params.MaxResult
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTestGridProjectsInput{}
	}

	return &ListTestGridProjectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestGridProjectsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTestGridProjects page.
func (p *ListTestGridProjectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestGridProjectsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResult = limit

	result, err := p.client.ListTestGridProjects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTestGridProjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListTestGridProjects",
	}
}
