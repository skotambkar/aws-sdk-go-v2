// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetExportJobs = "GetExportJobs"

// GetExportJobsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of all the export jobs
// for an application.
//
//    // Example sending a request using GetExportJobsRequest.
//    req := client.GetExportJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetExportJobs
func (c *Client) GetExportJobsRequest(input *types.GetExportJobsInput) GetExportJobsRequest {
	op := &aws.Operation{
		Name:       opGetExportJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/jobs/export",
	}

	if input == nil {
		input = &types.GetExportJobsInput{}
	}

	req := c.newRequest(op, input, &types.GetExportJobsOutput{})
	return GetExportJobsRequest{Request: req, Input: input, Copy: c.GetExportJobsRequest}
}

// GetExportJobsRequest is the request type for the
// GetExportJobs API operation.
type GetExportJobsRequest struct {
	*aws.Request
	Input *types.GetExportJobsInput
	Copy  func(*types.GetExportJobsInput) GetExportJobsRequest
}

// Send marshals and sends the GetExportJobs API request.
func (r GetExportJobsRequest) Send(ctx context.Context) (*GetExportJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetExportJobsResponse{
		GetExportJobsOutput: r.Request.Data.(*types.GetExportJobsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetExportJobsResponse is the response type for the
// GetExportJobs API operation.
type GetExportJobsResponse struct {
	*types.GetExportJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetExportJobs request.
func (r *GetExportJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
