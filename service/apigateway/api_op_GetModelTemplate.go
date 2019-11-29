// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetModelTemplate = "GetModelTemplate"

// GetModelTemplateRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Generates a sample mapping template that can be used to transform a payload
// into the structure of a model.
//
//    // Example sending a request using GetModelTemplateRequest.
//    req := client.GetModelTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetModelTemplateRequest(input *types.GetModelTemplateInput) GetModelTemplateRequest {
	op := &aws.Operation{
		Name:       opGetModelTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/models/{model_name}/default_template",
	}

	if input == nil {
		input = &types.GetModelTemplateInput{}
	}

	req := c.newRequest(op, input, &types.GetModelTemplateOutput{})
	return GetModelTemplateRequest{Request: req, Input: input, Copy: c.GetModelTemplateRequest}
}

// GetModelTemplateRequest is the request type for the
// GetModelTemplate API operation.
type GetModelTemplateRequest struct {
	*aws.Request
	Input *types.GetModelTemplateInput
	Copy  func(*types.GetModelTemplateInput) GetModelTemplateRequest
}

// Send marshals and sends the GetModelTemplate API request.
func (r GetModelTemplateRequest) Send(ctx context.Context) (*GetModelTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetModelTemplateResponse{
		GetModelTemplateOutput: r.Request.Data.(*types.GetModelTemplateOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetModelTemplateResponse is the response type for the
// GetModelTemplate API operation.
type GetModelTemplateResponse struct {
	*types.GetModelTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetModelTemplate request.
func (r *GetModelTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
