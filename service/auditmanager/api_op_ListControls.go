// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of controls from AWS Audit Manager.
func (c *Client) ListControls(ctx context.Context, params *ListControlsInput, optFns ...func(*Options)) (*ListControlsOutput, error) {
	if params == nil {
		params = &ListControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListControls", params, optFns, c.addOperationListControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListControlsInput struct {

	// The type of control, such as standard or custom.
	//
	// This member is required.
	ControlType types.ControlType

	// Represents the maximum number of results per page, or per API request call.
	MaxResults *int32

	// The pagination token used to fetch the next set of results.
	NextToken *string
}

type ListControlsOutput struct {

	// The list of control metadata objects returned by the ListControls API.
	ControlMetadataList []types.ControlMetadata

	// The pagination token used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListControls{}, middleware.After)
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
	if err = addOpListControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListControls(options.Region), middleware.Before); err != nil {
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

// ListControlsAPIClient is a client that implements the ListControls operation.
type ListControlsAPIClient interface {
	ListControls(context.Context, *ListControlsInput, ...func(*Options)) (*ListControlsOutput, error)
}

var _ ListControlsAPIClient = (*Client)(nil)

// ListControlsPaginatorOptions is the paginator options for ListControls
type ListControlsPaginatorOptions struct {
	// Represents the maximum number of results per page, or per API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListControlsPaginator is a paginator for ListControls
type ListControlsPaginator struct {
	options   ListControlsPaginatorOptions
	client    ListControlsAPIClient
	params    *ListControlsInput
	nextToken *string
	firstPage bool
}

// NewListControlsPaginator returns a new ListControlsPaginator
func NewListControlsPaginator(client ListControlsAPIClient, params *ListControlsInput, optFns ...func(*ListControlsPaginatorOptions)) *ListControlsPaginator {
	if params == nil {
		params = &ListControlsInput{}
	}

	options := ListControlsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListControlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListControlsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListControls page.
func (p *ListControlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListControlsOutput, error) {
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

	result, err := p.client.ListControls(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "ListControls",
	}
}
