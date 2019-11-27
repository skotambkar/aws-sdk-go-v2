// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyTrafficMirrorFilterRule = "ModifyTrafficMirrorFilterRule"

// ModifyTrafficMirrorFilterRuleRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified Traffic Mirror rule.
//
// DestinationCidrBlock and SourceCidrBlock must both be an IPv4 range or an
// IPv6 range.
//
//    // Example sending a request using ModifyTrafficMirrorFilterRuleRequest.
//    req := client.ModifyTrafficMirrorFilterRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTrafficMirrorFilterRule
func (c *Client) ModifyTrafficMirrorFilterRuleRequest(input *types.ModifyTrafficMirrorFilterRuleInput) ModifyTrafficMirrorFilterRuleRequest {
	op := &aws.Operation{
		Name:       opModifyTrafficMirrorFilterRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyTrafficMirrorFilterRuleInput{}
	}

	req := c.newRequest(op, input, &types.ModifyTrafficMirrorFilterRuleOutput{})
	return ModifyTrafficMirrorFilterRuleRequest{Request: req, Input: input, Copy: c.ModifyTrafficMirrorFilterRuleRequest}
}

// ModifyTrafficMirrorFilterRuleRequest is the request type for the
// ModifyTrafficMirrorFilterRule API operation.
type ModifyTrafficMirrorFilterRuleRequest struct {
	*aws.Request
	Input *types.ModifyTrafficMirrorFilterRuleInput
	Copy  func(*types.ModifyTrafficMirrorFilterRuleInput) ModifyTrafficMirrorFilterRuleRequest
}

// Send marshals and sends the ModifyTrafficMirrorFilterRule API request.
func (r ModifyTrafficMirrorFilterRuleRequest) Send(ctx context.Context) (*ModifyTrafficMirrorFilterRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyTrafficMirrorFilterRuleResponse{
		ModifyTrafficMirrorFilterRuleOutput: r.Request.Data.(*types.ModifyTrafficMirrorFilterRuleOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyTrafficMirrorFilterRuleResponse is the response type for the
// ModifyTrafficMirrorFilterRule API operation.
type ModifyTrafficMirrorFilterRuleResponse struct {
	*types.ModifyTrafficMirrorFilterRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyTrafficMirrorFilterRule request.
func (r *ModifyTrafficMirrorFilterRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
