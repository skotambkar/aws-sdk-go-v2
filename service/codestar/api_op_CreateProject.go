// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestar/types"
)

const opCreateProject = "CreateProject"

// CreateProjectRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Creates a project, including project resources. This action creates a project
// based on a submitted project request. A set of source code files and a toolchain
// template file can be included with the project request. If these are not
// provided, an empty project is created.
//
//    // Example sending a request using CreateProjectRequest.
//    req := client.CreateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/CreateProject
func (c *Client) CreateProjectRequest(input *types.CreateProjectInput) CreateProjectRequest {
	op := &aws.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateProjectInput{}
	}

	req := c.newRequest(op, input, &types.CreateProjectOutput{})
	return CreateProjectRequest{Request: req, Input: input, Copy: c.CreateProjectRequest}
}

// CreateProjectRequest is the request type for the
// CreateProject API operation.
type CreateProjectRequest struct {
	*aws.Request
	Input *types.CreateProjectInput
	Copy  func(*types.CreateProjectInput) CreateProjectRequest
}

// Send marshals and sends the CreateProject API request.
func (r CreateProjectRequest) Send(ctx context.Context) (*CreateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectResponse{
		CreateProjectOutput: r.Request.Data.(*types.CreateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectResponse is the response type for the
// CreateProject API operation.
type CreateProjectResponse struct {
	*types.CreateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProject request.
func (r *CreateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
