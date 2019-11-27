// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetModel = "GetModel"

// GetModelRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Describes an existing model defined for a RestApi resource.
//
//    // Example sending a request using GetModelRequest.
//    req := client.GetModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetModelRequest(input *types.GetModelInput) GetModelRequest {
	op := &aws.Operation{
		Name:       opGetModel,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/models/{model_name}",
	}

	if input == nil {
		input = &types.GetModelInput{}
	}

	req := c.newRequest(op, input, &types.GetModelOutput{})
	return GetModelRequest{Request: req, Input: input, Copy: c.GetModelRequest}
}

// GetModelRequest is the request type for the
// GetModel API operation.
type GetModelRequest struct {
	*aws.Request
	Input *types.GetModelInput
	Copy  func(*types.GetModelInput) GetModelRequest
}

// Send marshals and sends the GetModel API request.
func (r GetModelRequest) Send(ctx context.Context) (*GetModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetModelResponse{
		GetModelOutput: r.Request.Data.(*types.GetModelOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetModelResponse is the response type for the
// GetModel API operation.
type GetModelResponse struct {
	*types.GetModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetModel request.
func (r *GetModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
