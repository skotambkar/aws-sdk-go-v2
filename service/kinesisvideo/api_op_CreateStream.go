// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
)

const opCreateStream = "CreateStream"

// CreateStreamRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Creates a new Kinesis video stream.
//
// When you create a new stream, Kinesis Video Streams assigns it a version
// number. When you change the stream's metadata, Kinesis Video Streams updates
// the version.
//
// CreateStream is an asynchronous operation.
//
// For information about how the service works, see How it Works (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/how-it-works.html).
//
// You must have permissions for the KinesisVideo:CreateStream action.
//
//    // Example sending a request using CreateStreamRequest.
//    req := client.CreateStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/CreateStream
func (c *Client) CreateStreamRequest(input *types.CreateStreamInput) CreateStreamRequest {
	op := &aws.Operation{
		Name:       opCreateStream,
		HTTPMethod: "POST",
		HTTPPath:   "/createStream",
	}

	if input == nil {
		input = &types.CreateStreamInput{}
	}

	req := c.newRequest(op, input, &types.CreateStreamOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateStreamMarshaler{Input: input}.GetNamedBuildHandler())

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
