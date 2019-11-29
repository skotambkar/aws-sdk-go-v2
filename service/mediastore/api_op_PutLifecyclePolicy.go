// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
)

const opPutLifecyclePolicy = "PutLifecyclePolicy"

// PutLifecyclePolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Writes an object lifecycle policy to a container. If the container already
// has an object lifecycle policy, the service replaces the existing policy
// with the new policy. It takes up to 20 minutes for the change to take effect.
//
// For information about how to construct an object lifecycle policy, see Components
// of an Object Lifecycle Policy (https://docs.aws.amazon.com/mediastore/latest/ug/policies-object-lifecycle-components.html).
//
//    // Example sending a request using PutLifecyclePolicyRequest.
//    req := client.PutLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutLifecyclePolicy
func (c *Client) PutLifecyclePolicyRequest(input *types.PutLifecyclePolicyInput) PutLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opPutLifecyclePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &types.PutLifecyclePolicyOutput{})
	return PutLifecyclePolicyRequest{Request: req, Input: input, Copy: c.PutLifecyclePolicyRequest}
}

// PutLifecyclePolicyRequest is the request type for the
// PutLifecyclePolicy API operation.
type PutLifecyclePolicyRequest struct {
	*aws.Request
	Input *types.PutLifecyclePolicyInput
	Copy  func(*types.PutLifecyclePolicyInput) PutLifecyclePolicyRequest
}

// Send marshals and sends the PutLifecyclePolicy API request.
func (r PutLifecyclePolicyRequest) Send(ctx context.Context) (*PutLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLifecyclePolicyResponse{
		PutLifecyclePolicyOutput: r.Request.Data.(*types.PutLifecyclePolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLifecyclePolicyResponse is the response type for the
// PutLifecyclePolicy API operation.
type PutLifecyclePolicyResponse struct {
	*types.PutLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLifecyclePolicy request.
func (r *PutLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
