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

// Returns a list of the actions taken in a TestGridSession.
func (c *Client) ListTestGridSessionActions(ctx context.Context, params *ListTestGridSessionActionsInput, optFns ...func(*Options)) (*ListTestGridSessionActionsOutput, error) {
	if params == nil {
		params = &ListTestGridSessionActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestGridSessionActions", params, optFns, addOperationListTestGridSessionActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestGridSessionActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestGridSessionActionsInput struct {

	// The ARN of the session to retrieve.
	//
	// This member is required.
	SessionArn *string

	// The maximum number of sessions to return per response.
	MaxResult *int32

	// Pagination token.
	NextToken *string
}

type ListTestGridSessionActionsOutput struct {

	// The action taken by the session.
	Actions []types.TestGridSessionAction

	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTestGridSessionActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTestGridSessionActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTestGridSessionActions{}, middleware.After)
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
	if err = addOpListTestGridSessionActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestGridSessionActions(options.Region), middleware.Before); err != nil {
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

// ListTestGridSessionActionsAPIClient is a client that implements the
// ListTestGridSessionActions operation.
type ListTestGridSessionActionsAPIClient interface {
	ListTestGridSessionActions(context.Context, *ListTestGridSessionActionsInput, ...func(*Options)) (*ListTestGridSessionActionsOutput, error)
}

var _ ListTestGridSessionActionsAPIClient = (*Client)(nil)

// ListTestGridSessionActionsPaginatorOptions is the paginator options for
// ListTestGridSessionActions
type ListTestGridSessionActionsPaginatorOptions struct {
	// The maximum number of sessions to return per response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestGridSessionActionsPaginator is a paginator for
// ListTestGridSessionActions
type ListTestGridSessionActionsPaginator struct {
	options   ListTestGridSessionActionsPaginatorOptions
	client    ListTestGridSessionActionsAPIClient
	params    *ListTestGridSessionActionsInput
	nextToken *string
	firstPage bool
}

// NewListTestGridSessionActionsPaginator returns a new
// ListTestGridSessionActionsPaginator
func NewListTestGridSessionActionsPaginator(client ListTestGridSessionActionsAPIClient, params *ListTestGridSessionActionsInput, optFns ...func(*ListTestGridSessionActionsPaginatorOptions)) *ListTestGridSessionActionsPaginator {
	options := ListTestGridSessionActionsPaginatorOptions{}
	if params.MaxResult != nil {
		options.Limit = *params.MaxResult
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTestGridSessionActionsInput{}
	}

	return &ListTestGridSessionActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestGridSessionActionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTestGridSessionActions page.
func (p *ListTestGridSessionActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestGridSessionActionsOutput, error) {
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

	result, err := p.client.ListTestGridSessionActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTestGridSessionActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListTestGridSessionActions",
	}
}
