// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opListByteMatchSets = "ListByteMatchSets"

// ListByteMatchSetsRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns an array of ByteMatchSetSummary objects.
//
//    // Example sending a request using ListByteMatchSetsRequest.
//    req := client.ListByteMatchSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListByteMatchSets
func (c *Client) ListByteMatchSetsRequest(input *types.ListByteMatchSetsInput) ListByteMatchSetsRequest {
	op := &aws.Operation{
		Name:       opListByteMatchSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListByteMatchSetsInput{}
	}

	req := c.newRequest(op, input, &types.ListByteMatchSetsOutput{})
	return ListByteMatchSetsRequest{Request: req, Input: input, Copy: c.ListByteMatchSetsRequest}
}

// ListByteMatchSetsRequest is the request type for the
// ListByteMatchSets API operation.
type ListByteMatchSetsRequest struct {
	*aws.Request
	Input *types.ListByteMatchSetsInput
	Copy  func(*types.ListByteMatchSetsInput) ListByteMatchSetsRequest
}

// Send marshals and sends the ListByteMatchSets API request.
func (r ListByteMatchSetsRequest) Send(ctx context.Context) (*ListByteMatchSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListByteMatchSetsResponse{
		ListByteMatchSetsOutput: r.Request.Data.(*types.ListByteMatchSetsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListByteMatchSetsResponse is the response type for the
// ListByteMatchSets API operation.
type ListByteMatchSetsResponse struct {
	*types.ListByteMatchSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListByteMatchSets request.
func (r *ListByteMatchSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
