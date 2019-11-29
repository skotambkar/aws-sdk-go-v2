// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
)

const opCreateJob = "CreateJob"

// CreateJobRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Create a new transcoding job. For information about jobs and job settings,
// see the User Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
//
//    // Example sending a request using CreateJobRequest.
//    req := client.CreateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/CreateJob
func (c *Client) CreateJobRequest(input *types.CreateJobInput) CreateJobRequest {
	op := &aws.Operation{
		Name:       opCreateJob,
		HTTPMethod: "POST",
		HTTPPath:   "/2017-08-29/jobs",
	}

	if input == nil {
		input = &types.CreateJobInput{}
	}

	req := c.newRequest(op, input, &types.CreateJobOutput{})
	return CreateJobRequest{Request: req, Input: input, Copy: c.CreateJobRequest}
}

// CreateJobRequest is the request type for the
// CreateJob API operation.
type CreateJobRequest struct {
	*aws.Request
	Input *types.CreateJobInput
	Copy  func(*types.CreateJobInput) CreateJobRequest
}

// Send marshals and sends the CreateJob API request.
func (r CreateJobRequest) Send(ctx context.Context) (*CreateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJobResponse{
		CreateJobOutput: r.Request.Data.(*types.CreateJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJobResponse is the response type for the
// CreateJob API operation.
type CreateJobResponse struct {
	*types.CreateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJob request.
func (r *CreateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
