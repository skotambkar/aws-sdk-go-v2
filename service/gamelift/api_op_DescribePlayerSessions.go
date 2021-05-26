// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves properties for one or more player sessions. This action can be used in
// the following ways:
//
// * To retrieve a specific player session, provide the player
// session ID only.
//
// * To retrieve all player sessions in a game session, provide
// the game session ID only.
//
// * To retrieve all player sessions for a specific
// player, provide a player ID only.
//
// To request player sessions, specify either a
// player session ID, game session ID, or player ID. You can filter this request by
// player session status. Use the pagination parameters to retrieve results as a
// set of sequential pages. If successful, a PlayerSession object is returned for
// each session that matches the request. Available in Amazon GameLift Local.
// Related actions CreatePlayerSession | CreatePlayerSessions |
// DescribePlayerSessions | StartGameSessionPlacement |
// DescribeGameSessionPlacement | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribePlayerSessions(ctx context.Context, params *DescribePlayerSessionsInput, optFns ...func(*Options)) (*DescribePlayerSessionsOutput, error) {
	if params == nil {
		params = &DescribePlayerSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePlayerSessions", params, optFns, c.addOperationDescribePlayerSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePlayerSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribePlayerSessionsInput struct {

	// A unique identifier for the game session to retrieve player sessions for.
	GameSessionId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. If a player session ID is specified,
	// this parameter is ignored.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value. If a player session ID is
	// specified, this parameter is ignored.
	NextToken *string

	// A unique identifier for a player to retrieve player sessions for.
	PlayerId *string

	// A unique identifier for a player session to retrieve.
	PlayerSessionId *string

	// Player session status to filter results on. Possible player session statuses
	// include the following:
	//
	// * RESERVED -- The player session request has been
	// received, but the player has not yet connected to the server process and/or been
	// validated.
	//
	// * ACTIVE -- The player has been validated by the server process and
	// is currently connected.
	//
	// * COMPLETED -- The player connection has been
	// dropped.
	//
	// * TIMEDOUT -- A player session request was received, but the player
	// did not connect and/or was not validated within the timeout limit (60 seconds).
	PlayerSessionStatusFilter *string
}

// Represents the returned data in response to a request operation.
type DescribePlayerSessionsOutput struct {

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// A collection of objects containing properties for each player session that
	// matches the request.
	PlayerSessions []types.PlayerSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribePlayerSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePlayerSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePlayerSessions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePlayerSessions(options.Region), middleware.Before); err != nil {
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

// DescribePlayerSessionsAPIClient is a client that implements the
// DescribePlayerSessions operation.
type DescribePlayerSessionsAPIClient interface {
	DescribePlayerSessions(context.Context, *DescribePlayerSessionsInput, ...func(*Options)) (*DescribePlayerSessionsOutput, error)
}

var _ DescribePlayerSessionsAPIClient = (*Client)(nil)

// DescribePlayerSessionsPaginatorOptions is the paginator options for
// DescribePlayerSessions
type DescribePlayerSessionsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. If a player session ID is specified,
	// this parameter is ignored.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePlayerSessionsPaginator is a paginator for DescribePlayerSessions
type DescribePlayerSessionsPaginator struct {
	options   DescribePlayerSessionsPaginatorOptions
	client    DescribePlayerSessionsAPIClient
	params    *DescribePlayerSessionsInput
	nextToken *string
	firstPage bool
}

// NewDescribePlayerSessionsPaginator returns a new DescribePlayerSessionsPaginator
func NewDescribePlayerSessionsPaginator(client DescribePlayerSessionsAPIClient, params *DescribePlayerSessionsInput, optFns ...func(*DescribePlayerSessionsPaginatorOptions)) *DescribePlayerSessionsPaginator {
	if params == nil {
		params = &DescribePlayerSessionsInput{}
	}

	options := DescribePlayerSessionsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePlayerSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePlayerSessionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribePlayerSessions page.
func (p *DescribePlayerSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePlayerSessionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribePlayerSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePlayerSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribePlayerSessions",
	}
}
