// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDeletePolicy = "DeletePolicy"

// DeletePolicyRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes the specified policy.
//
// A policy cannot be deleted if it has non-default versions or it is attached
// to any certificate.
//
// To delete a policy, use the DeletePolicyVersion API to delete all non-default
// versions of the policy; use the DetachPrincipalPolicy API to detach the policy
// from any certificate; and then use the DeletePolicy API to delete the policy.
//
// When a policy is deleted using DeletePolicy, its default version is deleted
// with it.
//
//    // Example sending a request using DeletePolicyRequest.
//    req := client.DeletePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeletePolicyRequest(input *types.DeletePolicyInput) DeletePolicyRequest {
	op := &aws.Operation{
		Name:       opDeletePolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/policies/{policyName}",
	}

	if input == nil {
		input = &types.DeletePolicyInput{}
	}

	req := c.newRequest(op, input, &types.DeletePolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeletePolicyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePolicyRequest{Request: req, Input: input, Copy: c.DeletePolicyRequest}
}

// DeletePolicyRequest is the request type for the
// DeletePolicy API operation.
type DeletePolicyRequest struct {
	*aws.Request
	Input *types.DeletePolicyInput
	Copy  func(*types.DeletePolicyInput) DeletePolicyRequest
}

// Send marshals and sends the DeletePolicy API request.
func (r DeletePolicyRequest) Send(ctx context.Context) (*DeletePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePolicyResponse{
		DeletePolicyOutput: r.Request.Data.(*types.DeletePolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePolicyResponse is the response type for the
// DeletePolicy API operation.
type DeletePolicyResponse struct {
	*types.DeletePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePolicy request.
func (r *DeletePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
