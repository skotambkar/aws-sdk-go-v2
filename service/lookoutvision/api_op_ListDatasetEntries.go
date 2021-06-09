// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the JSON Lines within a dataset. An Amazon Lookout for Vision JSON Line
// contains the anomaly information for a single image, including the image
// location and the assigned label. This operation requires permissions to perform
// the lookoutvision:ListDatasetEntries operation.
func (c *Client) ListDatasetEntries(ctx context.Context, params *ListDatasetEntriesInput, optFns ...func(*Options)) (*ListDatasetEntriesOutput, error) {
	if params == nil {
		params = &ListDatasetEntriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasetEntries", params, optFns, c.addOperationListDatasetEntriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetEntriesInput struct {

	// The type of the dataset that you want to list. Specify train to list the
	// training dataset. Specify test to list the test dataset. If you have a single
	// dataset project, specify train.
	//
	// This member is required.
	DatasetType *string

	// The name of the project that contains the dataset that you want to list.
	//
	// This member is required.
	ProjectName *string

	// Only includes entries after the specified date in the response. For example,
	// 2020-06-23T00:00:00.
	AfterCreationDate *time.Time

	// Specify normal to include only normal images. Specify anomaly to only include
	// anomalous entries. If you don't specify a value, Amazon Lookout for Vision
	// returns normal and anomalous images.
	AnomalyClass *string

	// Only includes entries before the specified date in the response. For example,
	// 2020-06-23T00:00:00.
	BeforeCreationDate *time.Time

	// Specify true to include labeled entries, otherwise specify false. If you don't
	// specify a value, Lookout for Vision returns all entries.
	Labeled *bool

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Lookout for Vision returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of dataset entries.
	NextToken *string

	// Perform a "contains" search on the values of the source-ref key within the
	// dataset. For example a value of "IMG_17" returns all JSON Lines where the
	// source-ref key value matches *IMG_17*.
	SourceRefContains *string
}

type ListDatasetEntriesOutput struct {

	// A list of the entries (JSON Lines) within the dataset.
	DatasetEntries []string

	// If the response is truncated, Amazon Lookout for Vision returns this token that
	// you can use in the subsequent request to retrieve the next set ofdataset
	// entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListDatasetEntriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDatasetEntries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDatasetEntries{}, middleware.After)
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
	if err = addOpListDatasetEntriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetEntries(options.Region), middleware.Before); err != nil {
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

// ListDatasetEntriesAPIClient is a client that implements the ListDatasetEntries
// operation.
type ListDatasetEntriesAPIClient interface {
	ListDatasetEntries(context.Context, *ListDatasetEntriesInput, ...func(*Options)) (*ListDatasetEntriesOutput, error)
}

var _ ListDatasetEntriesAPIClient = (*Client)(nil)

// ListDatasetEntriesPaginatorOptions is the paginator options for
// ListDatasetEntries
type ListDatasetEntriesPaginatorOptions struct {
	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatasetEntriesPaginator is a paginator for ListDatasetEntries
type ListDatasetEntriesPaginator struct {
	options   ListDatasetEntriesPaginatorOptions
	client    ListDatasetEntriesAPIClient
	params    *ListDatasetEntriesInput
	nextToken *string
	firstPage bool
}

// NewListDatasetEntriesPaginator returns a new ListDatasetEntriesPaginator
func NewListDatasetEntriesPaginator(client ListDatasetEntriesAPIClient, params *ListDatasetEntriesInput, optFns ...func(*ListDatasetEntriesPaginatorOptions)) *ListDatasetEntriesPaginator {
	if params == nil {
		params = &ListDatasetEntriesInput{}
	}

	options := ListDatasetEntriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatasetEntriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatasetEntriesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDatasetEntries page.
func (p *ListDatasetEntriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatasetEntriesOutput, error) {
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

	result, err := p.client.ListDatasetEntries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatasetEntries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "ListDatasetEntries",
	}
}
