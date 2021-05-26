// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API takes either a ProvisonedProductId or a ProvisionedProductName, along
// with a list of one or more output keys, and responds with the key/value pairs of
// those outputs.
func (c *Client) GetProvisionedProductOutputs(ctx context.Context, params *GetProvisionedProductOutputsInput, optFns ...func(*Options)) (*GetProvisionedProductOutputsOutput, error) {
	if params == nil {
		params = &GetProvisionedProductOutputsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProvisionedProductOutputs", params, optFns, c.addOperationGetProvisionedProductOutputsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProvisionedProductOutputsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProvisionedProductOutputsInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The list of keys that the API should return with their values. If none are
	// provided, the API will return all outputs of the provisioned product.
	OutputKeys []string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The identifier of the provisioned product that you want the outputs from.
	ProvisionedProductId *string

	// The name of the provisioned product that you want the outputs from.
	ProvisionedProductName *string
}

type GetProvisionedProductOutputsOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Information about the product created as the result of a request. For example,
	// the output for a CloudFormation-backed product that creates an S3 bucket would
	// include the S3 bucket URL.
	Outputs []types.RecordOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetProvisionedProductOutputsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetProvisionedProductOutputs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetProvisionedProductOutputs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProvisionedProductOutputs(options.Region), middleware.Before); err != nil {
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

// GetProvisionedProductOutputsAPIClient is a client that implements the
// GetProvisionedProductOutputs operation.
type GetProvisionedProductOutputsAPIClient interface {
	GetProvisionedProductOutputs(context.Context, *GetProvisionedProductOutputsInput, ...func(*Options)) (*GetProvisionedProductOutputsOutput, error)
}

var _ GetProvisionedProductOutputsAPIClient = (*Client)(nil)

// GetProvisionedProductOutputsPaginatorOptions is the paginator options for
// GetProvisionedProductOutputs
type GetProvisionedProductOutputsPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetProvisionedProductOutputsPaginator is a paginator for
// GetProvisionedProductOutputs
type GetProvisionedProductOutputsPaginator struct {
	options   GetProvisionedProductOutputsPaginatorOptions
	client    GetProvisionedProductOutputsAPIClient
	params    *GetProvisionedProductOutputsInput
	nextToken *string
	firstPage bool
}

// NewGetProvisionedProductOutputsPaginator returns a new
// GetProvisionedProductOutputsPaginator
func NewGetProvisionedProductOutputsPaginator(client GetProvisionedProductOutputsAPIClient, params *GetProvisionedProductOutputsInput, optFns ...func(*GetProvisionedProductOutputsPaginatorOptions)) *GetProvisionedProductOutputsPaginator {
	if params == nil {
		params = &GetProvisionedProductOutputsInput{}
	}

	options := GetProvisionedProductOutputsPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetProvisionedProductOutputsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetProvisionedProductOutputsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetProvisionedProductOutputs page.
func (p *GetProvisionedProductOutputsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetProvisionedProductOutputsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.GetProvisionedProductOutputs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetProvisionedProductOutputs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "GetProvisionedProductOutputs",
	}
}
