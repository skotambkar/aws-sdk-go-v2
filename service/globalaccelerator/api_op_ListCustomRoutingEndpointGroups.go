// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the endpoint groups that are associated with a listener for a custom
// routing accelerator.
func (c *Client) ListCustomRoutingEndpointGroups(ctx context.Context, params *ListCustomRoutingEndpointGroupsInput, optFns ...func(*Options)) (*ListCustomRoutingEndpointGroupsOutput, error) {
	if params == nil {
		params = &ListCustomRoutingEndpointGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomRoutingEndpointGroups", params, optFns, c.addOperationListCustomRoutingEndpointGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomRoutingEndpointGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomRoutingEndpointGroupsInput struct {

	// The Amazon Resource Name (ARN) of the listener to list endpoint groups for.
	//
	// This member is required.
	ListenerArn *string

	// The number of endpoint group objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
}

type ListCustomRoutingEndpointGroupsOutput struct {

	// The list of the endpoint groups associated with a listener for a custom routing
	// accelerator.
	EndpointGroups []types.CustomRoutingEndpointGroup

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListCustomRoutingEndpointGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCustomRoutingEndpointGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCustomRoutingEndpointGroups{}, middleware.After)
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
	if err = addOpListCustomRoutingEndpointGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomRoutingEndpointGroups(options.Region), middleware.Before); err != nil {
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

// ListCustomRoutingEndpointGroupsAPIClient is a client that implements the
// ListCustomRoutingEndpointGroups operation.
type ListCustomRoutingEndpointGroupsAPIClient interface {
	ListCustomRoutingEndpointGroups(context.Context, *ListCustomRoutingEndpointGroupsInput, ...func(*Options)) (*ListCustomRoutingEndpointGroupsOutput, error)
}

var _ ListCustomRoutingEndpointGroupsAPIClient = (*Client)(nil)

// ListCustomRoutingEndpointGroupsPaginatorOptions is the paginator options for
// ListCustomRoutingEndpointGroups
type ListCustomRoutingEndpointGroupsPaginatorOptions struct {
	// The number of endpoint group objects that you want to return with this call. The
	// default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomRoutingEndpointGroupsPaginator is a paginator for
// ListCustomRoutingEndpointGroups
type ListCustomRoutingEndpointGroupsPaginator struct {
	options   ListCustomRoutingEndpointGroupsPaginatorOptions
	client    ListCustomRoutingEndpointGroupsAPIClient
	params    *ListCustomRoutingEndpointGroupsInput
	nextToken *string
	firstPage bool
}

// NewListCustomRoutingEndpointGroupsPaginator returns a new
// ListCustomRoutingEndpointGroupsPaginator
func NewListCustomRoutingEndpointGroupsPaginator(client ListCustomRoutingEndpointGroupsAPIClient, params *ListCustomRoutingEndpointGroupsInput, optFns ...func(*ListCustomRoutingEndpointGroupsPaginatorOptions)) *ListCustomRoutingEndpointGroupsPaginator {
	if params == nil {
		params = &ListCustomRoutingEndpointGroupsInput{}
	}

	options := ListCustomRoutingEndpointGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomRoutingEndpointGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomRoutingEndpointGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCustomRoutingEndpointGroups page.
func (p *ListCustomRoutingEndpointGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomRoutingEndpointGroupsOutput, error) {
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

	result, err := p.client.ListCustomRoutingEndpointGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomRoutingEndpointGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListCustomRoutingEndpointGroups",
	}
}
