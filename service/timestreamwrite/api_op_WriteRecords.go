// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The WriteRecords operation enables you to write your time series data into
// Timestream. You can specify a single data point or a batch of data points to be
// inserted into the system. Timestream offers you with a flexible schema that auto
// detects the column names and data types for your Timestream tables based on the
// dimension names and data types of the data points you specify when invoking
// writes into the database. Timestream support eventual consistency read
// semantics. This means that when you query data immediately after writing a batch
// of data into Timestream, the query results might not reflect the results of a
// recently completed write operation. The results may also include some stale
// data. If you repeat the query request after a short time, the results should
// return the latest data. Service quotas apply. For more information, see Access
// Management
// (https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html) in
// the Timestream Developer Guide.
func (c *Client) WriteRecords(ctx context.Context, params *WriteRecordsInput, optFns ...func(*Options)) (*WriteRecordsOutput, error) {
	if params == nil {
		params = &WriteRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "WriteRecords", params, optFns, addOperationWriteRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*WriteRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type WriteRecordsInput struct {

	// The name of the Timestream database.
	//
	// This member is required.
	DatabaseName *string

	// An array of records containing the unique dimension and measure attributes for
	// each time series data point.
	//
	// This member is required.
	Records []types.Record

	// The name of the Timesream table.
	//
	// This member is required.
	TableName *string

	// A record containing the common measure and dimension attributes shared across
	// all the records in the request. The measure and dimension attributes specified
	// in here will be merged with the measure and dimension attributes in the records
	// object when the data is written into Timestream.
	CommonAttributes *types.Record
}

type WriteRecordsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationWriteRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpWriteRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpWriteRecords{}, middleware.After)
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
	if err = addOpWriteRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opWriteRecords(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opWriteRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "WriteRecords",
	}
}
