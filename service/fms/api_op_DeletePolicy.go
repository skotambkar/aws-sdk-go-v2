// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
)

const opDeletePolicy = "DeletePolicy"

// DeletePolicyRequest returns a request value for making API operation for
// Firewall Management Service.
//
// Permanently deletes an AWS Firewall Manager policy.
//
//    // Example sending a request using DeletePolicyRequest.
//    req := client.DeletePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/DeletePolicy
func (c *Client) DeletePolicyRequest(input *types.DeletePolicyInput) DeletePolicyRequest {
	op := &aws.Operation{
		Name:       opDeletePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeletePolicyInput{}
	}

	req := c.newRequest(op, input, &types.DeletePolicyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
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
