// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opGetProject = "GetProject"

// GetProjectRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about a project.
//
//    // Example sending a request using GetProjectRequest.
//    req := client.GetProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetProject
func (c *Client) GetProjectRequest(input *types.GetProjectInput) GetProjectRequest {
	op := &aws.Operation{
		Name:       opGetProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetProjectInput{}
	}

	req := c.newRequest(op, input, &types.GetProjectOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetProjectMarshaler{Input: input}.GetNamedBuildHandler())

	return GetProjectRequest{Request: req, Input: input, Copy: c.GetProjectRequest}
}

// GetProjectRequest is the request type for the
// GetProject API operation.
type GetProjectRequest struct {
	*aws.Request
	Input *types.GetProjectInput
	Copy  func(*types.GetProjectInput) GetProjectRequest
}

// Send marshals and sends the GetProject API request.
func (r GetProjectRequest) Send(ctx context.Context) (*GetProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProjectResponse{
		GetProjectOutput: r.Request.Data.(*types.GetProjectOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetProjectResponse is the response type for the
// GetProject API operation.
type GetProjectResponse struct {
	*types.GetProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProject request.
func (r *GetProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
