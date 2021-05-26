// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all current job definitions.
func (c *Client) GetJobs(ctx context.Context, params *GetJobsInput, optFns ...func(*Options)) (*GetJobsOutput, error) {
	if params == nil {
		params = &GetJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJobs", params, optFns, c.addOperationGetJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobsInput struct {

	// The maximum size of the response.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type GetJobsOutput struct {

	// A list of job definitions.
	Jobs []types.Job

	// A continuation token, if not all job definitions have yet been returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobs(options.Region), middleware.Before); err != nil {
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

// GetJobsAPIClient is a client that implements the GetJobs operation.
type GetJobsAPIClient interface {
	GetJobs(context.Context, *GetJobsInput, ...func(*Options)) (*GetJobsOutput, error)
}

var _ GetJobsAPIClient = (*Client)(nil)

// GetJobsPaginatorOptions is the paginator options for GetJobs
type GetJobsPaginatorOptions struct {
	// The maximum size of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetJobsPaginator is a paginator for GetJobs
type GetJobsPaginator struct {
	options   GetJobsPaginatorOptions
	client    GetJobsAPIClient
	params    *GetJobsInput
	nextToken *string
	firstPage bool
}

// NewGetJobsPaginator returns a new GetJobsPaginator
func NewGetJobsPaginator(client GetJobsAPIClient, params *GetJobsInput, optFns ...func(*GetJobsPaginatorOptions)) *GetJobsPaginator {
	if params == nil {
		params = &GetJobsInput{}
	}

	options := GetJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetJobs page.
func (p *GetJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetJobsOutput, error) {
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

	result, err := p.client.GetJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetJobs",
	}
}
