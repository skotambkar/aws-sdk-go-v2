// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the assignee of the specified AWS account with the specified permission
// set.
func (c *Client) ListAccountAssignments(ctx context.Context, params *ListAccountAssignmentsInput, optFns ...func(*Options)) (*ListAccountAssignmentsOutput, error) {
	if params == nil {
		params = &ListAccountAssignmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountAssignments", params, optFns, c.addOperationListAccountAssignmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountAssignmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountAssignmentsInput struct {

	// The identifier of the AWS account from which to list the assignments.
	//
	// This member is required.
	AccountId *string

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set from which to list assignments.
	//
	// This member is required.
	PermissionSetArn *string

	// The maximum number of results to display for the assignment.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string
}

type ListAccountAssignmentsOutput struct {

	// The list of assignments that match the input AWS account and permission set.
	AccountAssignments []types.AccountAssignment

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListAccountAssignmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccountAssignments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccountAssignments{}, middleware.After)
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
	if err = addOpListAccountAssignmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountAssignments(options.Region), middleware.Before); err != nil {
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

// ListAccountAssignmentsAPIClient is a client that implements the
// ListAccountAssignments operation.
type ListAccountAssignmentsAPIClient interface {
	ListAccountAssignments(context.Context, *ListAccountAssignmentsInput, ...func(*Options)) (*ListAccountAssignmentsOutput, error)
}

var _ ListAccountAssignmentsAPIClient = (*Client)(nil)

// ListAccountAssignmentsPaginatorOptions is the paginator options for
// ListAccountAssignments
type ListAccountAssignmentsPaginatorOptions struct {
	// The maximum number of results to display for the assignment.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountAssignmentsPaginator is a paginator for ListAccountAssignments
type ListAccountAssignmentsPaginator struct {
	options   ListAccountAssignmentsPaginatorOptions
	client    ListAccountAssignmentsAPIClient
	params    *ListAccountAssignmentsInput
	nextToken *string
	firstPage bool
}

// NewListAccountAssignmentsPaginator returns a new ListAccountAssignmentsPaginator
func NewListAccountAssignmentsPaginator(client ListAccountAssignmentsAPIClient, params *ListAccountAssignmentsInput, optFns ...func(*ListAccountAssignmentsPaginatorOptions)) *ListAccountAssignmentsPaginator {
	if params == nil {
		params = &ListAccountAssignmentsInput{}
	}

	options := ListAccountAssignmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountAssignmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountAssignmentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAccountAssignments page.
func (p *ListAccountAssignmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountAssignmentsOutput, error) {
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

	result, err := p.client.ListAccountAssignments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccountAssignments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "ListAccountAssignments",
	}
}
