// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/firehose/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

const opCreateDeliveryStream = "CreateDeliveryStream"

// CreateDeliveryStreamRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Creates a Kinesis Data Firehose delivery stream.
//
// By default, you can create up to 50 delivery streams per AWS Region.
//
// This is an asynchronous operation that immediately returns. The initial status
// of the delivery stream is CREATING. After the delivery stream is created,
// its status is ACTIVE and it now accepts data. If the delivery stream creation
// fails, the status transitions to CREATING_FAILED. Attempts to send data to
// a delivery stream that is not in the ACTIVE state cause an exception. To
// check the state of a delivery stream, use DescribeDeliveryStream.
//
// If the status of a delivery stream is CREATING_FAILED, this status doesn't
// change, and you can't invoke CreateDeliveryStream again on it. However, you
// can invoke the DeleteDeliveryStream operation to delete it.
//
// A Kinesis Data Firehose delivery stream can be configured to receive records
// directly from providers using PutRecord or PutRecordBatch, or it can be configured
// to use an existing Kinesis stream as its source. To specify a Kinesis data
// stream as input, set the DeliveryStreamType parameter to KinesisStreamAsSource,
// and provide the Kinesis stream Amazon Resource Name (ARN) and role ARN in
// the KinesisStreamSourceConfiguration parameter.
//
// To create a delivery stream with server-side encryption (SSE) enabled, include
// DeliveryStreamEncryptionConfigurationInput in your request. This is optional.
// You can also invoke StartDeliveryStreamEncryption to turn on SSE for an existing
// delivery stream that doesn't have SSE enabled.
//
// A delivery stream is configured with a single destination: Amazon S3, Amazon
// ES, Amazon Redshift, or Splunk. You must specify only one of the following
// destination configuration parameters: ExtendedS3DestinationConfiguration,
// S3DestinationConfiguration, ElasticsearchDestinationConfiguration, RedshiftDestinationConfiguration,
// or SplunkDestinationConfiguration.
//
// When you specify S3DestinationConfiguration, you can also provide the following
// optional values: BufferingHints, EncryptionConfiguration, and CompressionFormat.
// By default, if no BufferingHints value is provided, Kinesis Data Firehose
// buffers data up to 5 MB or for 5 minutes, whichever condition is satisfied
// first. BufferingHints is a hint, so there are some cases where the service
// cannot adhere to these conditions strictly. For example, record boundaries
// might be such that the size is a little over or under the configured buffering
// size. By default, no encryption is performed. We strongly recommend that
// you enable encryption to ensure secure data storage in Amazon S3.
//
// A few notes about Amazon Redshift as a destination:
//
//    * An Amazon Redshift destination requires an S3 bucket as intermediate
//    location. Kinesis Data Firehose first delivers data to Amazon S3 and then
//    uses COPY syntax to load data into an Amazon Redshift table. This is specified
//    in the RedshiftDestinationConfiguration.S3Configuration parameter.
//
//    * The compression formats SNAPPY or ZIP cannot be specified in RedshiftDestinationConfiguration.S3Configuration
//    because the Amazon Redshift COPY operation that reads from the S3 bucket
//    doesn't support these compression formats.
//
//    * We strongly recommend that you use the user name and password you provide
//    exclusively with Kinesis Data Firehose, and that the permissions for the
//    account are restricted for Amazon Redshift INSERT permissions.
//
// Kinesis Data Firehose assumes the IAM role that is configured as part of
// the destination. The role should allow the Kinesis Data Firehose principal
// to assume the role, and the role should have permissions that allow the service
// to deliver the data. For more information, see Grant Kinesis Data Firehose
// Access to an Amazon S3 Destination (https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-s3)
// in the Amazon Kinesis Data Firehose Developer Guide.
//
//    // Example sending a request using CreateDeliveryStreamRequest.
//    req := client.CreateDeliveryStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/CreateDeliveryStream
func (c *Client) CreateDeliveryStreamRequest(input *types.CreateDeliveryStreamInput) CreateDeliveryStreamRequest {
	op := &aws.Operation{
		Name:       opCreateDeliveryStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDeliveryStreamInput{}
	}

	req := c.newRequest(op, input, &types.CreateDeliveryStreamOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateDeliveryStreamMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDeliveryStreamRequest{Request: req, Input: input, Copy: c.CreateDeliveryStreamRequest}
}

// CreateDeliveryStreamRequest is the request type for the
// CreateDeliveryStream API operation.
type CreateDeliveryStreamRequest struct {
	*aws.Request
	Input *types.CreateDeliveryStreamInput
	Copy  func(*types.CreateDeliveryStreamInput) CreateDeliveryStreamRequest
}

// Send marshals and sends the CreateDeliveryStream API request.
func (r CreateDeliveryStreamRequest) Send(ctx context.Context) (*CreateDeliveryStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeliveryStreamResponse{
		CreateDeliveryStreamOutput: r.Request.Data.(*types.CreateDeliveryStreamOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeliveryStreamResponse is the response type for the
// CreateDeliveryStream API operation.
type CreateDeliveryStreamResponse struct {
	*types.CreateDeliveryStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeliveryStream request.
func (r *CreateDeliveryStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
