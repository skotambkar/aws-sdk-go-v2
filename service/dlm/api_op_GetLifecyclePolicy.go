// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dlm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/dlm/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
)

const opGetLifecyclePolicy = "GetLifecyclePolicy"

// GetLifecyclePolicyRequest returns a request value for making API operation for
// Amazon Data Lifecycle Manager.
//
// Gets detailed information about the specified lifecycle policy.
//
//    // Example sending a request using GetLifecyclePolicyRequest.
//    req := client.GetLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dlm-2018-01-12/GetLifecyclePolicy
func (c *Client) GetLifecyclePolicyRequest(input *types.GetLifecyclePolicyInput) GetLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opGetLifecyclePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/policies/{policyId}/",
	}

	if input == nil {
		input = &types.GetLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &types.GetLifecyclePolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetLifecyclePolicyMarshaler{Input: input}.GetNamedBuildHandler())

	return GetLifecyclePolicyRequest{Request: req, Input: input, Copy: c.GetLifecyclePolicyRequest}
}

// GetLifecyclePolicyRequest is the request type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyRequest struct {
	*aws.Request
	Input *types.GetLifecyclePolicyInput
	Copy  func(*types.GetLifecyclePolicyInput) GetLifecyclePolicyRequest
}

// Send marshals and sends the GetLifecyclePolicy API request.
func (r GetLifecyclePolicyRequest) Send(ctx context.Context) (*GetLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLifecyclePolicyResponse{
		GetLifecyclePolicyOutput: r.Request.Data.(*types.GetLifecyclePolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLifecyclePolicyResponse is the response type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyResponse struct {
	*types.GetLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLifecyclePolicy request.
func (r *GetLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
