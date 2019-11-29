// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opGetRegexMatchSet = "GetRegexMatchSet"

// GetRegexMatchSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns the RegexMatchSet specified by RegexMatchSetId.
//
//    // Example sending a request using GetRegexMatchSetRequest.
//    req := client.GetRegexMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetRegexMatchSet
func (c *Client) GetRegexMatchSetRequest(input *types.GetRegexMatchSetInput) GetRegexMatchSetRequest {
	op := &aws.Operation{
		Name:       opGetRegexMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRegexMatchSetInput{}
	}

	req := c.newRequest(op, input, &types.GetRegexMatchSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetRegexMatchSetMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRegexMatchSetRequest{Request: req, Input: input, Copy: c.GetRegexMatchSetRequest}
}

// GetRegexMatchSetRequest is the request type for the
// GetRegexMatchSet API operation.
type GetRegexMatchSetRequest struct {
	*aws.Request
	Input *types.GetRegexMatchSetInput
	Copy  func(*types.GetRegexMatchSetInput) GetRegexMatchSetRequest
}

// Send marshals and sends the GetRegexMatchSet API request.
func (r GetRegexMatchSetRequest) Send(ctx context.Context) (*GetRegexMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRegexMatchSetResponse{
		GetRegexMatchSetOutput: r.Request.Data.(*types.GetRegexMatchSetOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRegexMatchSetResponse is the response type for the
// GetRegexMatchSet API operation.
type GetRegexMatchSetResponse struct {
	*types.GetRegexMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRegexMatchSet request.
func (r *GetRegexMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
