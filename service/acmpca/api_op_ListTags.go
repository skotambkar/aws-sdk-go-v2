// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the tags, if any, that are associated with your private CA or one that has
// been shared with you. Tags are labels that you can use to identify and organize
// your CAs. Each tag consists of a key and an optional value. Call the
// TagCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_TagCertificateAuthority.html)
// action to add one or more tags to your CA. Call the UntagCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_UntagCertificateAuthority.html)
// action to remove tags.
func (c *Client) ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) {
	if params == nil {
		params = &ListTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTags", params, optFns, addOperationListTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called the
	// CreateCertificateAuthority
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html)
	// action. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken value
	// in a subsequent request to retrieve additional items.
	MaxResults *int32

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of NextToken from
	// the response you just received.
	NextToken *string
}

type ListTagsOutput struct {

	// When the list is truncated, this value is present and should be used for the
	// NextToken parameter in a subsequent pagination request.
	NextToken *string

	// The tags associated with your private CA.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTags{}, middleware.After)
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
	if err = addOpListTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTags(options.Region), middleware.Before); err != nil {
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

// ListTagsAPIClient is a client that implements the ListTags operation.
type ListTagsAPIClient interface {
	ListTags(context.Context, *ListTagsInput, ...func(*Options)) (*ListTagsOutput, error)
}

var _ ListTagsAPIClient = (*Client)(nil)

// ListTagsPaginatorOptions is the paginator options for ListTags
type ListTagsPaginatorOptions struct {
	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken value
	// in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTagsPaginator is a paginator for ListTags
type ListTagsPaginator struct {
	options   ListTagsPaginatorOptions
	client    ListTagsAPIClient
	params    *ListTagsInput
	nextToken *string
	firstPage bool
}

// NewListTagsPaginator returns a new ListTagsPaginator
func NewListTagsPaginator(client ListTagsAPIClient, params *ListTagsInput, optFns ...func(*ListTagsPaginatorOptions)) *ListTagsPaginator {
	options := ListTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTagsInput{}
	}

	return &ListTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTagsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTags page.
func (p *ListTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTagsOutput, error) {
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

	result, err := p.client.ListTags(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "ListTags",
	}
}
