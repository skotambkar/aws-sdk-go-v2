// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
)

const opUpdateProject = "UpdateProject"

// UpdateProjectRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Changes the settings of a build project.
//
//    // Example sending a request using UpdateProjectRequest.
//    req := client.UpdateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/UpdateProject
func (c *Client) UpdateProjectRequest(input *types.UpdateProjectInput) UpdateProjectRequest {
	op := &aws.Operation{
		Name:       opUpdateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateProjectInput{}
	}

	req := c.newRequest(op, input, &types.UpdateProjectOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateProjectMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateProjectRequest{Request: req, Input: input, Copy: c.UpdateProjectRequest}
}

// UpdateProjectRequest is the request type for the
// UpdateProject API operation.
type UpdateProjectRequest struct {
	*aws.Request
	Input *types.UpdateProjectInput
	Copy  func(*types.UpdateProjectInput) UpdateProjectRequest
}

// Send marshals and sends the UpdateProject API request.
func (r UpdateProjectRequest) Send(ctx context.Context) (*UpdateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProjectResponse{
		UpdateProjectOutput: r.Request.Data.(*types.UpdateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProjectResponse is the response type for the
// UpdateProject API operation.
type UpdateProjectResponse struct {
	*types.UpdateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProject request.
func (r *UpdateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
