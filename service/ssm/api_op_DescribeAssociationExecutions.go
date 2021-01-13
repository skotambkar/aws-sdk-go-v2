// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this API action to view all executions for a specific association ID.
func (c *Client) DescribeAssociationExecutions(ctx context.Context, params *DescribeAssociationExecutionsInput, optFns ...func(*Options)) (*DescribeAssociationExecutionsOutput, error) {
	if params == nil {
		params = &DescribeAssociationExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssociationExecutions", params, optFns, addOperationDescribeAssociationExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssociationExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssociationExecutionsInput struct {

	// The association ID for which you want to view execution history details.
	//
	// This member is required.
	AssociationId *string

	// Filters for the request. You can specify the following filters and values.
	// ExecutionId (EQUAL) Status (EQUAL) CreatedTime (EQUAL, GREATER_THAN, LESS_THAN)
	Filters []types.AssociationExecutionFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type DescribeAssociationExecutionsOutput struct {

	// A list of the executions for the specified association ID.
	AssociationExecutions []types.AssociationExecution

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAssociationExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssociationExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssociationExecutions{}, middleware.After)
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
	if err = addOpDescribeAssociationExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssociationExecutions(options.Region), middleware.Before); err != nil {
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

// DescribeAssociationExecutionsAPIClient is a client that implements the
// DescribeAssociationExecutions operation.
type DescribeAssociationExecutionsAPIClient interface {
	DescribeAssociationExecutions(context.Context, *DescribeAssociationExecutionsInput, ...func(*Options)) (*DescribeAssociationExecutionsOutput, error)
}

var _ DescribeAssociationExecutionsAPIClient = (*Client)(nil)

// DescribeAssociationExecutionsPaginatorOptions is the paginator options for
// DescribeAssociationExecutions
type DescribeAssociationExecutionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAssociationExecutionsPaginator is a paginator for
// DescribeAssociationExecutions
type DescribeAssociationExecutionsPaginator struct {
	options   DescribeAssociationExecutionsPaginatorOptions
	client    DescribeAssociationExecutionsAPIClient
	params    *DescribeAssociationExecutionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAssociationExecutionsPaginator returns a new
// DescribeAssociationExecutionsPaginator
func NewDescribeAssociationExecutionsPaginator(client DescribeAssociationExecutionsAPIClient, params *DescribeAssociationExecutionsInput, optFns ...func(*DescribeAssociationExecutionsPaginatorOptions)) *DescribeAssociationExecutionsPaginator {
	options := DescribeAssociationExecutionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeAssociationExecutionsInput{}
	}

	return &DescribeAssociationExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAssociationExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeAssociationExecutions page.
func (p *DescribeAssociationExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAssociationExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeAssociationExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAssociationExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeAssociationExecutions",
	}
}
