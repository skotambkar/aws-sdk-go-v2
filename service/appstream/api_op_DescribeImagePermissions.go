// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list that describes the permissions for shared AWS account IDs on a
// private image that you own.
func (c *Client) DescribeImagePermissions(ctx context.Context, params *DescribeImagePermissionsInput, optFns ...func(*Options)) (*DescribeImagePermissionsOutput, error) {
	if params == nil {
		params = &DescribeImagePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImagePermissions", params, optFns, c.addOperationDescribeImagePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImagePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImagePermissionsInput struct {

	// The name of the private image for which to describe permissions. The image must
	// be one that you own.
	//
	// This member is required.
	Name *string

	// The maximum size of each page of results.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The 12-digit identifier of one or more AWS accounts with which the image is
	// shared.
	SharedAwsAccountIds []string
}

type DescribeImagePermissionsOutput struct {

	// The name of the private image.
	Name *string

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// The permissions for a private image that you own.
	SharedImagePermissionsList []types.SharedImagePermissions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeImagePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImagePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImagePermissions{}, middleware.After)
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
	if err = addOpDescribeImagePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImagePermissions(options.Region), middleware.Before); err != nil {
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

// DescribeImagePermissionsAPIClient is a client that implements the
// DescribeImagePermissions operation.
type DescribeImagePermissionsAPIClient interface {
	DescribeImagePermissions(context.Context, *DescribeImagePermissionsInput, ...func(*Options)) (*DescribeImagePermissionsOutput, error)
}

var _ DescribeImagePermissionsAPIClient = (*Client)(nil)

// DescribeImagePermissionsPaginatorOptions is the paginator options for
// DescribeImagePermissions
type DescribeImagePermissionsPaginatorOptions struct {
	// The maximum size of each page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeImagePermissionsPaginator is a paginator for DescribeImagePermissions
type DescribeImagePermissionsPaginator struct {
	options   DescribeImagePermissionsPaginatorOptions
	client    DescribeImagePermissionsAPIClient
	params    *DescribeImagePermissionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeImagePermissionsPaginator returns a new
// DescribeImagePermissionsPaginator
func NewDescribeImagePermissionsPaginator(client DescribeImagePermissionsAPIClient, params *DescribeImagePermissionsInput, optFns ...func(*DescribeImagePermissionsPaginatorOptions)) *DescribeImagePermissionsPaginator {
	if params == nil {
		params = &DescribeImagePermissionsInput{}
	}

	options := DescribeImagePermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeImagePermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeImagePermissionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeImagePermissions page.
func (p *DescribeImagePermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeImagePermissionsOutput, error) {
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

	result, err := p.client.DescribeImagePermissions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeImagePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DescribeImagePermissions",
	}
}
