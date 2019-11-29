// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opListRules = "ListRules"

// ListRulesRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns an array of RuleSummary objects.
//
//    // Example sending a request using ListRulesRequest.
//    req := client.ListRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListRules
func (c *Client) ListRulesRequest(input *types.ListRulesInput) ListRulesRequest {
	op := &aws.Operation{
		Name:       opListRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListRulesInput{}
	}

	req := c.newRequest(op, input, &types.ListRulesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListRulesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListRulesRequest{Request: req, Input: input, Copy: c.ListRulesRequest}
}

// ListRulesRequest is the request type for the
// ListRules API operation.
type ListRulesRequest struct {
	*aws.Request
	Input *types.ListRulesInput
	Copy  func(*types.ListRulesInput) ListRulesRequest
}

// Send marshals and sends the ListRules API request.
func (r ListRulesRequest) Send(ctx context.Context) (*ListRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRulesResponse{
		ListRulesOutput: r.Request.Data.(*types.ListRulesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRulesResponse is the response type for the
// ListRules API operation.
type ListRulesResponse struct {
	*types.ListRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRules request.
func (r *ListRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
