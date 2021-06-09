// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all of the variables or the specific variable. This is a paginated API.
// Providing null maxSizePerPage results in retrieving maximum of 100 records per
// page. If you provide maxSizePerPage the value must be between 50 and 100. To get
// the next page result, a provide a pagination token from GetVariablesResult as
// part of your request. Null pagination token fetches the records from the
// beginning.
func (c *Client) GetVariables(ctx context.Context, params *GetVariablesInput, optFns ...func(*Options)) (*GetVariablesOutput, error) {
	if params == nil {
		params = &GetVariablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVariables", params, optFns, c.addOperationGetVariablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVariablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVariablesInput struct {

	// The max size per page determined for the get variable request.
	MaxResults *int32

	// The name of the variable.
	Name *string

	// The next page token of the get variable request.
	NextToken *string
}

type GetVariablesOutput struct {

	// The next page token to be used in subsequent requests.
	NextToken *string

	// The names of the variables returned.
	Variables []types.Variable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetVariablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetVariables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetVariables{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVariables(options.Region), middleware.Before); err != nil {
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

// GetVariablesAPIClient is a client that implements the GetVariables operation.
type GetVariablesAPIClient interface {
	GetVariables(context.Context, *GetVariablesInput, ...func(*Options)) (*GetVariablesOutput, error)
}

var _ GetVariablesAPIClient = (*Client)(nil)

// GetVariablesPaginatorOptions is the paginator options for GetVariables
type GetVariablesPaginatorOptions struct {
	// The max size per page determined for the get variable request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetVariablesPaginator is a paginator for GetVariables
type GetVariablesPaginator struct {
	options   GetVariablesPaginatorOptions
	client    GetVariablesAPIClient
	params    *GetVariablesInput
	nextToken *string
	firstPage bool
}

// NewGetVariablesPaginator returns a new GetVariablesPaginator
func NewGetVariablesPaginator(client GetVariablesAPIClient, params *GetVariablesInput, optFns ...func(*GetVariablesPaginatorOptions)) *GetVariablesPaginator {
	if params == nil {
		params = &GetVariablesInput{}
	}

	options := GetVariablesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetVariablesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetVariablesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetVariables page.
func (p *GetVariablesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetVariablesOutput, error) {
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

	result, err := p.client.GetVariables(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetVariables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetVariables",
	}
}
