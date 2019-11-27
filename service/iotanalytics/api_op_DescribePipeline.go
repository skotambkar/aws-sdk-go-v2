// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opDescribePipeline = "DescribePipeline"

// DescribePipelineRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about a pipeline.
//
//    // Example sending a request using DescribePipelineRequest.
//    req := client.DescribePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribePipeline
func (c *Client) DescribePipelineRequest(input *types.DescribePipelineInput) DescribePipelineRequest {
	op := &aws.Operation{
		Name:       opDescribePipeline,
		HTTPMethod: "GET",
		HTTPPath:   "/pipelines/{pipelineName}",
	}

	if input == nil {
		input = &types.DescribePipelineInput{}
	}

	req := c.newRequest(op, input, &types.DescribePipelineOutput{})
	return DescribePipelineRequest{Request: req, Input: input, Copy: c.DescribePipelineRequest}
}

// DescribePipelineRequest is the request type for the
// DescribePipeline API operation.
type DescribePipelineRequest struct {
	*aws.Request
	Input *types.DescribePipelineInput
	Copy  func(*types.DescribePipelineInput) DescribePipelineRequest
}

// Send marshals and sends the DescribePipeline API request.
func (r DescribePipelineRequest) Send(ctx context.Context) (*DescribePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePipelineResponse{
		DescribePipelineOutput: r.Request.Data.(*types.DescribePipelineOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePipelineResponse is the response type for the
// DescribePipeline API operation.
type DescribePipelineResponse struct {
	*types.DescribePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePipeline request.
func (r *DescribePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
