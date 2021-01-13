// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/honeycode/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListTables API allows you to retrieve a list of all the tables in a
// workbook.
func (c *Client) ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) {
	if params == nil {
		params = &ListTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTables", params, optFns, addOperationListTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTablesInput struct {

	// The ID of the workbook whose tables are being retrieved. If a workbook with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	// The maximum number of tables to return in each page of the results.
	MaxResults *int32

	// This parameter is optional. If a nextToken is not specified, the API returns the
	// first page of data. Pagination tokens expire after 1 hour. If you use a token
	// that was returned more than an hour back, the API will throw
	// ValidationException.
	NextToken *string
}

type ListTablesOutput struct {

	// The list of tables in the workbook.
	//
	// This member is required.
	Tables []types.Table

	// Provides the pagination token to load the next page if there are more results
	// matching the request. If a pagination token is not present in the response, it
	// means that all data matching the request has been loaded.
	NextToken *string

	// Indicates the cursor of the workbook at which the data returned by this request
	// is read. Workbook cursor keeps increasing with every update and the increments
	// are not sequential.
	WorkbookCursor int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTables{}, middleware.After)
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
	if err = addOpListTablesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTables(options.Region), middleware.Before); err != nil {
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

// ListTablesAPIClient is a client that implements the ListTables operation.
type ListTablesAPIClient interface {
	ListTables(context.Context, *ListTablesInput, ...func(*Options)) (*ListTablesOutput, error)
}

var _ ListTablesAPIClient = (*Client)(nil)

// ListTablesPaginatorOptions is the paginator options for ListTables
type ListTablesPaginatorOptions struct {
	// The maximum number of tables to return in each page of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTablesPaginator is a paginator for ListTables
type ListTablesPaginator struct {
	options   ListTablesPaginatorOptions
	client    ListTablesAPIClient
	params    *ListTablesInput
	nextToken *string
	firstPage bool
}

// NewListTablesPaginator returns a new ListTablesPaginator
func NewListTablesPaginator(client ListTablesAPIClient, params *ListTablesInput, optFns ...func(*ListTablesPaginatorOptions)) *ListTablesPaginator {
	options := ListTablesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTablesInput{}
	}

	return &ListTablesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTablesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTables page.
func (p *ListTablesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTablesOutput, error) {
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

	result, err := p.client.ListTables(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "honeycode",
		OperationName: "ListTables",
	}
}
