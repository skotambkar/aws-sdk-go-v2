// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists origin access identities.
func (c *Client) ListCloudFrontOriginAccessIdentities(ctx context.Context, params *ListCloudFrontOriginAccessIdentitiesInput, optFns ...func(*Options)) (*ListCloudFrontOriginAccessIdentitiesOutput, error) {
	if params == nil {
		params = &ListCloudFrontOriginAccessIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCloudFrontOriginAccessIdentities", params, optFns, addOperationListCloudFrontOriginAccessIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCloudFrontOriginAccessIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list origin access identities.
type ListCloudFrontOriginAccessIdentitiesInput struct {

	// Use this when paginating results to indicate where to begin in your list of
	// origin access identities. The results include identities in the list that occur
	// after the marker. To get the next page of results, set the Marker to the value
	// of the NextMarker from the current page's response (which is also the ID of the
	// last identity on that page).
	Marker *string

	// The maximum number of origin access identities you want in the response body.
	MaxItems *int32
}

// The returned result of the corresponding request.
type ListCloudFrontOriginAccessIdentitiesOutput struct {

	// The CloudFrontOriginAccessIdentityList type.
	CloudFrontOriginAccessIdentityList *types.CloudFrontOriginAccessIdentityList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCloudFrontOriginAccessIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListCloudFrontOriginAccessIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListCloudFrontOriginAccessIdentities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCloudFrontOriginAccessIdentities(options.Region), middleware.Before); err != nil {
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

// ListCloudFrontOriginAccessIdentitiesAPIClient is a client that implements the
// ListCloudFrontOriginAccessIdentities operation.
type ListCloudFrontOriginAccessIdentitiesAPIClient interface {
	ListCloudFrontOriginAccessIdentities(context.Context, *ListCloudFrontOriginAccessIdentitiesInput, ...func(*Options)) (*ListCloudFrontOriginAccessIdentitiesOutput, error)
}

var _ ListCloudFrontOriginAccessIdentitiesAPIClient = (*Client)(nil)

// ListCloudFrontOriginAccessIdentitiesPaginatorOptions is the paginator options
// for ListCloudFrontOriginAccessIdentities
type ListCloudFrontOriginAccessIdentitiesPaginatorOptions struct {
	// The maximum number of origin access identities you want in the response body.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCloudFrontOriginAccessIdentitiesPaginator is a paginator for
// ListCloudFrontOriginAccessIdentities
type ListCloudFrontOriginAccessIdentitiesPaginator struct {
	options   ListCloudFrontOriginAccessIdentitiesPaginatorOptions
	client    ListCloudFrontOriginAccessIdentitiesAPIClient
	params    *ListCloudFrontOriginAccessIdentitiesInput
	nextToken *string
	firstPage bool
}

// NewListCloudFrontOriginAccessIdentitiesPaginator returns a new
// ListCloudFrontOriginAccessIdentitiesPaginator
func NewListCloudFrontOriginAccessIdentitiesPaginator(client ListCloudFrontOriginAccessIdentitiesAPIClient, params *ListCloudFrontOriginAccessIdentitiesInput, optFns ...func(*ListCloudFrontOriginAccessIdentitiesPaginatorOptions)) *ListCloudFrontOriginAccessIdentitiesPaginator {
	options := ListCloudFrontOriginAccessIdentitiesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListCloudFrontOriginAccessIdentitiesInput{}
	}

	return &ListCloudFrontOriginAccessIdentitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCloudFrontOriginAccessIdentitiesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCloudFrontOriginAccessIdentities page.
func (p *ListCloudFrontOriginAccessIdentitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCloudFrontOriginAccessIdentitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListCloudFrontOriginAccessIdentities(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.CloudFrontOriginAccessIdentityList != nil {
		p.nextToken = result.CloudFrontOriginAccessIdentityList.NextMarker
	}

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCloudFrontOriginAccessIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListCloudFrontOriginAccessIdentities",
	}
}
