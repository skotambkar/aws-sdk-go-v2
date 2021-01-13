// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListBonusPayments operation retrieves the amounts of bonuses you have paid
// to Workers for a given HIT or assignment.
func (c *Client) ListBonusPayments(ctx context.Context, params *ListBonusPaymentsInput, optFns ...func(*Options)) (*ListBonusPaymentsOutput, error) {
	if params == nil {
		params = &ListBonusPaymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBonusPayments", params, optFns, addOperationListBonusPaymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBonusPaymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBonusPaymentsInput struct {

	// The ID of the assignment associated with the bonus payments to retrieve. If
	// specified, only bonus payments for the given assignment are returned. Either the
	// HITId parameter or the AssignmentId parameter must be specified
	AssignmentId *string

	// The ID of the HIT associated with the bonus payments to retrieve. If not
	// specified, all bonus payments for all assignments for the given HIT are
	// returned. Either the HITId parameter or the AssignmentId parameter must be
	// specified
	HITId *string

	MaxResults *int32

	// Pagination token
	NextToken *string
}

type ListBonusPaymentsOutput struct {

	// A successful request to the ListBonusPayments operation returns a list of
	// BonusPayment objects.
	BonusPayments []types.BonusPayment

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of bonus payments on this page in the filtered results list,
	// equivalent to the number of bonus payments being returned by this call.
	NumResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBonusPaymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBonusPayments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBonusPayments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBonusPayments(options.Region), middleware.Before); err != nil {
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

// ListBonusPaymentsAPIClient is a client that implements the ListBonusPayments
// operation.
type ListBonusPaymentsAPIClient interface {
	ListBonusPayments(context.Context, *ListBonusPaymentsInput, ...func(*Options)) (*ListBonusPaymentsOutput, error)
}

var _ ListBonusPaymentsAPIClient = (*Client)(nil)

// ListBonusPaymentsPaginatorOptions is the paginator options for ListBonusPayments
type ListBonusPaymentsPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBonusPaymentsPaginator is a paginator for ListBonusPayments
type ListBonusPaymentsPaginator struct {
	options   ListBonusPaymentsPaginatorOptions
	client    ListBonusPaymentsAPIClient
	params    *ListBonusPaymentsInput
	nextToken *string
	firstPage bool
}

// NewListBonusPaymentsPaginator returns a new ListBonusPaymentsPaginator
func NewListBonusPaymentsPaginator(client ListBonusPaymentsAPIClient, params *ListBonusPaymentsInput, optFns ...func(*ListBonusPaymentsPaginatorOptions)) *ListBonusPaymentsPaginator {
	options := ListBonusPaymentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListBonusPaymentsInput{}
	}

	return &ListBonusPaymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBonusPaymentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListBonusPayments page.
func (p *ListBonusPaymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBonusPaymentsOutput, error) {
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

	result, err := p.client.ListBonusPayments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBonusPayments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListBonusPayments",
	}
}
