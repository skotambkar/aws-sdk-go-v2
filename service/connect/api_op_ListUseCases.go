// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the use cases.
func (c *Client) ListUseCases(ctx context.Context, params *ListUseCasesInput, optFns ...func(*Options)) (*ListUseCasesOutput, error) {
	if params == nil {
		params = &ListUseCasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUseCases", params, optFns, c.addOperationListUseCasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUseCasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides summary information about the use cases for the specified Amazon
// Connect AppIntegration association.
type ListUseCasesInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier for the integration association.
	//
	// This member is required.
	IntegrationAssociationId *string

	// The maximum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
}

type ListUseCasesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The use cases.
	UseCaseSummaryList []types.UseCase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListUseCasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUseCases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUseCases{}, middleware.After)
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
	if err = addOpListUseCasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUseCases(options.Region), middleware.Before); err != nil {
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

// ListUseCasesAPIClient is a client that implements the ListUseCases operation.
type ListUseCasesAPIClient interface {
	ListUseCases(context.Context, *ListUseCasesInput, ...func(*Options)) (*ListUseCasesOutput, error)
}

var _ ListUseCasesAPIClient = (*Client)(nil)

// ListUseCasesPaginatorOptions is the paginator options for ListUseCases
type ListUseCasesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUseCasesPaginator is a paginator for ListUseCases
type ListUseCasesPaginator struct {
	options   ListUseCasesPaginatorOptions
	client    ListUseCasesAPIClient
	params    *ListUseCasesInput
	nextToken *string
	firstPage bool
}

// NewListUseCasesPaginator returns a new ListUseCasesPaginator
func NewListUseCasesPaginator(client ListUseCasesAPIClient, params *ListUseCasesInput, optFns ...func(*ListUseCasesPaginatorOptions)) *ListUseCasesPaginator {
	if params == nil {
		params = &ListUseCasesInput{}
	}

	options := ListUseCasesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUseCasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUseCasesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListUseCases page.
func (p *ListUseCasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUseCasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListUseCases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUseCases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListUseCases",
	}
}
