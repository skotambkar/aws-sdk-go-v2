// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of table names associated with the current account and
// endpoint. The output from ListTables is paginated, with each page returning a
// maximum of 100 table names.
func (c *Client) ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) {
	if params == nil {
		params = &ListTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTables", params, optFns, c.addOperationListTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListTables operation.
type ListTablesInput struct {

	// The first table name that this operation will evaluate. Use the value that was
	// returned for LastEvaluatedTableName in a previous operation, so that you can
	// obtain the next page of results.
	ExclusiveStartTableName *string

	// A maximum number of table names to return. If this parameter is not specified,
	// the limit is 100.
	Limit *int32
}

// Represents the output of a ListTables operation.
type ListTablesOutput struct {

	// The name of the last table in the current page of results. Use this value as the
	// ExclusiveStartTableName in a new request to obtain the next page of results,
	// until all the table names are returned. If you do not receive a
	// LastEvaluatedTableName value in the response, this means that there are no more
	// table names to be retrieved.
	LastEvaluatedTableName *string

	// The names of the tables associated with the current account at the current
	// endpoint. The maximum size of this array is 100. If LastEvaluatedTableName also
	// appears in the output, you can use this value as the ExclusiveStartTableName
	// parameter in a subsequent ListTables request and obtain the next page of
	// results.
	TableNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListTables{}, middleware.After)
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
	if err = addOpListTablesDiscoverEndpointMiddleware(stack, options, c); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func addOpListTablesDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpListTablesDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpListTablesDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*ListTablesInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

// ListTablesAPIClient is a client that implements the ListTables operation.
type ListTablesAPIClient interface {
	ListTables(context.Context, *ListTablesInput, ...func(*Options)) (*ListTablesOutput, error)
}

var _ ListTablesAPIClient = (*Client)(nil)

// ListTablesPaginatorOptions is the paginator options for ListTables
type ListTablesPaginatorOptions struct {
	// A maximum number of table names to return. If this parameter is not specified,
	// the limit is 100.
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
	if params == nil {
		params = &ListTablesInput{}
	}

	options := ListTablesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
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
	params.ExclusiveStartTableName = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListTables(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.LastEvaluatedTableName

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListTables",
	}
}
