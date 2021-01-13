// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of AWS IoT Greengrass core devices.
func (c *Client) ListCoreDevices(ctx context.Context, params *ListCoreDevicesInput, optFns ...func(*Options)) (*ListCoreDevicesOutput, error) {
	if params == nil {
		params = &ListCoreDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCoreDevices", params, optFns, addOperationListCoreDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCoreDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCoreDevicesInput struct {

	// The maximum number of results to be returned per paginated request.
	MaxResults int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The core device status by which to filter. If you specify this parameter, the
	// list includes only core devices that have this status. Choose one of the
	// following options:
	//
	// * HEALTHY – The AWS IoT Greengrass Core software and all
	// components run on the core device without issue.
	//
	// * UNHEALTHY – The AWS IoT
	// Greengrass Core software or a component is in a failed state on the core device.
	Status types.CoreDeviceStatus

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the AWS IoT thing group by which to filter. If you specify this parameter, the
	// list includes only core devices that are members of this thing group.
	ThingGroupArn *string
}

type ListCoreDevicesOutput struct {

	// A list that summarizes each core device.
	CoreDevices []types.CoreDevice

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCoreDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCoreDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCoreDevices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCoreDevices(options.Region), middleware.Before); err != nil {
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

// ListCoreDevicesAPIClient is a client that implements the ListCoreDevices
// operation.
type ListCoreDevicesAPIClient interface {
	ListCoreDevices(context.Context, *ListCoreDevicesInput, ...func(*Options)) (*ListCoreDevicesOutput, error)
}

var _ ListCoreDevicesAPIClient = (*Client)(nil)

// ListCoreDevicesPaginatorOptions is the paginator options for ListCoreDevices
type ListCoreDevicesPaginatorOptions struct {
	// The maximum number of results to be returned per paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCoreDevicesPaginator is a paginator for ListCoreDevices
type ListCoreDevicesPaginator struct {
	options   ListCoreDevicesPaginatorOptions
	client    ListCoreDevicesAPIClient
	params    *ListCoreDevicesInput
	nextToken *string
	firstPage bool
}

// NewListCoreDevicesPaginator returns a new ListCoreDevicesPaginator
func NewListCoreDevicesPaginator(client ListCoreDevicesAPIClient, params *ListCoreDevicesInput, optFns ...func(*ListCoreDevicesPaginatorOptions)) *ListCoreDevicesPaginator {
	options := ListCoreDevicesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListCoreDevicesInput{}
	}

	return &ListCoreDevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCoreDevicesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCoreDevices page.
func (p *ListCoreDevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCoreDevicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListCoreDevices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCoreDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListCoreDevices",
	}
}
