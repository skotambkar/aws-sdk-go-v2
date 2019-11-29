// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListJobs = "ListJobs"

// ListJobsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists jobs.
//
//    // Example sending a request using ListJobsRequest.
//    req := client.ListJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListJobsRequest(input *types.ListJobsInput) ListJobsRequest {
	op := &aws.Operation{
		Name:       opListJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/jobs",
	}

	if input == nil {
		input = &types.ListJobsInput{}
	}

	req := c.newRequest(op, input, &types.ListJobsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListJobsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListJobsRequest{Request: req, Input: input, Copy: c.ListJobsRequest}
}

// ListJobsRequest is the request type for the
// ListJobs API operation.
type ListJobsRequest struct {
	*aws.Request
	Input *types.ListJobsInput
	Copy  func(*types.ListJobsInput) ListJobsRequest
}

// Send marshals and sends the ListJobs API request.
func (r ListJobsRequest) Send(ctx context.Context) (*ListJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobsResponse{
		ListJobsOutput: r.Request.Data.(*types.ListJobsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListJobsResponse is the response type for the
// ListJobs API operation.
type ListJobsResponse struct {
	*types.ListJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobs request.
func (r *ListJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
