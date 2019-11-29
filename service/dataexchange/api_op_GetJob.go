// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
)

const opGetJob = "GetJob"

// GetJobRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation returns information about a job.
//
//    // Example sending a request using GetJobRequest.
//    req := client.GetJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/GetJob
func (c *Client) GetJobRequest(input *types.GetJobInput) GetJobRequest {
	op := &aws.Operation{
		Name:       opGetJob,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/jobs/{JobId}",
	}

	if input == nil {
		input = &types.GetJobInput{}
	}

	req := c.newRequest(op, input, &types.GetJobOutput{})
	return GetJobRequest{Request: req, Input: input, Copy: c.GetJobRequest}
}

// GetJobRequest is the request type for the
// GetJob API operation.
type GetJobRequest struct {
	*aws.Request
	Input *types.GetJobInput
	Copy  func(*types.GetJobInput) GetJobRequest
}

// Send marshals and sends the GetJob API request.
func (r GetJobRequest) Send(ctx context.Context) (*GetJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobResponse{
		GetJobOutput: r.Request.Data.(*types.GetJobOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobResponse is the response type for the
// GetJob API operation.
type GetJobResponse struct {
	*types.GetJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJob request.
func (r *GetJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
