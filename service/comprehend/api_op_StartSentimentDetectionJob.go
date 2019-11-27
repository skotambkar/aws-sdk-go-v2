// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opStartSentimentDetectionJob = "StartSentimentDetectionJob"

// StartSentimentDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Starts an asynchronous sentiment detection job for a collection of documents.
// use the operation to track the status of a job.
//
//    // Example sending a request using StartSentimentDetectionJobRequest.
//    req := client.StartSentimentDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StartSentimentDetectionJob
func (c *Client) StartSentimentDetectionJobRequest(input *types.StartSentimentDetectionJobInput) StartSentimentDetectionJobRequest {
	op := &aws.Operation{
		Name:       opStartSentimentDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartSentimentDetectionJobInput{}
	}

	req := c.newRequest(op, input, &types.StartSentimentDetectionJobOutput{})
	return StartSentimentDetectionJobRequest{Request: req, Input: input, Copy: c.StartSentimentDetectionJobRequest}
}

// StartSentimentDetectionJobRequest is the request type for the
// StartSentimentDetectionJob API operation.
type StartSentimentDetectionJobRequest struct {
	*aws.Request
	Input *types.StartSentimentDetectionJobInput
	Copy  func(*types.StartSentimentDetectionJobInput) StartSentimentDetectionJobRequest
}

// Send marshals and sends the StartSentimentDetectionJob API request.
func (r StartSentimentDetectionJobRequest) Send(ctx context.Context) (*StartSentimentDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSentimentDetectionJobResponse{
		StartSentimentDetectionJobOutput: r.Request.Data.(*types.StartSentimentDetectionJobOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSentimentDetectionJobResponse is the response type for the
// StartSentimentDetectionJob API operation.
type StartSentimentDetectionJobResponse struct {
	*types.StartSentimentDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSentimentDetectionJob request.
func (r *StartSentimentDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
