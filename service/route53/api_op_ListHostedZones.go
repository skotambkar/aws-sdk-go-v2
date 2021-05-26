// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of the public and private hosted zones that are associated with
// the current AWS account. The response includes a HostedZones child element for
// each hosted zone. Amazon Route 53 returns a maximum of 100 items in each
// response. If you have a lot of hosted zones, you can use the maxitems parameter
// to list them in groups of up to 100.
func (c *Client) ListHostedZones(ctx context.Context, params *ListHostedZonesInput, optFns ...func(*Options)) (*ListHostedZonesOutput, error) {
	if params == nil {
		params = &ListHostedZonesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHostedZones", params, optFns, c.addOperationListHostedZonesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHostedZonesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve a list of the public and private hosted zones that are
// associated with the current AWS account.
type ListHostedZonesInput struct {

	// If you're using reusable delegation sets and you want to list all of the hosted
	// zones that are associated with a reusable delegation set, specify the ID of that
	// reusable delegation set.
	DelegationSetId *string

	// If the value of IsTruncated in the previous response was true, you have more
	// hosted zones. To get more hosted zones, submit another ListHostedZones request.
	// For the value of marker, specify the value of NextMarker from the previous
	// response, which is the ID of the first hosted zone that Amazon Route 53 will
	// return if you submit another request. If the value of IsTruncated in the
	// previous response was false, there are no more hosted zones to get.
	Marker *string

	// (Optional) The maximum number of hosted zones that you want Amazon Route 53 to
	// return. If you have more than maxitems hosted zones, the value of IsTruncated in
	// the response is true, and the value of NextMarker is the hosted zone ID of the
	// first hosted zone that Route 53 will return if you submit another request.
	MaxItems *int32
}

type ListHostedZonesOutput struct {

	// A complex type that contains general information about the hosted zone.
	//
	// This member is required.
	HostedZones []types.HostedZone

	// A flag indicating whether there are more hosted zones to be listed. If the
	// response was truncated, you can get more hosted zones by submitting another
	// ListHostedZones request and specifying the value of NextMarker in the marker
	// parameter.
	//
	// This member is required.
	IsTruncated bool

	// For the second and subsequent calls to ListHostedZones, Marker is the value that
	// you specified for the marker parameter in the request that produced the current
	// response.
	//
	// This member is required.
	Marker *string

	// The value that you specified for the maxitems parameter in the call to
	// ListHostedZones that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// If IsTruncated is true, the value of NextMarker identifies the first hosted zone
	// in the next group of hosted zones. Submit another ListHostedZones request, and
	// specify the value of NextMarker from the response in the marker parameter. This
	// element is present only if IsTruncated is true.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListHostedZonesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListHostedZones{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListHostedZones{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHostedZones(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListHostedZonesAPIClient is a client that implements the ListHostedZones
// operation.
type ListHostedZonesAPIClient interface {
	ListHostedZones(context.Context, *ListHostedZonesInput, ...func(*Options)) (*ListHostedZonesOutput, error)
}

var _ ListHostedZonesAPIClient = (*Client)(nil)

// ListHostedZonesPaginatorOptions is the paginator options for ListHostedZones
type ListHostedZonesPaginatorOptions struct {
	// (Optional) The maximum number of hosted zones that you want Amazon Route 53 to
	// return. If you have more than maxitems hosted zones, the value of IsTruncated in
	// the response is true, and the value of NextMarker is the hosted zone ID of the
	// first hosted zone that Route 53 will return if you submit another request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHostedZonesPaginator is a paginator for ListHostedZones
type ListHostedZonesPaginator struct {
	options   ListHostedZonesPaginatorOptions
	client    ListHostedZonesAPIClient
	params    *ListHostedZonesInput
	nextToken *string
	firstPage bool
}

// NewListHostedZonesPaginator returns a new ListHostedZonesPaginator
func NewListHostedZonesPaginator(client ListHostedZonesAPIClient, params *ListHostedZonesInput, optFns ...func(*ListHostedZonesPaginatorOptions)) *ListHostedZonesPaginator {
	if params == nil {
		params = &ListHostedZonesInput{}
	}

	options := ListHostedZonesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHostedZonesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHostedZonesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListHostedZones page.
func (p *ListHostedZonesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHostedZonesOutput, error) {
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

	result, err := p.client.ListHostedZones(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHostedZones(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListHostedZones",
	}
}
