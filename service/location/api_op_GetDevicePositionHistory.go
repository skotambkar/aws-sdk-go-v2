// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the device position history from a tracker resource within a specified
// range of time. Device positions are deleted after 1 year.
func (c *Client) GetDevicePositionHistory(ctx context.Context, params *GetDevicePositionHistoryInput, optFns ...func(*Options)) (*GetDevicePositionHistoryOutput, error) {
	if params == nil {
		params = &GetDevicePositionHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevicePositionHistory", params, optFns, c.addOperationGetDevicePositionHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDevicePositionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDevicePositionHistoryInput struct {

	// The device whose position history you want to retrieve.
	//
	// This member is required.
	DeviceId *string

	// The tracker resource receiving the request for the device position history.
	//
	// This member is required.
	TrackerName *string

	// Specify the end time for the position history in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ. By default, the value will be the time that the
	// request is made. Requirement:
	//
	// * The time specified for EndTimeExclusive must be
	// after the time for StartTimeInclusive.
	EndTimeExclusive *time.Time

	// The pagination token specifying which page of results to return in the response.
	// If no token is provided, the default page is the first page. Default value: null
	NextToken *string

	// Specify the start time for the position history in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ. By default, the value will be 24 hours prior to the
	// time that the request is made. Requirement:
	//
	// * The time specified for
	// StartTimeInclusive must be before EndTimeExclusive.
	StartTimeInclusive *time.Time
}

type GetDevicePositionHistoryOutput struct {

	// Contains the position history details for the requested device.
	//
	// This member is required.
	DevicePositions []types.DevicePosition

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetDevicePositionHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDevicePositionHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDevicePositionHistory{}, middleware.After)
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
	if err = addOpGetDevicePositionHistoryValidationMiddleware(stack); err != nil {
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

// GetDevicePositionHistoryAPIClient is a client that implements the
// GetDevicePositionHistory operation.
type GetDevicePositionHistoryAPIClient interface {
	GetDevicePositionHistory(context.Context, *GetDevicePositionHistoryInput, ...func(*Options)) (*GetDevicePositionHistoryOutput, error)
}

var _ GetDevicePositionHistoryAPIClient = (*Client)(nil)

// GetDevicePositionHistoryPaginatorOptions is the paginator options for
// GetDevicePositionHistory
type GetDevicePositionHistoryPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDevicePositionHistoryPaginator is a paginator for GetDevicePositionHistory
type GetDevicePositionHistoryPaginator struct {
	options   GetDevicePositionHistoryPaginatorOptions
	client    GetDevicePositionHistoryAPIClient
	params    *GetDevicePositionHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetDevicePositionHistoryPaginator returns a new
// GetDevicePositionHistoryPaginator
func NewGetDevicePositionHistoryPaginator(client GetDevicePositionHistoryAPIClient, params *GetDevicePositionHistoryInput, optFns ...func(*GetDevicePositionHistoryPaginatorOptions)) *GetDevicePositionHistoryPaginator {
	if params == nil {
		params = &GetDevicePositionHistoryInput{}
	}

	options := GetDevicePositionHistoryPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetDevicePositionHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDevicePositionHistoryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetDevicePositionHistory page.
func (p *GetDevicePositionHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDevicePositionHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetDevicePositionHistory(ctx, &params, optFns...)
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
