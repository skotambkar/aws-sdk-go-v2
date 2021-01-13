// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of asset relationships for an asset. You can use this
// operation to identify an asset's root asset and all associated assets between
// that asset and its root.
func (c *Client) ListAssetRelationships(ctx context.Context, params *ListAssetRelationshipsInput, optFns ...func(*Options)) (*ListAssetRelationshipsOutput, error) {
	if params == nil {
		params = &ListAssetRelationshipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetRelationships", params, optFns, addOperationListAssetRelationshipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetRelationshipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetRelationshipsInput struct {

	// The ID of the asset.
	//
	// This member is required.
	AssetId *string

	// The type of traversal to use to identify asset relationships. Choose the
	// following option:
	//
	// * PATH_TO_ROOT – Identify the asset's parent assets up to the
	// root asset. The asset that you specify in assetId is the first result in the
	// list of assetRelationshipSummaries, and the root asset is the last result.
	//
	// This member is required.
	TraversalType types.TraversalType

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListAssetRelationshipsOutput struct {

	// A list that summarizes each asset relationship.
	//
	// This member is required.
	AssetRelationshipSummaries []types.AssetRelationshipSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAssetRelationshipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetRelationships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetRelationships{}, middleware.After)
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
	if err = addEndpointPrefix_opListAssetRelationshipsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAssetRelationshipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetRelationships(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListAssetRelationshipsMiddleware struct {
}

func (*endpointPrefix_opListAssetRelationshipsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAssetRelationshipsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "model." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListAssetRelationshipsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListAssetRelationshipsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListAssetRelationshipsAPIClient is a client that implements the
// ListAssetRelationships operation.
type ListAssetRelationshipsAPIClient interface {
	ListAssetRelationships(context.Context, *ListAssetRelationshipsInput, ...func(*Options)) (*ListAssetRelationshipsOutput, error)
}

var _ ListAssetRelationshipsAPIClient = (*Client)(nil)

// ListAssetRelationshipsPaginatorOptions is the paginator options for
// ListAssetRelationships
type ListAssetRelationshipsPaginatorOptions struct {
	// The maximum number of results to be returned per paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetRelationshipsPaginator is a paginator for ListAssetRelationships
type ListAssetRelationshipsPaginator struct {
	options   ListAssetRelationshipsPaginatorOptions
	client    ListAssetRelationshipsAPIClient
	params    *ListAssetRelationshipsInput
	nextToken *string
	firstPage bool
}

// NewListAssetRelationshipsPaginator returns a new ListAssetRelationshipsPaginator
func NewListAssetRelationshipsPaginator(client ListAssetRelationshipsAPIClient, params *ListAssetRelationshipsInput, optFns ...func(*ListAssetRelationshipsPaginatorOptions)) *ListAssetRelationshipsPaginator {
	options := ListAssetRelationshipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListAssetRelationshipsInput{}
	}

	return &ListAssetRelationshipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetRelationshipsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAssetRelationships page.
func (p *ListAssetRelationshipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetRelationshipsOutput, error) {
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

	result, err := p.client.ListAssetRelationships(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssetRelationships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListAssetRelationships",
	}
}
