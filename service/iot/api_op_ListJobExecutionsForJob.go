// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListJobExecutionsForJob = "ListJobExecutionsForJob"

// ListJobExecutionsForJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the job executions for a job.
//
//    // Example sending a request using ListJobExecutionsForJobRequest.
//    req := client.ListJobExecutionsForJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListJobExecutionsForJobRequest(input *types.ListJobExecutionsForJobInput) ListJobExecutionsForJobRequest {
	op := &aws.Operation{
		Name:       opListJobExecutionsForJob,
		HTTPMethod: "GET",
		HTTPPath:   "/jobs/{jobId}/things",
	}

	if input == nil {
		input = &types.ListJobExecutionsForJobInput{}
	}

	req := c.newRequest(op, input, &types.ListJobExecutionsForJobOutput{})
	return ListJobExecutionsForJobRequest{Request: req, Input: input, Copy: c.ListJobExecutionsForJobRequest}
}

// ListJobExecutionsForJobRequest is the request type for the
// ListJobExecutionsForJob API operation.
type ListJobExecutionsForJobRequest struct {
	*aws.Request
	Input *types.ListJobExecutionsForJobInput
	Copy  func(*types.ListJobExecutionsForJobInput) ListJobExecutionsForJobRequest
}

// Send marshals and sends the ListJobExecutionsForJob API request.
func (r ListJobExecutionsForJobRequest) Send(ctx context.Context) (*ListJobExecutionsForJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobExecutionsForJobResponse{
		ListJobExecutionsForJobOutput: r.Request.Data.(*types.ListJobExecutionsForJobOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListJobExecutionsForJobResponse is the response type for the
// ListJobExecutionsForJob API operation.
type ListJobExecutionsForJobResponse struct {
	*types.ListJobExecutionsForJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobExecutionsForJob request.
func (r *ListJobExecutionsForJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
