// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
)

const opStartPHIDetectionJob = "StartPHIDetectionJob"

// StartPHIDetectionJobRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// Starts an asynchronous job to detect protected health information (PHI).
// Use the DescribePHIDetectionJob operation to track the status of a job.
//
//    // Example sending a request using StartPHIDetectionJobRequest.
//    req := client.StartPHIDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/StartPHIDetectionJob
func (c *Client) StartPHIDetectionJobRequest(input *types.StartPHIDetectionJobInput) StartPHIDetectionJobRequest {
	op := &aws.Operation{
		Name:       opStartPHIDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartPHIDetectionJobInput{}
	}

	req := c.newRequest(op, input, &types.StartPHIDetectionJobOutput{})
	return StartPHIDetectionJobRequest{Request: req, Input: input, Copy: c.StartPHIDetectionJobRequest}
}

// StartPHIDetectionJobRequest is the request type for the
// StartPHIDetectionJob API operation.
type StartPHIDetectionJobRequest struct {
	*aws.Request
	Input *types.StartPHIDetectionJobInput
	Copy  func(*types.StartPHIDetectionJobInput) StartPHIDetectionJobRequest
}

// Send marshals and sends the StartPHIDetectionJob API request.
func (r StartPHIDetectionJobRequest) Send(ctx context.Context) (*StartPHIDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartPHIDetectionJobResponse{
		StartPHIDetectionJobOutput: r.Request.Data.(*types.StartPHIDetectionJobOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartPHIDetectionJobResponse is the response type for the
// StartPHIDetectionJob API operation.
type StartPHIDetectionJobResponse struct {
	*types.StartPHIDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartPHIDetectionJob request.
func (r *StartPHIDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
