// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/appstream/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opStartImageBuilder = "StartImageBuilder"

// StartImageBuilderRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Starts the specified image builder.
//
//    // Example sending a request using StartImageBuilderRequest.
//    req := client.StartImageBuilderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/StartImageBuilder
func (c *Client) StartImageBuilderRequest(input *types.StartImageBuilderInput) StartImageBuilderRequest {
	op := &aws.Operation{
		Name:       opStartImageBuilder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartImageBuilderInput{}
	}

	req := c.newRequest(op, input, &types.StartImageBuilderOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StartImageBuilderMarshaler{Input: input}.GetNamedBuildHandler())

	return StartImageBuilderRequest{Request: req, Input: input, Copy: c.StartImageBuilderRequest}
}

// StartImageBuilderRequest is the request type for the
// StartImageBuilder API operation.
type StartImageBuilderRequest struct {
	*aws.Request
	Input *types.StartImageBuilderInput
	Copy  func(*types.StartImageBuilderInput) StartImageBuilderRequest
}

// Send marshals and sends the StartImageBuilder API request.
func (r StartImageBuilderRequest) Send(ctx context.Context) (*StartImageBuilderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartImageBuilderResponse{
		StartImageBuilderOutput: r.Request.Data.(*types.StartImageBuilderOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartImageBuilderResponse is the response type for the
// StartImageBuilder API operation.
type StartImageBuilderResponse struct {
	*types.StartImageBuilderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartImageBuilder request.
func (r *StartImageBuilderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
