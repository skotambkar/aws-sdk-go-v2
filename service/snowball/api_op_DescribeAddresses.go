// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a specified number of ADDRESS objects. Calling this API in one of the US
// regions will return addresses from the list of all addresses associated with
// this account in all US regions.
func (c *Client) DescribeAddresses(ctx context.Context, params *DescribeAddressesInput, optFns ...func(*Options)) (*DescribeAddressesOutput, error) {
	if params == nil {
		params = &DescribeAddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAddresses", params, optFns, addOperationDescribeAddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddressesInput struct {

	// The number of ADDRESS objects to return.
	MaxResults *int32

	// HTTP requests are stateless. To identify what object comes "next" in the list of
	// ADDRESS objects, you have the option of specifying a value for NextToken as the
	// starting point for your list of returned addresses.
	NextToken *string
}

type DescribeAddressesOutput struct {

	// The Snow device shipping addresses that were created for this account.
	Addresses []types.Address

	// HTTP requests are stateless. If you use the automatically generated NextToken
	// value in your next DescribeAddresses call, your list of returned addresses will
	// start from this point in the array.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAddresses{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddresses(options.Region), middleware.Before); err != nil {
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

// DescribeAddressesAPIClient is a client that implements the DescribeAddresses
// operation.
type DescribeAddressesAPIClient interface {
	DescribeAddresses(context.Context, *DescribeAddressesInput, ...func(*Options)) (*DescribeAddressesOutput, error)
}

var _ DescribeAddressesAPIClient = (*Client)(nil)

// DescribeAddressesPaginatorOptions is the paginator options for DescribeAddresses
type DescribeAddressesPaginatorOptions struct {
	// The number of ADDRESS objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAddressesPaginator is a paginator for DescribeAddresses
type DescribeAddressesPaginator struct {
	options   DescribeAddressesPaginatorOptions
	client    DescribeAddressesAPIClient
	params    *DescribeAddressesInput
	nextToken *string
	firstPage bool
}

// NewDescribeAddressesPaginator returns a new DescribeAddressesPaginator
func NewDescribeAddressesPaginator(client DescribeAddressesAPIClient, params *DescribeAddressesInput, optFns ...func(*DescribeAddressesPaginatorOptions)) *DescribeAddressesPaginator {
	options := DescribeAddressesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeAddressesInput{}
	}

	return &DescribeAddressesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAddressesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeAddresses page.
func (p *DescribeAddressesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAddressesOutput, error) {
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

	result, err := p.client.DescribeAddresses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAddresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "DescribeAddresses",
	}
}
