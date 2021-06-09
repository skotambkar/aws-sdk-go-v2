// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the VpcLinks collection under the caller's account in a selected region.
func (c *Client) GetVpcLinks(ctx context.Context, params *GetVpcLinksInput, optFns ...func(*Options)) (*GetVpcLinksOutput, error) {
	if params == nil {
		params = &GetVpcLinksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVpcLinks", params, optFns, c.addOperationGetVpcLinksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVpcLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the VpcLinks collection under the caller's account in a selected region.
type GetVpcLinksInput struct {

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string
}

// The collection of VPC links under the caller's account in a region. Getting
// Started with Private Integrations
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-with-private-integration.html),
// Set up Private Integrations
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-private-integration.html)
type GetVpcLinksOutput struct {

	// The current page of elements from this collection.
	Items []types.VpcLink

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetVpcLinksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVpcLinks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVpcLinks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVpcLinks(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetVpcLinksAPIClient is a client that implements the GetVpcLinks operation.
type GetVpcLinksAPIClient interface {
	GetVpcLinks(context.Context, *GetVpcLinksInput, ...func(*Options)) (*GetVpcLinksOutput, error)
}

var _ GetVpcLinksAPIClient = (*Client)(nil)

// GetVpcLinksPaginatorOptions is the paginator options for GetVpcLinks
type GetVpcLinksPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetVpcLinksPaginator is a paginator for GetVpcLinks
type GetVpcLinksPaginator struct {
	options   GetVpcLinksPaginatorOptions
	client    GetVpcLinksAPIClient
	params    *GetVpcLinksInput
	nextToken *string
	firstPage bool
}

// NewGetVpcLinksPaginator returns a new GetVpcLinksPaginator
func NewGetVpcLinksPaginator(client GetVpcLinksAPIClient, params *GetVpcLinksInput, optFns ...func(*GetVpcLinksPaginatorOptions)) *GetVpcLinksPaginator {
	if params == nil {
		params = &GetVpcLinksInput{}
	}

	options := GetVpcLinksPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetVpcLinksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetVpcLinksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetVpcLinks page.
func (p *GetVpcLinksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetVpcLinksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetVpcLinks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetVpcLinks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetVpcLinks",
	}
}
