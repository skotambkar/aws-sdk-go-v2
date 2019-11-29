// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opGetRuleGroup = "GetRuleGroup"

// GetRuleGroupRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the RuleGroup that is specified by the RuleGroupId that you included
// in the GetRuleGroup request.
//
// To view the rules in a rule group, use ListActivatedRulesInRuleGroup.
//
//    // Example sending a request using GetRuleGroupRequest.
//    req := client.GetRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRuleGroup
func (c *Client) GetRuleGroupRequest(input *types.GetRuleGroupInput) GetRuleGroupRequest {
	op := &aws.Operation{
		Name:       opGetRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRuleGroupInput{}
	}

	req := c.newRequest(op, input, &types.GetRuleGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetRuleGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRuleGroupRequest{Request: req, Input: input, Copy: c.GetRuleGroupRequest}
}

// GetRuleGroupRequest is the request type for the
// GetRuleGroup API operation.
type GetRuleGroupRequest struct {
	*aws.Request
	Input *types.GetRuleGroupInput
	Copy  func(*types.GetRuleGroupInput) GetRuleGroupRequest
}

// Send marshals and sends the GetRuleGroup API request.
func (r GetRuleGroupRequest) Send(ctx context.Context) (*GetRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRuleGroupResponse{
		GetRuleGroupOutput: r.Request.Data.(*types.GetRuleGroupOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRuleGroupResponse is the response type for the
// GetRuleGroup API operation.
type GetRuleGroupResponse struct {
	*types.GetRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRuleGroup request.
func (r *GetRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
