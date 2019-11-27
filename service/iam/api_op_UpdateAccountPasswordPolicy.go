// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opUpdateAccountPasswordPolicy = "UpdateAccountPasswordPolicy"

// UpdateAccountPasswordPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Updates the password policy settings for the AWS account.
//
//    * This operation does not support partial updates. No parameters are required,
//    but if you do not specify a parameter, that parameter's value reverts
//    to its default value. See the Request Parameters section for each parameter's
//    default value. Also note that some parameters do not allow the default
//    parameter to be explicitly set. Instead, to invoke the default value,
//    do not include that parameter when you invoke the operation.
//
// For more information about using a password policy, see Managing an IAM Password
// Policy (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingPasswordPolicies.html)
// in the IAM User Guide.
//
//    // Example sending a request using UpdateAccountPasswordPolicyRequest.
//    req := client.UpdateAccountPasswordPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateAccountPasswordPolicy
func (c *Client) UpdateAccountPasswordPolicyRequest(input *types.UpdateAccountPasswordPolicyInput) UpdateAccountPasswordPolicyRequest {
	op := &aws.Operation{
		Name:       opUpdateAccountPasswordPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateAccountPasswordPolicyInput{}
	}

	req := c.newRequest(op, input, &types.UpdateAccountPasswordPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateAccountPasswordPolicyRequest{Request: req, Input: input, Copy: c.UpdateAccountPasswordPolicyRequest}
}

// UpdateAccountPasswordPolicyRequest is the request type for the
// UpdateAccountPasswordPolicy API operation.
type UpdateAccountPasswordPolicyRequest struct {
	*aws.Request
	Input *types.UpdateAccountPasswordPolicyInput
	Copy  func(*types.UpdateAccountPasswordPolicyInput) UpdateAccountPasswordPolicyRequest
}

// Send marshals and sends the UpdateAccountPasswordPolicy API request.
func (r UpdateAccountPasswordPolicyRequest) Send(ctx context.Context) (*UpdateAccountPasswordPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccountPasswordPolicyResponse{
		UpdateAccountPasswordPolicyOutput: r.Request.Data.(*types.UpdateAccountPasswordPolicyOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccountPasswordPolicyResponse is the response type for the
// UpdateAccountPasswordPolicy API operation.
type UpdateAccountPasswordPolicyResponse struct {
	*types.UpdateAccountPasswordPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccountPasswordPolicy request.
func (r *UpdateAccountPasswordPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
