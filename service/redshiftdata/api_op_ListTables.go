// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the tables in a database. If neither SchemaPattern nor TablePattern are
// specified, then all tables in the database are returned. A token is returned to
// page through the table list. Depending on the authorization method, use one of
// the following combinations of request parameters:
//
// * AWS Secrets Manager -
// specify the Amazon Resource Name (ARN) of the secret and the cluster identifier
// that matches the cluster in the secret.
//
// * Temporary credentials - specify the
// cluster identifier, the database name, and the database user name. Permission to
// call the redshift:GetClusterCredentials operation is required to use this
// method.
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

type ListTablesInput struct {

	// The cluster identifier. This parameter is required when authenticating using
	// either AWS Secrets Manager or temporary credentials.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the database that contains the tables to list. If ConnectedDatabase
	// is not specified, this is also the database to connect to with your
	// authentication credentials.
	//
	// This member is required.
	Database *string

	// A database name. The connected database is specified when you connect with your
	// authentication credentials.
	ConnectedDatabase *string

	// The database user name. This parameter is required when authenticating using
	// temporary credentials.
	DbUser *string

	// The maximum number of tables to return in the response. If more tables exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
	MaxResults int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// A pattern to filter results by schema name. Within a schema pattern, "%" means
	// match any substring of 0 or more characters and "_" means match any one
	// character. Only schema name entries matching the search pattern are returned. If
	// SchemaPattern is not specified, then all tables that match TablePattern are
	// returned. If neither SchemaPattern or TablePattern are specified, then all
	// tables are returned.
	SchemaPattern *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using AWS Secrets Manager.
	SecretArn *string

	// A pattern to filter results by table name. Within a table pattern, "%" means
	// match any substring of 0 or more characters and "_" means match any one
	// character. Only table name entries matching the search pattern are returned. If
	// TablePattern is not specified, then all tables that match SchemaPatternare
	// returned. If neither SchemaPattern or TablePattern are specified, then all
	// tables are returned.
	TablePattern *string
}

type ListTablesOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The tables that match the request pattern.
	Tables []types.TableMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTables{}, middleware.After)
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
	if err = addOpListTablesValidationMiddleware(stack); err != nil {
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
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListTablesAPIClient is a client that implements the ListTables operation.
type ListTablesAPIClient interface {
	ListTables(context.Context, *ListTablesInput, ...func(*Options)) (*ListTablesOutput, error)
}

var _ ListTablesAPIClient = (*Client)(nil)

// ListTablesPaginatorOptions is the paginator options for ListTables
type ListTablesPaginatorOptions struct {
	// The maximum number of tables to return in the response. If more tables exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
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
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
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
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListTables(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "ListTables",
	}
}
