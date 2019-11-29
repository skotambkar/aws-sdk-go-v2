// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opEnableDirectory = "EnableDirectory"

// EnableDirectoryRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Enables the specified directory. Only disabled directories can be enabled.
// Once enabled, the directory can then be read and written to.
//
//    // Example sending a request using EnableDirectoryRequest.
//    req := client.EnableDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/EnableDirectory
func (c *Client) EnableDirectoryRequest(input *types.EnableDirectoryInput) EnableDirectoryRequest {
	op := &aws.Operation{
		Name:       opEnableDirectory,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/directory/enable",
	}

	if input == nil {
		input = &types.EnableDirectoryInput{}
	}

	req := c.newRequest(op, input, &types.EnableDirectoryOutput{})
	return EnableDirectoryRequest{Request: req, Input: input, Copy: c.EnableDirectoryRequest}
}

// EnableDirectoryRequest is the request type for the
// EnableDirectory API operation.
type EnableDirectoryRequest struct {
	*aws.Request
	Input *types.EnableDirectoryInput
	Copy  func(*types.EnableDirectoryInput) EnableDirectoryRequest
}

// Send marshals and sends the EnableDirectory API request.
func (r EnableDirectoryRequest) Send(ctx context.Context) (*EnableDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableDirectoryResponse{
		EnableDirectoryOutput: r.Request.Data.(*types.EnableDirectoryOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableDirectoryResponse is the response type for the
// EnableDirectory API operation.
type EnableDirectoryResponse struct {
	*types.EnableDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableDirectory request.
func (r *EnableDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
