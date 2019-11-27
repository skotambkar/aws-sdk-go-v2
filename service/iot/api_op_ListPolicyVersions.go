// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListPolicyVersions = "ListPolicyVersions"

// ListPolicyVersionsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the versions of the specified policy and identifies the default version.
//
//    // Example sending a request using ListPolicyVersionsRequest.
//    req := client.ListPolicyVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListPolicyVersionsRequest(input *types.ListPolicyVersionsInput) ListPolicyVersionsRequest {
	op := &aws.Operation{
		Name:       opListPolicyVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/policies/{policyName}/version",
	}

	if input == nil {
		input = &types.ListPolicyVersionsInput{}
	}

	req := c.newRequest(op, input, &types.ListPolicyVersionsOutput{})
	return ListPolicyVersionsRequest{Request: req, Input: input, Copy: c.ListPolicyVersionsRequest}
}

// ListPolicyVersionsRequest is the request type for the
// ListPolicyVersions API operation.
type ListPolicyVersionsRequest struct {
	*aws.Request
	Input *types.ListPolicyVersionsInput
	Copy  func(*types.ListPolicyVersionsInput) ListPolicyVersionsRequest
}

// Send marshals and sends the ListPolicyVersions API request.
func (r ListPolicyVersionsRequest) Send(ctx context.Context) (*ListPolicyVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPolicyVersionsResponse{
		ListPolicyVersionsOutput: r.Request.Data.(*types.ListPolicyVersionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPolicyVersionsResponse is the response type for the
// ListPolicyVersions API operation.
type ListPolicyVersionsResponse struct {
	*types.ListPolicyVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicyVersions request.
func (r *ListPolicyVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
