// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of cases that you specify by passing one or more case IDs. You
// can use the afterTime and beforeTime parameters to filter the cases by date. You
// can set values for the includeResolvedCases and includeCommunications parameters
// to specify how much information to return. The response returns the following in
// JSON format:
//
// * One or more CaseDetails
// (https://docs.aws.amazon.com/awssupport/latest/APIReference/API_CaseDetails.html)
// data types.
//
// * One or more nextToken values, which specify where to paginate the
// returned records represented by the CaseDetails objects.
//
// Case data is available
// for 12 months after creation. If a case was created more than 12 months ago, a
// request might return an error.
//
// * You must have a Business or Enterprise Support
// plan to use the AWS Support API.
//
// * If you call the AWS Support API from an
// account that does not have a Business or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeCases(ctx context.Context, params *DescribeCasesInput, optFns ...func(*Options)) (*DescribeCasesOutput, error) {
	if params == nil {
		params = &DescribeCasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCases", params, optFns, c.addOperationDescribeCasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCasesInput struct {

	// The start date for a filtered date search on support case communications. Case
	// communications are available for 12 months after creation.
	AfterTime *string

	// The end date for a filtered date search on support case communications. Case
	// communications are available for 12 months after creation.
	BeforeTime *string

	// A list of ID numbers of the support cases you want returned. The maximum number
	// of cases is 100.
	CaseIdList []string

	// The ID displayed for a case in the AWS Support Center user interface.
	DisplayId *string

	// Specifies whether to include communications in the DescribeCases response. By
	// default, communications are included.
	IncludeCommunications *bool

	// Specifies whether to include resolved support cases in the DescribeCases
	// response. By default, resolved cases aren't included.
	IncludeResolvedCases bool

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters must
	// be passed explicitly for operations that take them.
	Language *string

	// The maximum number of results to return before paginating.
	MaxResults *int32

	// A resumption point for pagination.
	NextToken *string
}

// Returns an array of CaseDetails
// (https://docs.aws.amazon.com/awssupport/latest/APIReference/API_CaseDetails.html)
// objects and a nextToken that defines a point for pagination in the result set.
type DescribeCasesOutput struct {

	// The details for the cases that match the request.
	Cases []types.CaseDetails

	// A resumption point for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeCasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCases{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCases(options.Region), middleware.Before); err != nil {
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

// DescribeCasesAPIClient is a client that implements the DescribeCases operation.
type DescribeCasesAPIClient interface {
	DescribeCases(context.Context, *DescribeCasesInput, ...func(*Options)) (*DescribeCasesOutput, error)
}

var _ DescribeCasesAPIClient = (*Client)(nil)

// DescribeCasesPaginatorOptions is the paginator options for DescribeCases
type DescribeCasesPaginatorOptions struct {
	// The maximum number of results to return before paginating.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCasesPaginator is a paginator for DescribeCases
type DescribeCasesPaginator struct {
	options   DescribeCasesPaginatorOptions
	client    DescribeCasesAPIClient
	params    *DescribeCasesInput
	nextToken *string
	firstPage bool
}

// NewDescribeCasesPaginator returns a new DescribeCasesPaginator
func NewDescribeCasesPaginator(client DescribeCasesAPIClient, params *DescribeCasesInput, optFns ...func(*DescribeCasesPaginatorOptions)) *DescribeCasesPaginator {
	if params == nil {
		params = &DescribeCasesInput{}
	}

	options := DescribeCasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCasesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeCases page.
func (p *DescribeCasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCasesOutput, error) {
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

	result, err := p.client.DescribeCases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeCases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeCases",
	}
}
