// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opListRegexMatchSets = "ListRegexMatchSets"

// ListRegexMatchSetsRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns an array of RegexMatchSetSummary objects.
//
//    // Example sending a request using ListRegexMatchSetsRequest.
//    req := client.ListRegexMatchSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/ListRegexMatchSets
func (c *Client) ListRegexMatchSetsRequest(input *types.ListRegexMatchSetsInput) ListRegexMatchSetsRequest {
	op := &aws.Operation{
		Name:       opListRegexMatchSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListRegexMatchSetsInput{}
	}

	req := c.newRequest(op, input, &types.ListRegexMatchSetsOutput{})
	return ListRegexMatchSetsRequest{Request: req, Input: input, Copy: c.ListRegexMatchSetsRequest}
}

// ListRegexMatchSetsRequest is the request type for the
// ListRegexMatchSets API operation.
type ListRegexMatchSetsRequest struct {
	*aws.Request
	Input *types.ListRegexMatchSetsInput
	Copy  func(*types.ListRegexMatchSetsInput) ListRegexMatchSetsRequest
}

// Send marshals and sends the ListRegexMatchSets API request.
func (r ListRegexMatchSetsRequest) Send(ctx context.Context) (*ListRegexMatchSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRegexMatchSetsResponse{
		ListRegexMatchSetsOutput: r.Request.Data.(*types.ListRegexMatchSetsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRegexMatchSetsResponse is the response type for the
// ListRegexMatchSets API operation.
type ListRegexMatchSetsResponse struct {
	*types.ListRegexMatchSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRegexMatchSets request.
func (r *ListRegexMatchSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
