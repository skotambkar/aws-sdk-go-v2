// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the AWS services that you enabled to integrate with your
// organization. After a service on this list creates the resources that it
// requires for the integration, it can perform operations on your organization and
// its accounts. For more information about integrating other services with AWS
// Organizations, including the list of services that currently work with
// Organizations, see Integrating AWS Organizations with Other AWS Services
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's management account or by a member account that is a delegated
// administrator for an AWS service.
func (c *Client) ListAWSServiceAccessForOrganization(ctx context.Context, params *ListAWSServiceAccessForOrganizationInput, optFns ...func(*Options)) (*ListAWSServiceAccessForOrganizationOutput, error) {
	if params == nil {
		params = &ListAWSServiceAccessForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAWSServiceAccessForOrganization", params, optFns, addOperationListAWSServiceAccessForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAWSServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAWSServiceAccessForOrganizationInput struct {

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListAWSServiceAccessForOrganizationOutput struct {

	// A list of the service principals for the services that are enabled to integrate
	// with your organization. Each principal is a structure that includes the name and
	// the date that it was enabled for integration with AWS Organizations.
	EnabledServicePrincipals []types.EnabledServicePrincipal

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAWSServiceAccessForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAWSServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAWSServiceAccessForOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAWSServiceAccessForOrganization(options.Region), middleware.Before); err != nil {
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

// ListAWSServiceAccessForOrganizationAPIClient is a client that implements the
// ListAWSServiceAccessForOrganization operation.
type ListAWSServiceAccessForOrganizationAPIClient interface {
	ListAWSServiceAccessForOrganization(context.Context, *ListAWSServiceAccessForOrganizationInput, ...func(*Options)) (*ListAWSServiceAccessForOrganizationOutput, error)
}

var _ ListAWSServiceAccessForOrganizationAPIClient = (*Client)(nil)

// ListAWSServiceAccessForOrganizationPaginatorOptions is the paginator options for
// ListAWSServiceAccessForOrganization
type ListAWSServiceAccessForOrganizationPaginatorOptions struct {
	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAWSServiceAccessForOrganizationPaginator is a paginator for
// ListAWSServiceAccessForOrganization
type ListAWSServiceAccessForOrganizationPaginator struct {
	options   ListAWSServiceAccessForOrganizationPaginatorOptions
	client    ListAWSServiceAccessForOrganizationAPIClient
	params    *ListAWSServiceAccessForOrganizationInput
	nextToken *string
	firstPage bool
}

// NewListAWSServiceAccessForOrganizationPaginator returns a new
// ListAWSServiceAccessForOrganizationPaginator
func NewListAWSServiceAccessForOrganizationPaginator(client ListAWSServiceAccessForOrganizationAPIClient, params *ListAWSServiceAccessForOrganizationInput, optFns ...func(*ListAWSServiceAccessForOrganizationPaginatorOptions)) *ListAWSServiceAccessForOrganizationPaginator {
	options := ListAWSServiceAccessForOrganizationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListAWSServiceAccessForOrganizationInput{}
	}

	return &ListAWSServiceAccessForOrganizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAWSServiceAccessForOrganizationPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAWSServiceAccessForOrganization page.
func (p *ListAWSServiceAccessForOrganizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAWSServiceAccessForOrganizationOutput, error) {
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

	result, err := p.client.ListAWSServiceAccessForOrganization(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAWSServiceAccessForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListAWSServiceAccessForOrganization",
	}
}
