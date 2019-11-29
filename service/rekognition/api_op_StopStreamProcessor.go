// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const opStopStreamProcessor = "StopStreamProcessor"

// StopStreamProcessorRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Stops a running stream processor that was created by CreateStreamProcessor.
//
//    // Example sending a request using StopStreamProcessorRequest.
//    req := client.StopStreamProcessorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StopStreamProcessorRequest(input *types.StopStreamProcessorInput) StopStreamProcessorRequest {
	op := &aws.Operation{
		Name:       opStopStreamProcessor,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopStreamProcessorInput{}
	}

	req := c.newRequest(op, input, &types.StopStreamProcessorOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StopStreamProcessorMarshaler{Input: input}.GetNamedBuildHandler())

	return StopStreamProcessorRequest{Request: req, Input: input, Copy: c.StopStreamProcessorRequest}
}

// StopStreamProcessorRequest is the request type for the
// StopStreamProcessor API operation.
type StopStreamProcessorRequest struct {
	*aws.Request
	Input *types.StopStreamProcessorInput
	Copy  func(*types.StopStreamProcessorInput) StopStreamProcessorRequest
}

// Send marshals and sends the StopStreamProcessor API request.
func (r StopStreamProcessorRequest) Send(ctx context.Context) (*StopStreamProcessorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopStreamProcessorResponse{
		StopStreamProcessorOutput: r.Request.Data.(*types.StopStreamProcessorOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopStreamProcessorResponse is the response type for the
// StopStreamProcessor API operation.
type StopStreamProcessorResponse struct {
	*types.StopStreamProcessorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopStreamProcessor request.
func (r *StopStreamProcessorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
