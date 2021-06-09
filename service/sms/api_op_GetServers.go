// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the servers in your server catalog. Before you can describe your
// servers, you must import them using ImportServerCatalog.
func (c *Client) GetServers(ctx context.Context, params *GetServersInput, optFns ...func(*Options)) (*GetServersOutput, error) {
	if params == nil {
		params = &GetServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServers", params, optFns, c.addOperationGetServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServersInput struct {

	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The server addresses.
	VmServerAddressList []types.VmServerAddress
}

type GetServersOutput struct {

	// The time when the server was last modified.
	LastModifiedOn *time.Time

	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The status of the server catalog.
	ServerCatalogStatus types.ServerCatalogStatus

	// Information about the servers.
	ServerList []types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServers(options.Region), middleware.Before); err != nil {
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

// GetServersAPIClient is a client that implements the GetServers operation.
type GetServersAPIClient interface {
	GetServers(context.Context, *GetServersInput, ...func(*Options)) (*GetServersOutput, error)
}

var _ GetServersAPIClient = (*Client)(nil)

// GetServersPaginatorOptions is the paginator options for GetServers
type GetServersPaginatorOptions struct {
	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetServersPaginator is a paginator for GetServers
type GetServersPaginator struct {
	options   GetServersPaginatorOptions
	client    GetServersAPIClient
	params    *GetServersInput
	nextToken *string
	firstPage bool
}

// NewGetServersPaginator returns a new GetServersPaginator
func NewGetServersPaginator(client GetServersAPIClient, params *GetServersInput, optFns ...func(*GetServersPaginatorOptions)) *GetServersPaginator {
	if params == nil {
		params = &GetServersInput{}
	}

	options := GetServersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetServersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetServers page.
func (p *GetServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetServersOutput, error) {
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

	result, err := p.client.GetServers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetServers",
	}
}
