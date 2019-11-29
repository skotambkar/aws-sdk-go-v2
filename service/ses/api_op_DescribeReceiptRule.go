// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opDescribeReceiptRule = "DescribeReceiptRule"

// DescribeReceiptRuleRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns the details of the specified receipt rule.
//
// For information about setting up receipt rules, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rules.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DescribeReceiptRuleRequest.
//    req := client.DescribeReceiptRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DescribeReceiptRule
func (c *Client) DescribeReceiptRuleRequest(input *types.DescribeReceiptRuleInput) DescribeReceiptRuleRequest {
	op := &aws.Operation{
		Name:       opDescribeReceiptRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeReceiptRuleInput{}
	}

	req := c.newRequest(op, input, &types.DescribeReceiptRuleOutput{})
	return DescribeReceiptRuleRequest{Request: req, Input: input, Copy: c.DescribeReceiptRuleRequest}
}

// DescribeReceiptRuleRequest is the request type for the
// DescribeReceiptRule API operation.
type DescribeReceiptRuleRequest struct {
	*aws.Request
	Input *types.DescribeReceiptRuleInput
	Copy  func(*types.DescribeReceiptRuleInput) DescribeReceiptRuleRequest
}

// Send marshals and sends the DescribeReceiptRule API request.
func (r DescribeReceiptRuleRequest) Send(ctx context.Context) (*DescribeReceiptRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReceiptRuleResponse{
		DescribeReceiptRuleOutput: r.Request.Data.(*types.DescribeReceiptRuleOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeReceiptRuleResponse is the response type for the
// DescribeReceiptRule API operation.
type DescribeReceiptRuleResponse struct {
	*types.DescribeReceiptRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReceiptRule request.
func (r *DescribeReceiptRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
