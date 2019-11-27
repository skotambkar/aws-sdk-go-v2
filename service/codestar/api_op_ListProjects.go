// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestar/types"
)

const opListProjects = "ListProjects"

// ListProjectsRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Lists all projects in AWS CodeStar associated with your AWS account.
//
//    // Example sending a request using ListProjectsRequest.
//    req := client.ListProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListProjects
func (c *Client) ListProjectsRequest(input *types.ListProjectsInput) ListProjectsRequest {
	op := &aws.Operation{
		Name:       opListProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListProjectsInput{}
	}

	req := c.newRequest(op, input, &types.ListProjectsOutput{})
	return ListProjectsRequest{Request: req, Input: input, Copy: c.ListProjectsRequest}
}

// ListProjectsRequest is the request type for the
// ListProjects API operation.
type ListProjectsRequest struct {
	*aws.Request
	Input *types.ListProjectsInput
	Copy  func(*types.ListProjectsInput) ListProjectsRequest
}

// Send marshals and sends the ListProjects API request.
func (r ListProjectsRequest) Send(ctx context.Context) (*ListProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProjectsResponse{
		ListProjectsOutput: r.Request.Data.(*types.ListProjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListProjectsResponse is the response type for the
// ListProjects API operation.
type ListProjectsResponse struct {
	*types.ListProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProjects request.
func (r *ListProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
