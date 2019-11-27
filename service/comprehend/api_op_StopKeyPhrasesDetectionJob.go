// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opStopKeyPhrasesDetectionJob = "StopKeyPhrasesDetectionJob"

// StopKeyPhrasesDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Stops a key phrases detection job in progress.
//
// If the job state is IN_PROGRESS the job is marked for termination and put
// into the STOP_REQUESTED state. If the job completes before it can be stopped,
// it is put into the COMPLETED state; otherwise the job is stopped and put
// into the STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the StopDominantLanguageDetectionJob
// operation, the operation returns a 400 Internal Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
//
//    // Example sending a request using StopKeyPhrasesDetectionJobRequest.
//    req := client.StopKeyPhrasesDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StopKeyPhrasesDetectionJob
func (c *Client) StopKeyPhrasesDetectionJobRequest(input *types.StopKeyPhrasesDetectionJobInput) StopKeyPhrasesDetectionJobRequest {
	op := &aws.Operation{
		Name:       opStopKeyPhrasesDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopKeyPhrasesDetectionJobInput{}
	}

	req := c.newRequest(op, input, &types.StopKeyPhrasesDetectionJobOutput{})
	return StopKeyPhrasesDetectionJobRequest{Request: req, Input: input, Copy: c.StopKeyPhrasesDetectionJobRequest}
}

// StopKeyPhrasesDetectionJobRequest is the request type for the
// StopKeyPhrasesDetectionJob API operation.
type StopKeyPhrasesDetectionJobRequest struct {
	*aws.Request
	Input *types.StopKeyPhrasesDetectionJobInput
	Copy  func(*types.StopKeyPhrasesDetectionJobInput) StopKeyPhrasesDetectionJobRequest
}

// Send marshals and sends the StopKeyPhrasesDetectionJob API request.
func (r StopKeyPhrasesDetectionJobRequest) Send(ctx context.Context) (*StopKeyPhrasesDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopKeyPhrasesDetectionJobResponse{
		StopKeyPhrasesDetectionJobOutput: r.Request.Data.(*types.StopKeyPhrasesDetectionJobOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopKeyPhrasesDetectionJobResponse is the response type for the
// StopKeyPhrasesDetectionJob API operation.
type StopKeyPhrasesDetectionJobResponse struct {
	*types.StopKeyPhrasesDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopKeyPhrasesDetectionJob request.
func (r *StopKeyPhrasesDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
