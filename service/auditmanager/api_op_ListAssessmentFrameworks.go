// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the frameworks available in the AWS Audit Manager framework
// library.
func (c *Client) ListAssessmentFrameworks(ctx context.Context, params *ListAssessmentFrameworksInput, optFns ...func(*Options)) (*ListAssessmentFrameworksOutput, error) {
	if params == nil {
		params = &ListAssessmentFrameworksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentFrameworks", params, optFns, c.addOperationListAssessmentFrameworksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentFrameworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentFrameworksInput struct {

	// The type of framework, such as standard or custom.
	//
	// This member is required.
	FrameworkType types.FrameworkType

	// Represents the maximum number of results per page, or per API request call.
	MaxResults *int32

	// The pagination token used to fetch the next set of results.
	NextToken *string
}

type ListAssessmentFrameworksOutput struct {

	// The list of metadata objects for the specified framework.
	FrameworkMetadataList []types.AssessmentFrameworkMetadata

	// The pagination token used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListAssessmentFrameworksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssessmentFrameworks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssessmentFrameworks{}, middleware.After)
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
	if err = addOpListAssessmentFrameworksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentFrameworks(options.Region), middleware.Before); err != nil {
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

// ListAssessmentFrameworksAPIClient is a client that implements the
// ListAssessmentFrameworks operation.
type ListAssessmentFrameworksAPIClient interface {
	ListAssessmentFrameworks(context.Context, *ListAssessmentFrameworksInput, ...func(*Options)) (*ListAssessmentFrameworksOutput, error)
}

var _ ListAssessmentFrameworksAPIClient = (*Client)(nil)

// ListAssessmentFrameworksPaginatorOptions is the paginator options for
// ListAssessmentFrameworks
type ListAssessmentFrameworksPaginatorOptions struct {
	// Represents the maximum number of results per page, or per API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentFrameworksPaginator is a paginator for ListAssessmentFrameworks
type ListAssessmentFrameworksPaginator struct {
	options   ListAssessmentFrameworksPaginatorOptions
	client    ListAssessmentFrameworksAPIClient
	params    *ListAssessmentFrameworksInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentFrameworksPaginator returns a new
// ListAssessmentFrameworksPaginator
func NewListAssessmentFrameworksPaginator(client ListAssessmentFrameworksAPIClient, params *ListAssessmentFrameworksInput, optFns ...func(*ListAssessmentFrameworksPaginatorOptions)) *ListAssessmentFrameworksPaginator {
	if params == nil {
		params = &ListAssessmentFrameworksInput{}
	}

	options := ListAssessmentFrameworksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentFrameworksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentFrameworksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAssessmentFrameworks page.
func (p *ListAssessmentFrameworksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentFrameworksOutput, error) {
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

	result, err := p.client.ListAssessmentFrameworks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssessmentFrameworks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "ListAssessmentFrameworks",
	}
}
