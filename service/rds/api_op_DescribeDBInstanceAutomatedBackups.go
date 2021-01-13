// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays backups for both current and deleted instances. For example, use this
// operation to find details about automated backups for previously deleted
// instances. Current instances with retention periods greater than zero (0) are
// returned for both the DescribeDBInstanceAutomatedBackups and DescribeDBInstances
// operations. All parameters are optional.
func (c *Client) DescribeDBInstanceAutomatedBackups(ctx context.Context, params *DescribeDBInstanceAutomatedBackupsInput, optFns ...func(*Options)) (*DescribeDBInstanceAutomatedBackupsOutput, error) {
	if params == nil {
		params = &DescribeDBInstanceAutomatedBackupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBInstanceAutomatedBackups", params, optFns, addOperationDescribeDBInstanceAutomatedBackupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBInstanceAutomatedBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Parameter input for DescribeDBInstanceAutomatedBackups.
type DescribeDBInstanceAutomatedBackupsInput struct {

	// The Amazon Resource Name (ARN) of the replicated automated backups, for example,
	// arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE.
	DBInstanceAutomatedBackupsArn *string

	// (Optional) The user-supplied instance identifier. If this parameter is
	// specified, it must match the identifier of an existing DB instance. It returns
	// information from the specific DB instance' automated backup. This parameter
	// isn't case-sensitive.
	DBInstanceIdentifier *string

	// The resource ID of the DB instance that is the source of the automated backup.
	// This parameter isn't case-sensitive.
	DbiResourceId *string

	// A filter that specifies which resources to return based on status. Supported
	// filters are the following:
	//
	// * status
	//
	// * active - automated backups for current
	// instances
	//
	// * retained - automated backups for deleted instances and after backup
	// replication is stopped
	//
	// * creating - automated backups that are waiting for the
	// first automated snapshot to be available
	//
	// * db-instance-id - Accepts DB instance
	// identifiers and Amazon Resource Names (ARNs). The results list includes only
	// information about the DB instance automated backups identified by these ARNs.
	//
	// *
	// dbi-resource-id - Accepts DB resource identifiers and Amazon Resource Names
	// (ARNs). The results list includes only information about the DB instance
	// resources identified by these ARNs.
	//
	// Returns all resources by default. The
	// status for each resource is specified in the response.
	Filters []types.Filter

	// The pagination token provided in the previous request. If this parameter is
	// specified the response includes only records beyond the marker, up to
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the
// DescribeDBInstanceAutomatedBackups action.
type DescribeDBInstanceAutomatedBackupsOutput struct {

	// A list of DBInstanceAutomatedBackup instances.
	DBInstanceAutomatedBackups []types.DBInstanceAutomatedBackup

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBInstanceAutomatedBackupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBInstanceAutomatedBackups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBInstanceAutomatedBackups{}, middleware.After)
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
	if err = addOpDescribeDBInstanceAutomatedBackupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBInstanceAutomatedBackups(options.Region), middleware.Before); err != nil {
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

// DescribeDBInstanceAutomatedBackupsAPIClient is a client that implements the
// DescribeDBInstanceAutomatedBackups operation.
type DescribeDBInstanceAutomatedBackupsAPIClient interface {
	DescribeDBInstanceAutomatedBackups(context.Context, *DescribeDBInstanceAutomatedBackupsInput, ...func(*Options)) (*DescribeDBInstanceAutomatedBackupsOutput, error)
}

var _ DescribeDBInstanceAutomatedBackupsAPIClient = (*Client)(nil)

// DescribeDBInstanceAutomatedBackupsPaginatorOptions is the paginator options for
// DescribeDBInstanceAutomatedBackups
type DescribeDBInstanceAutomatedBackupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBInstanceAutomatedBackupsPaginator is a paginator for
// DescribeDBInstanceAutomatedBackups
type DescribeDBInstanceAutomatedBackupsPaginator struct {
	options   DescribeDBInstanceAutomatedBackupsPaginatorOptions
	client    DescribeDBInstanceAutomatedBackupsAPIClient
	params    *DescribeDBInstanceAutomatedBackupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBInstanceAutomatedBackupsPaginator returns a new
// DescribeDBInstanceAutomatedBackupsPaginator
func NewDescribeDBInstanceAutomatedBackupsPaginator(client DescribeDBInstanceAutomatedBackupsAPIClient, params *DescribeDBInstanceAutomatedBackupsInput, optFns ...func(*DescribeDBInstanceAutomatedBackupsPaginatorOptions)) *DescribeDBInstanceAutomatedBackupsPaginator {
	options := DescribeDBInstanceAutomatedBackupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeDBInstanceAutomatedBackupsInput{}
	}

	return &DescribeDBInstanceAutomatedBackupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBInstanceAutomatedBackupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeDBInstanceAutomatedBackups page.
func (p *DescribeDBInstanceAutomatedBackupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBInstanceAutomatedBackupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBInstanceAutomatedBackups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBInstanceAutomatedBackups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBInstanceAutomatedBackups",
	}
}
