// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opReorderReceiptRuleSet = "ReorderReceiptRuleSet"

// ReorderReceiptRuleSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Reorders the receipt rules within a receipt rule set.
//
// All of the rules in the rule set must be represented in this request. That
// is, this API will return an error if the reorder request doesn't explicitly
// position all of the rules.
//
// For information about managing receipt rule sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using ReorderReceiptRuleSetRequest.
//    req := client.ReorderReceiptRuleSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ReorderReceiptRuleSet
func (c *Client) ReorderReceiptRuleSetRequest(input *types.ReorderReceiptRuleSetInput) ReorderReceiptRuleSetRequest {
	op := &aws.Operation{
		Name:       opReorderReceiptRuleSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ReorderReceiptRuleSetInput{}
	}

	req := c.newRequest(op, input, &types.ReorderReceiptRuleSetOutput{})
	return ReorderReceiptRuleSetRequest{Request: req, Input: input, Copy: c.ReorderReceiptRuleSetRequest}
}

// ReorderReceiptRuleSetRequest is the request type for the
// ReorderReceiptRuleSet API operation.
type ReorderReceiptRuleSetRequest struct {
	*aws.Request
	Input *types.ReorderReceiptRuleSetInput
	Copy  func(*types.ReorderReceiptRuleSetInput) ReorderReceiptRuleSetRequest
}

// Send marshals and sends the ReorderReceiptRuleSet API request.
func (r ReorderReceiptRuleSetRequest) Send(ctx context.Context) (*ReorderReceiptRuleSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReorderReceiptRuleSetResponse{
		ReorderReceiptRuleSetOutput: r.Request.Data.(*types.ReorderReceiptRuleSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReorderReceiptRuleSetResponse is the response type for the
// ReorderReceiptRuleSet API operation.
type ReorderReceiptRuleSetResponse struct {
	*types.ReorderReceiptRuleSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReorderReceiptRuleSet request.
func (r *ReorderReceiptRuleSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
