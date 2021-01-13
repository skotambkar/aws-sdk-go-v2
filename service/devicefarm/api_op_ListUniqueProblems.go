// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about unique problems, such as exceptions or crashes. Unique
// problems are defined as a single instance of an error across a run, job, or
// suite. For example, if a call in your application consistently raises an
// exception (OutOfBoundsException in MyActivity.java:386), ListUniqueProblems
// returns a single entry instead of many individual entries for that exception.
func (c *Client) ListUniqueProblems(ctx context.Context, params *ListUniqueProblemsInput, optFns ...func(*Options)) (*ListUniqueProblemsOutput, error) {
	if params == nil {
		params = &ListUniqueProblemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUniqueProblems", params, optFns, addOperationListUniqueProblemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUniqueProblemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the list unique problems operation.
type ListUniqueProblemsInput struct {

	// The unique problems' ARNs.
	//
	// This member is required.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the result of a list unique problems request.
type ListUniqueProblemsOutput struct {

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Information about the unique problems. Allowed values include:
	//
	// * PENDING
	//
	// *
	// PASSED
	//
	// * WARNED
	//
	// * FAILED
	//
	// * SKIPPED
	//
	// * ERRORED
	//
	// * STOPPED
	UniqueProblems map[string][]types.UniqueProblem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUniqueProblemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUniqueProblems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUniqueProblems{}, middleware.After)
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
	if err = addOpListUniqueProblemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUniqueProblems(options.Region), middleware.Before); err != nil {
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

// ListUniqueProblemsAPIClient is a client that implements the ListUniqueProblems
// operation.
type ListUniqueProblemsAPIClient interface {
	ListUniqueProblems(context.Context, *ListUniqueProblemsInput, ...func(*Options)) (*ListUniqueProblemsOutput, error)
}

var _ ListUniqueProblemsAPIClient = (*Client)(nil)

// ListUniqueProblemsPaginatorOptions is the paginator options for
// ListUniqueProblems
type ListUniqueProblemsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUniqueProblemsPaginator is a paginator for ListUniqueProblems
type ListUniqueProblemsPaginator struct {
	options   ListUniqueProblemsPaginatorOptions
	client    ListUniqueProblemsAPIClient
	params    *ListUniqueProblemsInput
	nextToken *string
	firstPage bool
}

// NewListUniqueProblemsPaginator returns a new ListUniqueProblemsPaginator
func NewListUniqueProblemsPaginator(client ListUniqueProblemsAPIClient, params *ListUniqueProblemsInput, optFns ...func(*ListUniqueProblemsPaginatorOptions)) *ListUniqueProblemsPaginator {
	options := ListUniqueProblemsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListUniqueProblemsInput{}
	}

	return &ListUniqueProblemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUniqueProblemsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListUniqueProblems page.
func (p *ListUniqueProblemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUniqueProblemsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListUniqueProblems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUniqueProblems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListUniqueProblems",
	}
}
