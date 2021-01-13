// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the file shares for a specific file gateway, or the list of file
// shares that belong to the calling user account. This operation is only supported
// for file gateways.
func (c *Client) ListFileShares(ctx context.Context, params *ListFileSharesInput, optFns ...func(*Options)) (*ListFileSharesOutput, error) {
	if params == nil {
		params = &ListFileSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFileShares", params, optFns, addOperationListFileSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFileSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListFileShareInput
type ListFileSharesInput struct {

	// The Amazon Resource Name (ARN) of the gateway whose file shares you want to
	// list. If this field is not present, all file shares under your account are
	// listed.
	GatewayARN *string

	// The maximum number of file shares to return in the response. The value must be
	// an integer with a value greater than zero. Optional.
	Limit *int32

	// Opaque pagination token returned from a previous ListFileShares operation. If
	// present, Marker specifies where to continue the list from after a previous call
	// to ListFileShares. Optional.
	Marker *string
}

// ListFileShareOutput
type ListFileSharesOutput struct {

	// An array of information about the file gateway's file shares.
	FileShareInfoList []types.FileShareInfo

	// If the request includes Marker, the response returns that value in this field.
	Marker *string

	// If a value is present, there are more file shares to return. In a subsequent
	// request, use NextMarker as the value for Marker to retrieve the next set of file
	// shares.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFileSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFileShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFileShares{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFileShares(options.Region), middleware.Before); err != nil {
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

// ListFileSharesAPIClient is a client that implements the ListFileShares
// operation.
type ListFileSharesAPIClient interface {
	ListFileShares(context.Context, *ListFileSharesInput, ...func(*Options)) (*ListFileSharesOutput, error)
}

var _ ListFileSharesAPIClient = (*Client)(nil)

// ListFileSharesPaginatorOptions is the paginator options for ListFileShares
type ListFileSharesPaginatorOptions struct {
	// The maximum number of file shares to return in the response. The value must be
	// an integer with a value greater than zero. Optional.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFileSharesPaginator is a paginator for ListFileShares
type ListFileSharesPaginator struct {
	options   ListFileSharesPaginatorOptions
	client    ListFileSharesAPIClient
	params    *ListFileSharesInput
	nextToken *string
	firstPage bool
}

// NewListFileSharesPaginator returns a new ListFileSharesPaginator
func NewListFileSharesPaginator(client ListFileSharesAPIClient, params *ListFileSharesInput, optFns ...func(*ListFileSharesPaginatorOptions)) *ListFileSharesPaginator {
	options := ListFileSharesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListFileSharesInput{}
	}

	return &ListFileSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFileSharesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFileShares page.
func (p *ListFileSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFileSharesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListFileShares(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFileShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListFileShares",
	}
}
