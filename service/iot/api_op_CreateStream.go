// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opCreateStream = "CreateStream"

// CreateStreamRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a stream for delivering one or more large files in chunks over MQTT.
// A stream transports data bytes in chunks or blocks packaged as MQTT messages
// from a source like S3. You can have one or more files associated with a stream.
// The total size of a file associated with the stream cannot exceed more than
// 2 MB. The stream will be created with version 0. If a stream is created with
// the same streamID as a stream that existed and was deleted within last 90
// days, we will resurrect that old stream by incrementing the version by 1.
//
//    // Example sending a request using CreateStreamRequest.
//    req := client.CreateStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateStreamRequest(input *types.CreateStreamInput) CreateStreamRequest {
	op := &aws.Operation{
		Name:       opCreateStream,
		HTTPMethod: "POST",
		HTTPPath:   "/streams/{streamId}",
	}

	if input == nil {
		input = &types.CreateStreamInput{}
	}

	req := c.newRequest(op, input, &types.CreateStreamOutput{})
	return CreateStreamRequest{Request: req, Input: input, Copy: c.CreateStreamRequest}
}

// CreateStreamRequest is the request type for the
// CreateStream API operation.
type CreateStreamRequest struct {
	*aws.Request
	Input *types.CreateStreamInput
	Copy  func(*types.CreateStreamInput) CreateStreamRequest
}

// Send marshals and sends the CreateStream API request.
func (r CreateStreamRequest) Send(ctx context.Context) (*CreateStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStreamResponse{
		CreateStreamOutput: r.Request.Data.(*types.CreateStreamOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStreamResponse is the response type for the
// CreateStream API operation.
type CreateStreamResponse struct {
	*types.CreateStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStream request.
func (r *CreateStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
