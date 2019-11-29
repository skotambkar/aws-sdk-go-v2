// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
)

const opDescribeStreamConsumer = "DescribeStreamConsumer"

// DescribeStreamConsumerRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// To get the description of a registered consumer, provide the ARN of the consumer.
// Alternatively, you can provide the ARN of the data stream and the name you
// gave the consumer when you registered it. You may also provide all three
// parameters, as long as they don't conflict with each other. If you don't
// know the name or ARN of the consumer that you want to describe, you can use
// the ListStreamConsumers operation to get a list of the descriptions of all
// the consumers that are currently registered with a given data stream.
//
// This operation has a limit of 20 transactions per second per account.
//
//    // Example sending a request using DescribeStreamConsumerRequest.
//    req := client.DescribeStreamConsumerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/DescribeStreamConsumer
func (c *Client) DescribeStreamConsumerRequest(input *types.DescribeStreamConsumerInput) DescribeStreamConsumerRequest {
	op := &aws.Operation{
		Name:       opDescribeStreamConsumer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeStreamConsumerInput{}
	}

	req := c.newRequest(op, input, &types.DescribeStreamConsumerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeStreamConsumerMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeStreamConsumerRequest{Request: req, Input: input, Copy: c.DescribeStreamConsumerRequest}
}

// DescribeStreamConsumerRequest is the request type for the
// DescribeStreamConsumer API operation.
type DescribeStreamConsumerRequest struct {
	*aws.Request
	Input *types.DescribeStreamConsumerInput
	Copy  func(*types.DescribeStreamConsumerInput) DescribeStreamConsumerRequest
}

// Send marshals and sends the DescribeStreamConsumer API request.
func (r DescribeStreamConsumerRequest) Send(ctx context.Context) (*DescribeStreamConsumerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStreamConsumerResponse{
		DescribeStreamConsumerOutput: r.Request.Data.(*types.DescribeStreamConsumerOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStreamConsumerResponse is the response type for the
// DescribeStreamConsumer API operation.
type DescribeStreamConsumerResponse struct {
	*types.DescribeStreamConsumerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStreamConsumer request.
func (r *DescribeStreamConsumerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
