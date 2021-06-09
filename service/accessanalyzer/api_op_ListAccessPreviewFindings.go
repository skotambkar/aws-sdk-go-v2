// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of access preview findings generated by the specified access
// preview.
func (c *Client) ListAccessPreviewFindings(ctx context.Context, params *ListAccessPreviewFindingsInput, optFns ...func(*Options)) (*ListAccessPreviewFindingsOutput, error) {
	if params == nil {
		params = &ListAccessPreviewFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessPreviewFindings", params, optFns, c.addOperationListAccessPreviewFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessPreviewFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessPreviewFindingsInput struct {

	// The unique ID for the access preview.
	//
	// This member is required.
	AccessPreviewId *string

	// The ARN of the analyzer
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources)
	// used to generate the access.
	//
	// This member is required.
	AnalyzerArn *string

	// Criteria to filter the returned findings.
	Filter map[string]types.Criterion

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string
}

type ListAccessPreviewFindingsOutput struct {

	// A list of access preview findings that match the specified filter criteria.
	//
	// This member is required.
	Findings []types.AccessPreviewFinding

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListAccessPreviewFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccessPreviewFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccessPreviewFindings{}, middleware.After)
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
	if err = addOpListAccessPreviewFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessPreviewFindings(options.Region), middleware.Before); err != nil {
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

// ListAccessPreviewFindingsAPIClient is a client that implements the
// ListAccessPreviewFindings operation.
type ListAccessPreviewFindingsAPIClient interface {
	ListAccessPreviewFindings(context.Context, *ListAccessPreviewFindingsInput, ...func(*Options)) (*ListAccessPreviewFindingsOutput, error)
}

var _ ListAccessPreviewFindingsAPIClient = (*Client)(nil)

// ListAccessPreviewFindingsPaginatorOptions is the paginator options for
// ListAccessPreviewFindings
type ListAccessPreviewFindingsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessPreviewFindingsPaginator is a paginator for ListAccessPreviewFindings
type ListAccessPreviewFindingsPaginator struct {
	options   ListAccessPreviewFindingsPaginatorOptions
	client    ListAccessPreviewFindingsAPIClient
	params    *ListAccessPreviewFindingsInput
	nextToken *string
	firstPage bool
}

// NewListAccessPreviewFindingsPaginator returns a new
// ListAccessPreviewFindingsPaginator
func NewListAccessPreviewFindingsPaginator(client ListAccessPreviewFindingsAPIClient, params *ListAccessPreviewFindingsInput, optFns ...func(*ListAccessPreviewFindingsPaginatorOptions)) *ListAccessPreviewFindingsPaginator {
	if params == nil {
		params = &ListAccessPreviewFindingsInput{}
	}

	options := ListAccessPreviewFindingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessPreviewFindingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessPreviewFindingsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAccessPreviewFindings page.
func (p *ListAccessPreviewFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessPreviewFindingsOutput, error) {
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

	result, err := p.client.ListAccessPreviewFindings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccessPreviewFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "ListAccessPreviewFindings",
	}
}
