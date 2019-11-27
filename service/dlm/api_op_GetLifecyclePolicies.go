// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dlm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
)

const opGetLifecyclePolicies = "GetLifecyclePolicies"

// GetLifecyclePoliciesRequest returns a request value for making API operation for
// Amazon Data Lifecycle Manager.
//
// Gets summary information about all or the specified data lifecycle policies.
//
// To get complete information about a policy, use GetLifecyclePolicy.
//
//    // Example sending a request using GetLifecyclePoliciesRequest.
//    req := client.GetLifecyclePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dlm-2018-01-12/GetLifecyclePolicies
func (c *Client) GetLifecyclePoliciesRequest(input *types.GetLifecyclePoliciesInput) GetLifecyclePoliciesRequest {
	op := &aws.Operation{
		Name:       opGetLifecyclePolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/policies",
	}

	if input == nil {
		input = &types.GetLifecyclePoliciesInput{}
	}

	req := c.newRequest(op, input, &types.GetLifecyclePoliciesOutput{})
	return GetLifecyclePoliciesRequest{Request: req, Input: input, Copy: c.GetLifecyclePoliciesRequest}
}

// GetLifecyclePoliciesRequest is the request type for the
// GetLifecyclePolicies API operation.
type GetLifecyclePoliciesRequest struct {
	*aws.Request
	Input *types.GetLifecyclePoliciesInput
	Copy  func(*types.GetLifecyclePoliciesInput) GetLifecyclePoliciesRequest
}

// Send marshals and sends the GetLifecyclePolicies API request.
func (r GetLifecyclePoliciesRequest) Send(ctx context.Context) (*GetLifecyclePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLifecyclePoliciesResponse{
		GetLifecyclePoliciesOutput: r.Request.Data.(*types.GetLifecyclePoliciesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLifecyclePoliciesResponse is the response type for the
// GetLifecyclePolicies API operation.
type GetLifecyclePoliciesResponse struct {
	*types.GetLifecyclePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLifecyclePolicies request.
func (r *GetLifecyclePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
