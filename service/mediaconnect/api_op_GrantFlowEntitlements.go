// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
)

const opGrantFlowEntitlements = "GrantFlowEntitlements"

// GrantFlowEntitlementsRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Grants entitlements to an existing flow.
//
//    // Example sending a request using GrantFlowEntitlementsRequest.
//    req := client.GrantFlowEntitlementsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/GrantFlowEntitlements
func (c *Client) GrantFlowEntitlementsRequest(input *types.GrantFlowEntitlementsInput) GrantFlowEntitlementsRequest {
	op := &aws.Operation{
		Name:       opGrantFlowEntitlements,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/flows/{flowArn}/entitlements",
	}

	if input == nil {
		input = &types.GrantFlowEntitlementsInput{}
	}

	req := c.newRequest(op, input, &types.GrantFlowEntitlementsOutput{})
	return GrantFlowEntitlementsRequest{Request: req, Input: input, Copy: c.GrantFlowEntitlementsRequest}
}

// GrantFlowEntitlementsRequest is the request type for the
// GrantFlowEntitlements API operation.
type GrantFlowEntitlementsRequest struct {
	*aws.Request
	Input *types.GrantFlowEntitlementsInput
	Copy  func(*types.GrantFlowEntitlementsInput) GrantFlowEntitlementsRequest
}

// Send marshals and sends the GrantFlowEntitlements API request.
func (r GrantFlowEntitlementsRequest) Send(ctx context.Context) (*GrantFlowEntitlementsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GrantFlowEntitlementsResponse{
		GrantFlowEntitlementsOutput: r.Request.Data.(*types.GrantFlowEntitlementsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GrantFlowEntitlementsResponse is the response type for the
// GrantFlowEntitlements API operation.
type GrantFlowEntitlementsResponse struct {
	*types.GrantFlowEntitlementsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GrantFlowEntitlements request.
func (r *GrantFlowEntitlementsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
