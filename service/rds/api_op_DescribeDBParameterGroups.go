// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DBParameterGroup descriptions. If a DBParameterGroupName is
// specified, the list will contain only the description of the specified DB
// parameter group.
func (c *Client) DescribeDBParameterGroups(ctx context.Context, params *DescribeDBParameterGroupsInput, optFns ...func(*Options)) (*DescribeDBParameterGroupsOutput, error) {
	if params == nil {
		params = &DescribeDBParameterGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBParameterGroups", params, optFns, c.addOperationDescribeDBParameterGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBParameterGroupsInput struct {

	// The name of a specific DB parameter group to return details for. Constraints:
	//
	// *
	// If supplied, must match the name of an existing DBClusterParameterGroup.
	DBParameterGroupName *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBParameterGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the DescribeDBParameterGroups
// action.
type DescribeDBParameterGroupsOutput struct {

	// A list of DBParameterGroup instances.
	DBParameterGroups []types.DBParameterGroup

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeDBParameterGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBParameterGroups{}, middleware.After)
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
	if err = addOpDescribeDBParameterGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBParameterGroups(options.Region), middleware.Before); err != nil {
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

// DescribeDBParameterGroupsAPIClient is a client that implements the
// DescribeDBParameterGroups operation.
type DescribeDBParameterGroupsAPIClient interface {
	DescribeDBParameterGroups(context.Context, *DescribeDBParameterGroupsInput, ...func(*Options)) (*DescribeDBParameterGroupsOutput, error)
}

var _ DescribeDBParameterGroupsAPIClient = (*Client)(nil)

// DescribeDBParameterGroupsPaginatorOptions is the paginator options for
// DescribeDBParameterGroups
type DescribeDBParameterGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBParameterGroupsPaginator is a paginator for DescribeDBParameterGroups
type DescribeDBParameterGroupsPaginator struct {
	options   DescribeDBParameterGroupsPaginatorOptions
	client    DescribeDBParameterGroupsAPIClient
	params    *DescribeDBParameterGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBParameterGroupsPaginator returns a new
// DescribeDBParameterGroupsPaginator
func NewDescribeDBParameterGroupsPaginator(client DescribeDBParameterGroupsAPIClient, params *DescribeDBParameterGroupsInput, optFns ...func(*DescribeDBParameterGroupsPaginatorOptions)) *DescribeDBParameterGroupsPaginator {
	if params == nil {
		params = &DescribeDBParameterGroupsInput{}
	}

	options := DescribeDBParameterGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBParameterGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBParameterGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeDBParameterGroups page.
func (p *DescribeDBParameterGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBParameterGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBParameterGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBParameterGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBParameterGroups",
	}
}
