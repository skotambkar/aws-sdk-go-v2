// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

const opCreateRule = "CreateRule"

// CreateRuleRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates a rule for the specified listener. The listener must be associated
// with an Application Load Balancer.
//
// Rules are evaluated in priority order, from the lowest value to the highest
// value. When the conditions for a rule are met, its actions are performed.
// If the conditions for no rules are met, the actions for the default rule
// are performed. For more information, see Listener Rules (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules)
// in the Application Load Balancers Guide.
//
// To view your current rules, use DescribeRules. To update a rule, use ModifyRule.
// To set the priorities of your rules, use SetRulePriorities. To delete a rule,
// use DeleteRule.
//
//    // Example sending a request using CreateRuleRequest.
//    req := client.CreateRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/CreateRule
func (c *Client) CreateRuleRequest(input *types.CreateRuleInput) CreateRuleRequest {
	op := &aws.Operation{
		Name:       opCreateRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateRuleInput{}
	}

	req := c.newRequest(op, input, &types.CreateRuleOutput{})
	return CreateRuleRequest{Request: req, Input: input, Copy: c.CreateRuleRequest}
}

// CreateRuleRequest is the request type for the
// CreateRule API operation.
type CreateRuleRequest struct {
	*aws.Request
	Input *types.CreateRuleInput
	Copy  func(*types.CreateRuleInput) CreateRuleRequest
}

// Send marshals and sends the CreateRule API request.
func (r CreateRuleRequest) Send(ctx context.Context) (*CreateRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRuleResponse{
		CreateRuleOutput: r.Request.Data.(*types.CreateRuleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRuleResponse is the response type for the
// CreateRule API operation.
type CreateRuleResponse struct {
	*types.CreateRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRule request.
func (r *CreateRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
