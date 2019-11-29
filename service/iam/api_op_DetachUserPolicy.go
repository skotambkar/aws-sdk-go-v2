// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opDetachUserPolicy = "DetachUserPolicy"

// DetachUserPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Removes the specified managed policy from the specified user.
//
// A user can also have inline policies embedded with it. To delete an inline
// policy, use the DeleteUserPolicy API. For information about policies, see
// Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using DetachUserPolicyRequest.
//    req := client.DetachUserPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DetachUserPolicy
func (c *Client) DetachUserPolicyRequest(input *types.DetachUserPolicyInput) DetachUserPolicyRequest {
	op := &aws.Operation{
		Name:       opDetachUserPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DetachUserPolicyInput{}
	}

	req := c.newRequest(op, input, &types.DetachUserPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DetachUserPolicyRequest{Request: req, Input: input, Copy: c.DetachUserPolicyRequest}
}

// DetachUserPolicyRequest is the request type for the
// DetachUserPolicy API operation.
type DetachUserPolicyRequest struct {
	*aws.Request
	Input *types.DetachUserPolicyInput
	Copy  func(*types.DetachUserPolicyInput) DetachUserPolicyRequest
}

// Send marshals and sends the DetachUserPolicy API request.
func (r DetachUserPolicyRequest) Send(ctx context.Context) (*DetachUserPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachUserPolicyResponse{
		DetachUserPolicyOutput: r.Request.Data.(*types.DetachUserPolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachUserPolicyResponse is the response type for the
// DetachUserPolicy API operation.
type DetachUserPolicyResponse struct {
	*types.DetachUserPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachUserPolicy request.
func (r *DetachUserPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
