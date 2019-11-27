// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
)

const opUpdatePipeline = "UpdatePipeline"

// UpdatePipelineRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// Use the UpdatePipeline operation to update settings for a pipeline.
//
// When you change pipeline settings, your changes take effect immediately.
// Jobs that you have already submitted and that Elastic Transcoder has not
// started to process are affected in addition to jobs that you submit after
// you change settings.
//
//    // Example sending a request using UpdatePipelineRequest.
//    req := client.UpdatePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdatePipelineRequest(input *types.UpdatePipelineInput) UpdatePipelineRequest {
	op := &aws.Operation{
		Name:       opUpdatePipeline,
		HTTPMethod: "PUT",
		HTTPPath:   "/2012-09-25/pipelines/{Id}",
	}

	if input == nil {
		input = &types.UpdatePipelineInput{}
	}

	req := c.newRequest(op, input, &types.UpdatePipelineOutput{})
	return UpdatePipelineRequest{Request: req, Input: input, Copy: c.UpdatePipelineRequest}
}

// UpdatePipelineRequest is the request type for the
// UpdatePipeline API operation.
type UpdatePipelineRequest struct {
	*aws.Request
	Input *types.UpdatePipelineInput
	Copy  func(*types.UpdatePipelineInput) UpdatePipelineRequest
}

// Send marshals and sends the UpdatePipeline API request.
func (r UpdatePipelineRequest) Send(ctx context.Context) (*UpdatePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePipelineResponse{
		UpdatePipelineOutput: r.Request.Data.(*types.UpdatePipelineOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePipelineResponse is the response type for the
// UpdatePipeline API operation.
type UpdatePipelineResponse struct {
	*types.UpdatePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePipeline request.
func (r *UpdatePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
