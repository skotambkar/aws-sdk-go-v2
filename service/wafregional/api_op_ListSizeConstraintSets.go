// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opListSizeConstraintSets = "ListSizeConstraintSets"

// ListSizeConstraintSetsRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns an array of SizeConstraintSetSummary objects.
//
//    // Example sending a request using ListSizeConstraintSetsRequest.
//    req := client.ListSizeConstraintSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListSizeConstraintSets
func (c *Client) ListSizeConstraintSetsRequest(input *types.ListSizeConstraintSetsInput) ListSizeConstraintSetsRequest {
	op := &aws.Operation{
		Name:       opListSizeConstraintSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListSizeConstraintSetsInput{}
	}

	req := c.newRequest(op, input, &types.ListSizeConstraintSetsOutput{})
	return ListSizeConstraintSetsRequest{Request: req, Input: input, Copy: c.ListSizeConstraintSetsRequest}
}

// ListSizeConstraintSetsRequest is the request type for the
// ListSizeConstraintSets API operation.
type ListSizeConstraintSetsRequest struct {
	*aws.Request
	Input *types.ListSizeConstraintSetsInput
	Copy  func(*types.ListSizeConstraintSetsInput) ListSizeConstraintSetsRequest
}

// Send marshals and sends the ListSizeConstraintSets API request.
func (r ListSizeConstraintSetsRequest) Send(ctx context.Context) (*ListSizeConstraintSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSizeConstraintSetsResponse{
		ListSizeConstraintSetsOutput: r.Request.Data.(*types.ListSizeConstraintSetsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSizeConstraintSetsResponse is the response type for the
// ListSizeConstraintSets API operation.
type ListSizeConstraintSetsResponse struct {
	*types.ListSizeConstraintSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSizeConstraintSets request.
func (r *ListSizeConstraintSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
