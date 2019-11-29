// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/firehose/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

const opListDeliveryStreams = "ListDeliveryStreams"

// ListDeliveryStreamsRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Lists your delivery streams in alphabetical order of their names.
//
// The number of delivery streams might be too large to return using a single
// call to ListDeliveryStreams. You can limit the number of delivery streams
// returned, using the Limit parameter. To determine whether there are more
// delivery streams to list, check the value of HasMoreDeliveryStreams in the
// output. If there are more delivery streams to list, you can request them
// by calling this operation again and setting the ExclusiveStartDeliveryStreamName
// parameter to the name of the last delivery stream returned in the last call.
//
//    // Example sending a request using ListDeliveryStreamsRequest.
//    req := client.ListDeliveryStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/ListDeliveryStreams
func (c *Client) ListDeliveryStreamsRequest(input *types.ListDeliveryStreamsInput) ListDeliveryStreamsRequest {
	op := &aws.Operation{
		Name:       opListDeliveryStreams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListDeliveryStreamsInput{}
	}

	req := c.newRequest(op, input, &types.ListDeliveryStreamsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListDeliveryStreamsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListDeliveryStreamsRequest{Request: req, Input: input, Copy: c.ListDeliveryStreamsRequest}
}

// ListDeliveryStreamsRequest is the request type for the
// ListDeliveryStreams API operation.
type ListDeliveryStreamsRequest struct {
	*aws.Request
	Input *types.ListDeliveryStreamsInput
	Copy  func(*types.ListDeliveryStreamsInput) ListDeliveryStreamsRequest
}

// Send marshals and sends the ListDeliveryStreams API request.
func (r ListDeliveryStreamsRequest) Send(ctx context.Context) (*ListDeliveryStreamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeliveryStreamsResponse{
		ListDeliveryStreamsOutput: r.Request.Data.(*types.ListDeliveryStreamsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDeliveryStreamsResponse is the response type for the
// ListDeliveryStreams API operation.
type ListDeliveryStreamsResponse struct {
	*types.ListDeliveryStreamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeliveryStreams request.
func (r *ListDeliveryStreamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
