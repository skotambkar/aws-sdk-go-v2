// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about human loops, given the specified parameters. If a
// human loop was deleted, it will not be included.
func (c *Client) ListHumanLoops(ctx context.Context, params *ListHumanLoopsInput, optFns ...func(*Options)) (*ListHumanLoopsOutput, error) {
	if params == nil {
		params = &ListHumanLoopsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHumanLoops", params, optFns, c.addOperationListHumanLoopsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHumanLoopsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHumanLoopsInput struct {

	// The Amazon Resource Name (ARN) of a flow definition.
	//
	// This member is required.
	FlowDefinitionArn *string

	// (Optional) The timestamp of the date when you want the human loops to begin in
	// ISO 8601 format. For example, 2020-02-24.
	CreationTimeAfter *time.Time

	// (Optional) The timestamp of the date before which you want the human loops to
	// begin in ISO 8601 format. For example, 2020-02-24.
	CreationTimeBefore *time.Time

	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults, then a NextToken is returned in the
	// output. You can use this token to display the next page of results.
	MaxResults int32

	// A token to display the next page of results.
	NextToken *string

	// Optional. The order for displaying results. Valid values: Ascending and
	// Descending.
	SortOrder types.SortOrder
}

type ListHumanLoopsOutput struct {

	// An array of objects that contain information about the human loops.
	//
	// This member is required.
	HumanLoopSummaries []types.HumanLoopSummary

	// A token to display the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListHumanLoopsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHumanLoops{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHumanLoops{}, middleware.After)
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
	if err = addOpListHumanLoopsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHumanLoops(options.Region), middleware.Before); err != nil {
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

// ListHumanLoopsAPIClient is a client that implements the ListHumanLoops
// operation.
type ListHumanLoopsAPIClient interface {
	ListHumanLoops(context.Context, *ListHumanLoopsInput, ...func(*Options)) (*ListHumanLoopsOutput, error)
}

var _ ListHumanLoopsAPIClient = (*Client)(nil)

// ListHumanLoopsPaginatorOptions is the paginator options for ListHumanLoops
type ListHumanLoopsPaginatorOptions struct {
	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults, then a NextToken is returned in the
	// output. You can use this token to display the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHumanLoopsPaginator is a paginator for ListHumanLoops
type ListHumanLoopsPaginator struct {
	options   ListHumanLoopsPaginatorOptions
	client    ListHumanLoopsAPIClient
	params    *ListHumanLoopsInput
	nextToken *string
	firstPage bool
}

// NewListHumanLoopsPaginator returns a new ListHumanLoopsPaginator
func NewListHumanLoopsPaginator(client ListHumanLoopsAPIClient, params *ListHumanLoopsInput, optFns ...func(*ListHumanLoopsPaginatorOptions)) *ListHumanLoopsPaginator {
	if params == nil {
		params = &ListHumanLoopsInput{}
	}

	options := ListHumanLoopsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHumanLoopsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHumanLoopsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListHumanLoops page.
func (p *ListHumanLoopsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHumanLoopsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListHumanLoops(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHumanLoops(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListHumanLoops",
	}
}
