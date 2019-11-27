// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opDescribeActiveReceiptRuleSet = "DescribeActiveReceiptRuleSet"

// DescribeActiveReceiptRuleSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns the metadata and receipt rules for the receipt rule set that is currently
// active.
//
// For information about setting up receipt rule sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DescribeActiveReceiptRuleSetRequest.
//    req := client.DescribeActiveReceiptRuleSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DescribeActiveReceiptRuleSet
func (c *Client) DescribeActiveReceiptRuleSetRequest(input *types.DescribeActiveReceiptRuleSetInput) DescribeActiveReceiptRuleSetRequest {
	op := &aws.Operation{
		Name:       opDescribeActiveReceiptRuleSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeActiveReceiptRuleSetInput{}
	}

	req := c.newRequest(op, input, &types.DescribeActiveReceiptRuleSetOutput{})
	return DescribeActiveReceiptRuleSetRequest{Request: req, Input: input, Copy: c.DescribeActiveReceiptRuleSetRequest}
}

// DescribeActiveReceiptRuleSetRequest is the request type for the
// DescribeActiveReceiptRuleSet API operation.
type DescribeActiveReceiptRuleSetRequest struct {
	*aws.Request
	Input *types.DescribeActiveReceiptRuleSetInput
	Copy  func(*types.DescribeActiveReceiptRuleSetInput) DescribeActiveReceiptRuleSetRequest
}

// Send marshals and sends the DescribeActiveReceiptRuleSet API request.
func (r DescribeActiveReceiptRuleSetRequest) Send(ctx context.Context) (*DescribeActiveReceiptRuleSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeActiveReceiptRuleSetResponse{
		DescribeActiveReceiptRuleSetOutput: r.Request.Data.(*types.DescribeActiveReceiptRuleSetOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeActiveReceiptRuleSetResponse is the response type for the
// DescribeActiveReceiptRuleSet API operation.
type DescribeActiveReceiptRuleSetResponse struct {
	*types.DescribeActiveReceiptRuleSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeActiveReceiptRuleSet request.
func (r *DescribeActiveReceiptRuleSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
