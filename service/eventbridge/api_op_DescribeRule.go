// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

const opDescribeRule = "DescribeRule"

// DescribeRuleRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Describes the specified rule.
//
// DescribeRule doesn't list the targets of a rule. To see the targets associated
// with a rule, use ListTargetsByRule.
//
//    // Example sending a request using DescribeRuleRequest.
//    req := client.DescribeRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/DescribeRule
func (c *Client) DescribeRuleRequest(input *types.DescribeRuleInput) DescribeRuleRequest {
	op := &aws.Operation{
		Name:       opDescribeRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeRuleInput{}
	}

	req := c.newRequest(op, input, &types.DescribeRuleOutput{})
	return DescribeRuleRequest{Request: req, Input: input, Copy: c.DescribeRuleRequest}
}

// DescribeRuleRequest is the request type for the
// DescribeRule API operation.
type DescribeRuleRequest struct {
	*aws.Request
	Input *types.DescribeRuleInput
	Copy  func(*types.DescribeRuleInput) DescribeRuleRequest
}

// Send marshals and sends the DescribeRule API request.
func (r DescribeRuleRequest) Send(ctx context.Context) (*DescribeRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRuleResponse{
		DescribeRuleOutput: r.Request.Data.(*types.DescribeRuleOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRuleResponse is the response type for the
// DescribeRule API operation.
type DescribeRuleResponse struct {
	*types.DescribeRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRule request.
func (r *DescribeRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
