// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opAttachRolePolicy = "AttachRolePolicy"

// AttachRolePolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Attaches the specified managed policy to the specified IAM role. When you
// attach a managed policy to a role, the managed policy becomes part of the
// role's permission (access) policy.
//
// You cannot use a managed policy as the role's trust policy. The role's trust
// policy is created at the same time as the role, using CreateRole. You can
// update a role's trust policy using UpdateAssumeRolePolicy.
//
// Use this API to attach a managed policy to a role. To embed an inline policy
// in a role, use PutRolePolicy. For more information about policies, see Managed
// Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using AttachRolePolicyRequest.
//    req := client.AttachRolePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/AttachRolePolicy
func (c *Client) AttachRolePolicyRequest(input *types.AttachRolePolicyInput) AttachRolePolicyRequest {
	op := &aws.Operation{
		Name:       opAttachRolePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AttachRolePolicyInput{}
	}

	req := c.newRequest(op, input, &types.AttachRolePolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AttachRolePolicyRequest{Request: req, Input: input, Copy: c.AttachRolePolicyRequest}
}

// AttachRolePolicyRequest is the request type for the
// AttachRolePolicy API operation.
type AttachRolePolicyRequest struct {
	*aws.Request
	Input *types.AttachRolePolicyInput
	Copy  func(*types.AttachRolePolicyInput) AttachRolePolicyRequest
}

// Send marshals and sends the AttachRolePolicy API request.
func (r AttachRolePolicyRequest) Send(ctx context.Context) (*AttachRolePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachRolePolicyResponse{
		AttachRolePolicyOutput: r.Request.Data.(*types.AttachRolePolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachRolePolicyResponse is the response type for the
// AttachRolePolicy API operation.
type AttachRolePolicyResponse struct {
	*types.AttachRolePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachRolePolicy request.
func (r *AttachRolePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
