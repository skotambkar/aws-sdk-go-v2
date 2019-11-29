// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opGetWebACL = "GetWebACL"

// GetWebACLRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the WebACL that is specified by WebACLId.
//
//    // Example sending a request using GetWebACLRequest.
//    req := client.GetWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetWebACL
func (c *Client) GetWebACLRequest(input *types.GetWebACLInput) GetWebACLRequest {
	op := &aws.Operation{
		Name:       opGetWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetWebACLInput{}
	}

	req := c.newRequest(op, input, &types.GetWebACLOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetWebACLMarshaler{Input: input}.GetNamedBuildHandler())

	return GetWebACLRequest{Request: req, Input: input, Copy: c.GetWebACLRequest}
}

// GetWebACLRequest is the request type for the
// GetWebACL API operation.
type GetWebACLRequest struct {
	*aws.Request
	Input *types.GetWebACLInput
	Copy  func(*types.GetWebACLInput) GetWebACLRequest
}

// Send marshals and sends the GetWebACL API request.
func (r GetWebACLRequest) Send(ctx context.Context) (*GetWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWebACLResponse{
		GetWebACLOutput: r.Request.Data.(*types.GetWebACLOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetWebACLResponse is the response type for the
// GetWebACL API operation.
type GetWebACLResponse struct {
	*types.GetWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWebACL request.
func (r *GetWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
