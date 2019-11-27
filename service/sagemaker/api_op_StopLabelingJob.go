// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opStopLabelingJob = "StopLabelingJob"

// StopLabelingJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Stops a running labeling job. A job that is stopped cannot be restarted.
// Any results obtained before the job is stopped are placed in the Amazon S3
// output bucket.
//
//    // Example sending a request using StopLabelingJobRequest.
//    req := client.StopLabelingJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopLabelingJob
func (c *Client) StopLabelingJobRequest(input *types.StopLabelingJobInput) StopLabelingJobRequest {
	op := &aws.Operation{
		Name:       opStopLabelingJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopLabelingJobInput{}
	}

	req := c.newRequest(op, input, &types.StopLabelingJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StopLabelingJobRequest{Request: req, Input: input, Copy: c.StopLabelingJobRequest}
}

// StopLabelingJobRequest is the request type for the
// StopLabelingJob API operation.
type StopLabelingJobRequest struct {
	*aws.Request
	Input *types.StopLabelingJobInput
	Copy  func(*types.StopLabelingJobInput) StopLabelingJobRequest
}

// Send marshals and sends the StopLabelingJob API request.
func (r StopLabelingJobRequest) Send(ctx context.Context) (*StopLabelingJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopLabelingJobResponse{
		StopLabelingJobOutput: r.Request.Data.(*types.StopLabelingJobOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopLabelingJobResponse is the response type for the
// StopLabelingJob API operation.
type StopLabelingJobResponse struct {
	*types.StopLabelingJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopLabelingJob request.
func (r *StopLabelingJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
