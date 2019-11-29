// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opUpdateSMBSecurityStrategy = "UpdateSMBSecurityStrategy"

// UpdateSMBSecurityStrategyRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates the SMB security strategy on a file gateway. This action is only
// supported in file gateways.
//
// This API is called Security level in the User Guide.
//
// A higher security level can affect performance of the gateway.
//
//    // Example sending a request using UpdateSMBSecurityStrategyRequest.
//    req := client.UpdateSMBSecurityStrategyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateSMBSecurityStrategy
func (c *Client) UpdateSMBSecurityStrategyRequest(input *types.UpdateSMBSecurityStrategyInput) UpdateSMBSecurityStrategyRequest {
	op := &aws.Operation{
		Name:       opUpdateSMBSecurityStrategy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateSMBSecurityStrategyInput{}
	}

	req := c.newRequest(op, input, &types.UpdateSMBSecurityStrategyOutput{})
	return UpdateSMBSecurityStrategyRequest{Request: req, Input: input, Copy: c.UpdateSMBSecurityStrategyRequest}
}

// UpdateSMBSecurityStrategyRequest is the request type for the
// UpdateSMBSecurityStrategy API operation.
type UpdateSMBSecurityStrategyRequest struct {
	*aws.Request
	Input *types.UpdateSMBSecurityStrategyInput
	Copy  func(*types.UpdateSMBSecurityStrategyInput) UpdateSMBSecurityStrategyRequest
}

// Send marshals and sends the UpdateSMBSecurityStrategy API request.
func (r UpdateSMBSecurityStrategyRequest) Send(ctx context.Context) (*UpdateSMBSecurityStrategyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSMBSecurityStrategyResponse{
		UpdateSMBSecurityStrategyOutput: r.Request.Data.(*types.UpdateSMBSecurityStrategyOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSMBSecurityStrategyResponse is the response type for the
// UpdateSMBSecurityStrategy API operation.
type UpdateSMBSecurityStrategyResponse struct {
	*types.UpdateSMBSecurityStrategyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSMBSecurityStrategy request.
func (r *UpdateSMBSecurityStrategyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
