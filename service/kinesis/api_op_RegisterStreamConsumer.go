// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
)

const opRegisterStreamConsumer = "RegisterStreamConsumer"

// RegisterStreamConsumerRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Registers a consumer with a Kinesis data stream. When you use this operation,
// the consumer you register can read data from the stream at a rate of up to
// 2 MiB per second. This rate is unaffected by the total number of consumers
// that read from the same stream.
//
// You can register up to 5 consumers per stream. A given consumer can only
// be registered with one stream.
//
// This operation has a limit of five transactions per second per account.
//
//    // Example sending a request using RegisterStreamConsumerRequest.
//    req := client.RegisterStreamConsumerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/RegisterStreamConsumer
func (c *Client) RegisterStreamConsumerRequest(input *types.RegisterStreamConsumerInput) RegisterStreamConsumerRequest {
	op := &aws.Operation{
		Name:       opRegisterStreamConsumer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterStreamConsumerInput{}
	}

	req := c.newRequest(op, input, &types.RegisterStreamConsumerOutput{})
	return RegisterStreamConsumerRequest{Request: req, Input: input, Copy: c.RegisterStreamConsumerRequest}
}

// RegisterStreamConsumerRequest is the request type for the
// RegisterStreamConsumer API operation.
type RegisterStreamConsumerRequest struct {
	*aws.Request
	Input *types.RegisterStreamConsumerInput
	Copy  func(*types.RegisterStreamConsumerInput) RegisterStreamConsumerRequest
}

// Send marshals and sends the RegisterStreamConsumer API request.
func (r RegisterStreamConsumerRequest) Send(ctx context.Context) (*RegisterStreamConsumerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterStreamConsumerResponse{
		RegisterStreamConsumerOutput: r.Request.Data.(*types.RegisterStreamConsumerOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterStreamConsumerResponse is the response type for the
// RegisterStreamConsumer API operation.
type RegisterStreamConsumerResponse struct {
	*types.RegisterStreamConsumerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterStreamConsumer request.
func (r *RegisterStreamConsumerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
