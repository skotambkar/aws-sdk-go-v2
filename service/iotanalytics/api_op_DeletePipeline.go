// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opDeletePipeline = "DeletePipeline"

// DeletePipelineRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Deletes the specified pipeline.
//
//    // Example sending a request using DeletePipelineRequest.
//    req := client.DeletePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DeletePipeline
func (c *Client) DeletePipelineRequest(input *types.DeletePipelineInput) DeletePipelineRequest {
	op := &aws.Operation{
		Name:       opDeletePipeline,
		HTTPMethod: "DELETE",
		HTTPPath:   "/pipelines/{pipelineName}",
	}

	if input == nil {
		input = &types.DeletePipelineInput{}
	}

	req := c.newRequest(op, input, &types.DeletePipelineOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePipelineRequest{Request: req, Input: input, Copy: c.DeletePipelineRequest}
}

// DeletePipelineRequest is the request type for the
// DeletePipeline API operation.
type DeletePipelineRequest struct {
	*aws.Request
	Input *types.DeletePipelineInput
	Copy  func(*types.DeletePipelineInput) DeletePipelineRequest
}

// Send marshals and sends the DeletePipeline API request.
func (r DeletePipelineRequest) Send(ctx context.Context) (*DeletePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePipelineResponse{
		DeletePipelineOutput: r.Request.Data.(*types.DeletePipelineOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePipelineResponse is the response type for the
// DeletePipeline API operation.
type DeletePipelineResponse struct {
	*types.DeletePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePipeline request.
func (r *DeletePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
