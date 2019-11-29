// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opGetRateBasedRule = "GetRateBasedRule"

// GetRateBasedRuleRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the RateBasedRule that is specified by the RuleId that you included
// in the GetRateBasedRule request.
//
//    // Example sending a request using GetRateBasedRuleRequest.
//    req := client.GetRateBasedRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRateBasedRule
func (c *Client) GetRateBasedRuleRequest(input *types.GetRateBasedRuleInput) GetRateBasedRuleRequest {
	op := &aws.Operation{
		Name:       opGetRateBasedRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRateBasedRuleInput{}
	}

	req := c.newRequest(op, input, &types.GetRateBasedRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetRateBasedRuleMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRateBasedRuleRequest{Request: req, Input: input, Copy: c.GetRateBasedRuleRequest}
}

// GetRateBasedRuleRequest is the request type for the
// GetRateBasedRule API operation.
type GetRateBasedRuleRequest struct {
	*aws.Request
	Input *types.GetRateBasedRuleInput
	Copy  func(*types.GetRateBasedRuleInput) GetRateBasedRuleRequest
}

// Send marshals and sends the GetRateBasedRule API request.
func (r GetRateBasedRuleRequest) Send(ctx context.Context) (*GetRateBasedRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRateBasedRuleResponse{
		GetRateBasedRuleOutput: r.Request.Data.(*types.GetRateBasedRuleOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRateBasedRuleResponse is the response type for the
// GetRateBasedRule API operation.
type GetRateBasedRuleResponse struct {
	*types.GetRateBasedRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRateBasedRule request.
func (r *GetRateBasedRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
