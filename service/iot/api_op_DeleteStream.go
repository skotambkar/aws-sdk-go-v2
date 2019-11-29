// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDeleteStream = "DeleteStream"

// DeleteStreamRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a stream.
//
//    // Example sending a request using DeleteStreamRequest.
//    req := client.DeleteStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteStreamRequest(input *types.DeleteStreamInput) DeleteStreamRequest {
	op := &aws.Operation{
		Name:       opDeleteStream,
		HTTPMethod: "DELETE",
		HTTPPath:   "/streams/{streamId}",
	}

	if input == nil {
		input = &types.DeleteStreamInput{}
	}

	req := c.newRequest(op, input, &types.DeleteStreamOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteStreamMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteStreamRequest{Request: req, Input: input, Copy: c.DeleteStreamRequest}
}

// DeleteStreamRequest is the request type for the
// DeleteStream API operation.
type DeleteStreamRequest struct {
	*aws.Request
	Input *types.DeleteStreamInput
	Copy  func(*types.DeleteStreamInput) DeleteStreamRequest
}

// Send marshals and sends the DeleteStream API request.
func (r DeleteStreamRequest) Send(ctx context.Context) (*DeleteStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteStreamResponse{
		DeleteStreamOutput: r.Request.Data.(*types.DeleteStreamOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteStreamResponse is the response type for the
// DeleteStream API operation.
type DeleteStreamResponse struct {
	*types.DeleteStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteStream request.
func (r *DeleteStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
