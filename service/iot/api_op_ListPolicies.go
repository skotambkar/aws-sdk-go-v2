// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListPolicies = "ListPolicies"

// ListPoliciesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists your policies.
//
//    // Example sending a request using ListPoliciesRequest.
//    req := client.ListPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListPoliciesRequest(input *types.ListPoliciesInput) ListPoliciesRequest {
	op := &aws.Operation{
		Name:       opListPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/policies",
	}

	if input == nil {
		input = &types.ListPoliciesInput{}
	}

	req := c.newRequest(op, input, &types.ListPoliciesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListPoliciesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPoliciesRequest{Request: req, Input: input, Copy: c.ListPoliciesRequest}
}

// ListPoliciesRequest is the request type for the
// ListPolicies API operation.
type ListPoliciesRequest struct {
	*aws.Request
	Input *types.ListPoliciesInput
	Copy  func(*types.ListPoliciesInput) ListPoliciesRequest
}

// Send marshals and sends the ListPolicies API request.
func (r ListPoliciesRequest) Send(ctx context.Context) (*ListPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPoliciesResponse{
		ListPoliciesOutput: r.Request.Data.(*types.ListPoliciesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPoliciesResponse is the response type for the
// ListPolicies API operation.
type ListPoliciesResponse struct {
	*types.ListPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicies request.
func (r *ListPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
