// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of running AWS App Runner services in your AWS account.
func (c *Client) ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) {
	if params == nil {
		params = &ListServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServices", params, optFns, c.addOperationListServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServicesInput struct {

	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults, the request
	// retrieves all available results in a single response.
	MaxResults *int32

	// A token from a previous result page. Used for a paginated request. The request
	// retrieves the next result page. All other parameter values must be identical to
	// the ones specified in the initial request. If you don't specify NextToken, the
	// request retrieves the first result page.
	NextToken *string
}

type ListServicesOutput struct {

	// A list of service summary information records. In a paginated request, the
	// request returns up to MaxResults records for each call.
	//
	// This member is required.
	ServiceSummaryList []types.ServiceSummary

	// The token that you can pass in a subsequent request to get the next result page.
	// It's returned in a paginated request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServices(options.Region), middleware.Before); err != nil {
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

// ListServicesAPIClient is a client that implements the ListServices operation.
type ListServicesAPIClient interface {
	ListServices(context.Context, *ListServicesInput, ...func(*Options)) (*ListServicesOutput, error)
}

var _ ListServicesAPIClient = (*Client)(nil)

// ListServicesPaginatorOptions is the paginator options for ListServices
type ListServicesPaginatorOptions struct {
	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults, the request
	// retrieves all available results in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServicesPaginator is a paginator for ListServices
type ListServicesPaginator struct {
	options   ListServicesPaginatorOptions
	client    ListServicesAPIClient
	params    *ListServicesInput
	nextToken *string
	firstPage bool
}

// NewListServicesPaginator returns a new ListServicesPaginator
func NewListServicesPaginator(client ListServicesAPIClient, params *ListServicesInput, optFns ...func(*ListServicesPaginatorOptions)) *ListServicesPaginator {
	if params == nil {
		params = &ListServicesInput{}
	}

	options := ListServicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServicesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListServices page.
func (p *ListServicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServicesOutput, error) {
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

	result, err := p.client.ListServices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "ListServices",
	}
}
