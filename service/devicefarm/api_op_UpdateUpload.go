// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opUpdateUpload = "UpdateUpload"

// UpdateUploadRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Update an uploaded test specification (test spec).
//
//    // Example sending a request using UpdateUploadRequest.
//    req := client.UpdateUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/UpdateUpload
func (c *Client) UpdateUploadRequest(input *types.UpdateUploadInput) UpdateUploadRequest {
	op := &aws.Operation{
		Name:       opUpdateUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateUploadInput{}
	}

	req := c.newRequest(op, input, &types.UpdateUploadOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateUploadMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateUploadRequest{Request: req, Input: input, Copy: c.UpdateUploadRequest}
}

// UpdateUploadRequest is the request type for the
// UpdateUpload API operation.
type UpdateUploadRequest struct {
	*aws.Request
	Input *types.UpdateUploadInput
	Copy  func(*types.UpdateUploadInput) UpdateUploadRequest
}

// Send marshals and sends the UpdateUpload API request.
func (r UpdateUploadRequest) Send(ctx context.Context) (*UpdateUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUploadResponse{
		UpdateUploadOutput: r.Request.Data.(*types.UpdateUploadOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUploadResponse is the response type for the
// UpdateUpload API operation.
type UpdateUploadResponse struct {
	*types.UpdateUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUpload request.
func (r *UpdateUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
