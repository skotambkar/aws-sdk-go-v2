// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opGetPolicy = "GetPolicy"

// GetPolicyRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the specified policy with the policy document of the
// default version.
//
//    // Example sending a request using GetPolicyRequest.
//    req := client.GetPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetPolicyRequest(input *types.GetPolicyInput) GetPolicyRequest {
	op := &aws.Operation{
		Name:       opGetPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/policies/{policyName}",
	}

	if input == nil {
		input = &types.GetPolicyInput{}
	}

	req := c.newRequest(op, input, &types.GetPolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetPolicyMarshaler{Input: input}.GetNamedBuildHandler())

	return GetPolicyRequest{Request: req, Input: input, Copy: c.GetPolicyRequest}
}

// GetPolicyRequest is the request type for the
// GetPolicy API operation.
type GetPolicyRequest struct {
	*aws.Request
	Input *types.GetPolicyInput
	Copy  func(*types.GetPolicyInput) GetPolicyRequest
}

// Send marshals and sends the GetPolicy API request.
func (r GetPolicyRequest) Send(ctx context.Context) (*GetPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPolicyResponse{
		GetPolicyOutput: r.Request.Data.(*types.GetPolicyOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPolicyResponse is the response type for the
// GetPolicy API operation.
type GetPolicyResponse struct {
	*types.GetPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPolicy request.
func (r *GetPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
