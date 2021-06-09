// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the details of one or more configuration aggregators. If the
// configuration aggregator is not specified, this action returns the details for
// all the configuration aggregators associated with the account.
func (c *Client) DescribeConfigurationAggregators(ctx context.Context, params *DescribeConfigurationAggregatorsInput, optFns ...func(*Options)) (*DescribeConfigurationAggregatorsOutput, error) {
	if params == nil {
		params = &DescribeConfigurationAggregatorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationAggregators", params, optFns, c.addOperationDescribeConfigurationAggregatorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationAggregatorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationAggregatorsInput struct {

	// The name of the configuration aggregators.
	ConfigurationAggregatorNames []string

	// The maximum number of configuration aggregators returned on each page. The
	// default is maximum. If you specify 0, AWS Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type DescribeConfigurationAggregatorsOutput struct {

	// Returns a ConfigurationAggregators object.
	ConfigurationAggregators []types.ConfigurationAggregator

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeConfigurationAggregatorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationAggregators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationAggregators{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationAggregators(options.Region), middleware.Before); err != nil {
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

// DescribeConfigurationAggregatorsAPIClient is a client that implements the
// DescribeConfigurationAggregators operation.
type DescribeConfigurationAggregatorsAPIClient interface {
	DescribeConfigurationAggregators(context.Context, *DescribeConfigurationAggregatorsInput, ...func(*Options)) (*DescribeConfigurationAggregatorsOutput, error)
}

var _ DescribeConfigurationAggregatorsAPIClient = (*Client)(nil)

// DescribeConfigurationAggregatorsPaginatorOptions is the paginator options for
// DescribeConfigurationAggregators
type DescribeConfigurationAggregatorsPaginatorOptions struct {
	// The maximum number of configuration aggregators returned on each page. The
	// default is maximum. If you specify 0, AWS Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConfigurationAggregatorsPaginator is a paginator for
// DescribeConfigurationAggregators
type DescribeConfigurationAggregatorsPaginator struct {
	options   DescribeConfigurationAggregatorsPaginatorOptions
	client    DescribeConfigurationAggregatorsAPIClient
	params    *DescribeConfigurationAggregatorsInput
	nextToken *string
	firstPage bool
}

// NewDescribeConfigurationAggregatorsPaginator returns a new
// DescribeConfigurationAggregatorsPaginator
func NewDescribeConfigurationAggregatorsPaginator(client DescribeConfigurationAggregatorsAPIClient, params *DescribeConfigurationAggregatorsInput, optFns ...func(*DescribeConfigurationAggregatorsPaginatorOptions)) *DescribeConfigurationAggregatorsPaginator {
	if params == nil {
		params = &DescribeConfigurationAggregatorsInput{}
	}

	options := DescribeConfigurationAggregatorsPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConfigurationAggregatorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConfigurationAggregatorsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeConfigurationAggregators page.
func (p *DescribeConfigurationAggregatorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConfigurationAggregatorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeConfigurationAggregators(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConfigurationAggregators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigurationAggregators",
	}
}
