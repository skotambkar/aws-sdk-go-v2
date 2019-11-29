// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opListRegexPatternSets = "ListRegexPatternSets"

// ListRegexPatternSetsRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns an array of RegexPatternSetSummary objects.
//
//    // Example sending a request using ListRegexPatternSetsRequest.
//    req := client.ListRegexPatternSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/ListRegexPatternSets
func (c *Client) ListRegexPatternSetsRequest(input *types.ListRegexPatternSetsInput) ListRegexPatternSetsRequest {
	op := &aws.Operation{
		Name:       opListRegexPatternSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListRegexPatternSetsInput{}
	}

	req := c.newRequest(op, input, &types.ListRegexPatternSetsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListRegexPatternSetsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListRegexPatternSetsRequest{Request: req, Input: input, Copy: c.ListRegexPatternSetsRequest}
}

// ListRegexPatternSetsRequest is the request type for the
// ListRegexPatternSets API operation.
type ListRegexPatternSetsRequest struct {
	*aws.Request
	Input *types.ListRegexPatternSetsInput
	Copy  func(*types.ListRegexPatternSetsInput) ListRegexPatternSetsRequest
}

// Send marshals and sends the ListRegexPatternSets API request.
func (r ListRegexPatternSetsRequest) Send(ctx context.Context) (*ListRegexPatternSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRegexPatternSetsResponse{
		ListRegexPatternSetsOutput: r.Request.Data.(*types.ListRegexPatternSetsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRegexPatternSetsResponse is the response type for the
// ListRegexPatternSets API operation.
type ListRegexPatternSetsResponse struct {
	*types.ListRegexPatternSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRegexPatternSets request.
func (r *ListRegexPatternSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
