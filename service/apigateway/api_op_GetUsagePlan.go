// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetUsagePlan = "GetUsagePlan"

// GetUsagePlanRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets a usage plan of a given plan identifier.
//
//    // Example sending a request using GetUsagePlanRequest.
//    req := client.GetUsagePlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetUsagePlanRequest(input *types.GetUsagePlanInput) GetUsagePlanRequest {
	op := &aws.Operation{
		Name:       opGetUsagePlan,
		HTTPMethod: "GET",
		HTTPPath:   "/usageplans/{usageplanId}",
	}

	if input == nil {
		input = &types.GetUsagePlanInput{}
	}

	req := c.newRequest(op, input, &types.GetUsagePlanOutput{})
	return GetUsagePlanRequest{Request: req, Input: input, Copy: c.GetUsagePlanRequest}
}

// GetUsagePlanRequest is the request type for the
// GetUsagePlan API operation.
type GetUsagePlanRequest struct {
	*aws.Request
	Input *types.GetUsagePlanInput
	Copy  func(*types.GetUsagePlanInput) GetUsagePlanRequest
}

// Send marshals and sends the GetUsagePlan API request.
func (r GetUsagePlanRequest) Send(ctx context.Context) (*GetUsagePlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUsagePlanResponse{
		GetUsagePlanOutput: r.Request.Data.(*types.GetUsagePlanOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUsagePlanResponse is the response type for the
// GetUsagePlan API operation.
type GetUsagePlanResponse struct {
	*types.GetUsagePlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUsagePlan request.
func (r *GetUsagePlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
