// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of contacts. If statusList contains AVAILABLE, the request must
// include groundStation, missionprofileArn, and satelliteArn.
func (c *Client) ListContacts(ctx context.Context, params *ListContactsInput, optFns ...func(*Options)) (*ListContactsOutput, error) {
	if params == nil {
		params = &ListContactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContacts", params, optFns, c.addOperationListContactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListContactsInput struct {

	// End time of a contact.
	//
	// This member is required.
	EndTime *time.Time

	// Start time of a contact.
	//
	// This member is required.
	StartTime *time.Time

	// Status of a contact reservation.
	//
	// This member is required.
	StatusList []types.ContactStatus

	// Name of a ground station.
	GroundStation *string

	// Maximum number of contacts returned.
	MaxResults *int32

	// ARN of a mission profile.
	MissionProfileArn *string

	// Next token returned in the request of a previous ListContacts call. Used to get
	// the next page of results.
	NextToken *string

	// ARN of a satellite.
	SatelliteArn *string
}

//
type ListContactsOutput struct {

	// List of contacts.
	ContactList []types.ContactData

	// Next token returned in the response of a previous ListContacts call. Used to get
	// the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListContactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContacts{}, middleware.After)
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
	if err = addOpListContactsValidationMiddleware(stack); err != nil {
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

// ListContactsAPIClient is a client that implements the ListContacts operation.
type ListContactsAPIClient interface {
	ListContacts(context.Context, *ListContactsInput, ...func(*Options)) (*ListContactsOutput, error)
}

var _ ListContactsAPIClient = (*Client)(nil)

// ListContactsPaginatorOptions is the paginator options for ListContacts
type ListContactsPaginatorOptions struct {
	// Maximum number of contacts returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContactsPaginator is a paginator for ListContacts
type ListContactsPaginator struct {
	options   ListContactsPaginatorOptions
	client    ListContactsAPIClient
	params    *ListContactsInput
	nextToken *string
	firstPage bool
}

// NewListContactsPaginator returns a new ListContactsPaginator
func NewListContactsPaginator(client ListContactsAPIClient, params *ListContactsInput, optFns ...func(*ListContactsPaginatorOptions)) *ListContactsPaginator {
	if params == nil {
		params = &ListContactsInput{}
	}

	options := ListContactsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContactsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListContacts page.
func (p *ListContactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContactsOutput, error) {
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

	result, err := p.client.ListContacts(ctx, &params, optFns...)
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
