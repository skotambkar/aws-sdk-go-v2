// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets revisions made to the specified system template. Only the previous 100
// revisions are stored. If the system has been deprecated, this action will return
// the revisions that occurred before its deprecation. This action won't work with
// systems that have been deleted.
func (c *Client) GetSystemTemplateRevisions(ctx context.Context, params *GetSystemTemplateRevisionsInput, optFns ...func(*Options)) (*GetSystemTemplateRevisionsOutput, error) {
	if params == nil {
		params = &GetSystemTemplateRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSystemTemplateRevisions", params, optFns, c.addOperationGetSystemTemplateRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSystemTemplateRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSystemTemplateRevisionsInput struct {

	// The ID of the system template. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type GetSystemTemplateRevisionsOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary data about the system template
	// revisions.
	Summaries []types.SystemTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetSystemTemplateRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSystemTemplateRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSystemTemplateRevisions{}, middleware.After)
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
	if err = addOpGetSystemTemplateRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSystemTemplateRevisions(options.Region), middleware.Before); err != nil {
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

// GetSystemTemplateRevisionsAPIClient is a client that implements the
// GetSystemTemplateRevisions operation.
type GetSystemTemplateRevisionsAPIClient interface {
	GetSystemTemplateRevisions(context.Context, *GetSystemTemplateRevisionsInput, ...func(*Options)) (*GetSystemTemplateRevisionsOutput, error)
}

var _ GetSystemTemplateRevisionsAPIClient = (*Client)(nil)

// GetSystemTemplateRevisionsPaginatorOptions is the paginator options for
// GetSystemTemplateRevisions
type GetSystemTemplateRevisionsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSystemTemplateRevisionsPaginator is a paginator for
// GetSystemTemplateRevisions
type GetSystemTemplateRevisionsPaginator struct {
	options   GetSystemTemplateRevisionsPaginatorOptions
	client    GetSystemTemplateRevisionsAPIClient
	params    *GetSystemTemplateRevisionsInput
	nextToken *string
	firstPage bool
}

// NewGetSystemTemplateRevisionsPaginator returns a new
// GetSystemTemplateRevisionsPaginator
func NewGetSystemTemplateRevisionsPaginator(client GetSystemTemplateRevisionsAPIClient, params *GetSystemTemplateRevisionsInput, optFns ...func(*GetSystemTemplateRevisionsPaginatorOptions)) *GetSystemTemplateRevisionsPaginator {
	if params == nil {
		params = &GetSystemTemplateRevisionsInput{}
	}

	options := GetSystemTemplateRevisionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSystemTemplateRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSystemTemplateRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetSystemTemplateRevisions page.
func (p *GetSystemTemplateRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSystemTemplateRevisionsOutput, error) {
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

	result, err := p.client.GetSystemTemplateRevisions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSystemTemplateRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetSystemTemplateRevisions",
	}
}
