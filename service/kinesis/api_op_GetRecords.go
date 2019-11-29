// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
)

const opGetRecords = "GetRecords"

// GetRecordsRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Gets data records from a Kinesis data stream's shard.
//
// Specify a shard iterator using the ShardIterator parameter. The shard iterator
// specifies the position in the shard from which you want to start reading
// data records sequentially. If there are no records available in the portion
// of the shard that the iterator points to, GetRecords returns an empty list.
// It might take multiple calls to get to a portion of the shard that contains
// records.
//
// You can scale by provisioning multiple shards per stream while considering
// service limits (for more information, see Amazon Kinesis Data Streams Limits
// (http://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide). Your application should
// have one thread per shard, each reading continuously from its stream. To
// read from a stream continually, call GetRecords in a loop. Use GetShardIterator
// to get the shard iterator to specify in the first GetRecords call. GetRecords
// returns a new shard iterator in NextShardIterator. Specify the shard iterator
// returned in NextShardIterator in subsequent calls to GetRecords. If the shard
// has been closed, the shard iterator can't return more data and GetRecords
// returns null in NextShardIterator. You can terminate the loop when the shard
// is closed, or when the shard iterator reaches the record with the sequence
// number or other attribute that marks it as the last record to process.
//
// Each data record can be up to 1 MiB in size, and each shard can read up to
// 2 MiB per second. You can ensure that your calls don't exceed the maximum
// supported size or throughput by using the Limit parameter to specify the
// maximum number of records that GetRecords can return. Consider your average
// record size when determining this limit. The maximum number of records that
// can be returned per call is 10,000.
//
// The size of the data returned by GetRecords varies depending on the utilization
// of the shard. The maximum size of data that GetRecords can return is 10 MiB.
// If a call returns this amount of data, subsequent calls made within the next
// 5 seconds throw ProvisionedThroughputExceededException. If there is insufficient
// provisioned throughput on the stream, subsequent calls made within the next
// 1 second throw ProvisionedThroughputExceededException. GetRecords doesn't
// return any data when it throws an exception. For this reason, we recommend
// that you wait 1 second between calls to GetRecords. However, it's possible
// that the application will get exceptions for longer than 1 second.
//
// To detect whether the application is falling behind in processing, you can
// use the MillisBehindLatest response attribute. You can also monitor the stream
// using CloudWatch metrics and other mechanisms (see Monitoring (http://docs.aws.amazon.com/kinesis/latest/dev/monitoring.html)
// in the Amazon Kinesis Data Streams Developer Guide).
//
// Each Amazon Kinesis record includes a value, ApproximateArrivalTimestamp,
// that is set when a stream successfully receives and stores a record. This
// is commonly referred to as a server-side time stamp, whereas a client-side
// time stamp is set when a data producer creates or sends the record to a stream
// (a data producer is any data source putting data records into a stream, for
// example with PutRecords). The time stamp has millisecond precision. There
// are no guarantees about the time stamp accuracy, or that the time stamp is
// always increasing. For example, records in a shard or across a stream might
// have time stamps that are out of order.
//
// This operation has a limit of five transactions per second per account.
//
//    // Example sending a request using GetRecordsRequest.
//    req := client.GetRecordsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/GetRecords
func (c *Client) GetRecordsRequest(input *types.GetRecordsInput) GetRecordsRequest {
	op := &aws.Operation{
		Name:       opGetRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRecordsInput{}
	}

	req := c.newRequest(op, input, &types.GetRecordsOutput{})
	return GetRecordsRequest{Request: req, Input: input, Copy: c.GetRecordsRequest}
}

// GetRecordsRequest is the request type for the
// GetRecords API operation.
type GetRecordsRequest struct {
	*aws.Request
	Input *types.GetRecordsInput
	Copy  func(*types.GetRecordsInput) GetRecordsRequest
}

// Send marshals and sends the GetRecords API request.
func (r GetRecordsRequest) Send(ctx context.Context) (*GetRecordsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRecordsResponse{
		GetRecordsOutput: r.Request.Data.(*types.GetRecordsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRecordsResponse is the response type for the
// GetRecords API operation.
type GetRecordsResponse struct {
	*types.GetRecordsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRecords request.
func (r *GetRecordsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
