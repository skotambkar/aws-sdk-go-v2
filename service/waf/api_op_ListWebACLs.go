// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opListWebACLs = "ListWebACLs"

// ListWebACLsRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns an array of WebACLSummary objects in the response.
//
//    // Example sending a request using ListWebACLsRequest.
//    req := client.ListWebACLsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/ListWebACLs
func (c *Client) ListWebACLsRequest(input *types.ListWebACLsInput) ListWebACLsRequest {
	op := &aws.Operation{
		Name:       opListWebACLs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListWebACLsInput{}
	}

	req := c.newRequest(op, input, &types.ListWebACLsOutput{})
	return ListWebACLsRequest{Request: req, Input: input, Copy: c.ListWebACLsRequest}
}

// ListWebACLsRequest is the request type for the
// ListWebACLs API operation.
type ListWebACLsRequest struct {
	*aws.Request
	Input *types.ListWebACLsInput
	Copy  func(*types.ListWebACLsInput) ListWebACLsRequest
}

// Send marshals and sends the ListWebACLs API request.
func (r ListWebACLsRequest) Send(ctx context.Context) (*ListWebACLsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListWebACLsResponse{
		ListWebACLsOutput: r.Request.Data.(*types.ListWebACLsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListWebACLsResponse is the response type for the
// ListWebACLs API operation.
type ListWebACLsResponse struct {
	*types.ListWebACLsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListWebACLs request.
func (r *ListWebACLsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
