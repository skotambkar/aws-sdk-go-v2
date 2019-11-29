// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opStopJob = "StopJob"

// StopJobRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Initiates a stop request for the current job. AWS Device Farm will immediately
// stop the job on the device where tests have not started executing, and you
// will not be billed for this device. On the device where tests have started
// executing, Setup Suite and Teardown Suite tests will run to completion before
// stopping execution on the device. You will be billed for Setup, Teardown,
// and any tests that were in progress or already completed.
//
//    // Example sending a request using StopJobRequest.
//    req := client.StopJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/StopJob
func (c *Client) StopJobRequest(input *types.StopJobInput) StopJobRequest {
	op := &aws.Operation{
		Name:       opStopJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopJobInput{}
	}

	req := c.newRequest(op, input, &types.StopJobOutput{})
	return StopJobRequest{Request: req, Input: input, Copy: c.StopJobRequest}
}

// StopJobRequest is the request type for the
// StopJob API operation.
type StopJobRequest struct {
	*aws.Request
	Input *types.StopJobInput
	Copy  func(*types.StopJobInput) StopJobRequest
}

// Send marshals and sends the StopJob API request.
func (r StopJobRequest) Send(ctx context.Context) (*StopJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopJobResponse{
		StopJobOutput: r.Request.Data.(*types.StopJobOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopJobResponse is the response type for the
// StopJob API operation.
type StopJobResponse struct {
	*types.StopJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopJob request.
func (r *StopJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
