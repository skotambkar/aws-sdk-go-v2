// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about events across your organization in AWS Organizations.
// You can use thefilters parameter to specify the events that you want to return.
// Events are returned in a summary form and don't include the affected accounts,
// detailed description, any additional metadata that depends on the event type, or
// any affected resources. To retrieve that information, use the following
// operations:
//
// * DescribeAffectedAccountsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedAccountsForOrganization.html)
//
// *
// DescribeEventDetailsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
//
// *
// DescribeAffectedEntitiesForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html)
//
// If
// you don't specify a filter, the DescribeEventsForOrganizations returns all
// events across your organization. Results are sorted by lastModifiedTime,
// starting with the most recent event. For more information about the different
// types of AWS Health events, see Event
// (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html). Before
// you can call this operation, you must first enable AWS Health to work with AWS
// Organizations. To do this, call the EnableHealthServiceAccessForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html)
// operation from your organization's master AWS account. This API operation uses
// pagination. Specify the nextToken parameter in the next request to return more
// results.
func (c *Client) DescribeEventsForOrganization(ctx context.Context, params *DescribeEventsForOrganizationInput, optFns ...func(*Options)) (*DescribeEventsForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeEventsForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventsForOrganization", params, optFns, addOperationDescribeEventsForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventsForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventsForOrganizationInput struct {

	// Values to narrow the results returned.
	Filter *types.OrganizationEventFilter

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string
}

type DescribeEventsForOrganizationOutput struct {

	// The events that match the specified filter criteria.
	Events []types.OrganizationEvent

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventsForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventsForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventsForOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventsForOrganization(options.Region), middleware.Before); err != nil {
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

// DescribeEventsForOrganizationAPIClient is a client that implements the
// DescribeEventsForOrganization operation.
type DescribeEventsForOrganizationAPIClient interface {
	DescribeEventsForOrganization(context.Context, *DescribeEventsForOrganizationInput, ...func(*Options)) (*DescribeEventsForOrganizationOutput, error)
}

var _ DescribeEventsForOrganizationAPIClient = (*Client)(nil)

// DescribeEventsForOrganizationPaginatorOptions is the paginator options for
// DescribeEventsForOrganization
type DescribeEventsForOrganizationPaginatorOptions struct {
	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEventsForOrganizationPaginator is a paginator for
// DescribeEventsForOrganization
type DescribeEventsForOrganizationPaginator struct {
	options   DescribeEventsForOrganizationPaginatorOptions
	client    DescribeEventsForOrganizationAPIClient
	params    *DescribeEventsForOrganizationInput
	nextToken *string
	firstPage bool
}

// NewDescribeEventsForOrganizationPaginator returns a new
// DescribeEventsForOrganizationPaginator
func NewDescribeEventsForOrganizationPaginator(client DescribeEventsForOrganizationAPIClient, params *DescribeEventsForOrganizationInput, optFns ...func(*DescribeEventsForOrganizationPaginatorOptions)) *DescribeEventsForOrganizationPaginator {
	options := DescribeEventsForOrganizationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeEventsForOrganizationInput{}
	}

	return &DescribeEventsForOrganizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEventsForOrganizationPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeEventsForOrganization page.
func (p *DescribeEventsForOrganizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEventsForOrganizationOutput, error) {
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

	result, err := p.client.DescribeEventsForOrganization(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEventsForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeEventsForOrganization",
	}
}
