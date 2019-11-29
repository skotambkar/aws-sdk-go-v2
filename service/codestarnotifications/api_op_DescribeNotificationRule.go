// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
)

const opDescribeNotificationRule = "DescribeNotificationRule"

// DescribeNotificationRuleRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Returns information about a specified notification rule.
//
//    // Example sending a request using DescribeNotificationRuleRequest.
//    req := client.DescribeNotificationRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/DescribeNotificationRule
func (c *Client) DescribeNotificationRuleRequest(input *types.DescribeNotificationRuleInput) DescribeNotificationRuleRequest {
	op := &aws.Operation{
		Name:       opDescribeNotificationRule,
		HTTPMethod: "POST",
		HTTPPath:   "/describeNotificationRule",
	}

	if input == nil {
		input = &types.DescribeNotificationRuleInput{}
	}

	req := c.newRequest(op, input, &types.DescribeNotificationRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeNotificationRuleMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeNotificationRuleRequest{Request: req, Input: input, Copy: c.DescribeNotificationRuleRequest}
}

// DescribeNotificationRuleRequest is the request type for the
// DescribeNotificationRule API operation.
type DescribeNotificationRuleRequest struct {
	*aws.Request
	Input *types.DescribeNotificationRuleInput
	Copy  func(*types.DescribeNotificationRuleInput) DescribeNotificationRuleRequest
}

// Send marshals and sends the DescribeNotificationRule API request.
func (r DescribeNotificationRuleRequest) Send(ctx context.Context) (*DescribeNotificationRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNotificationRuleResponse{
		DescribeNotificationRuleOutput: r.Request.Data.(*types.DescribeNotificationRuleOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeNotificationRuleResponse is the response type for the
// DescribeNotificationRule API operation.
type DescribeNotificationRuleResponse struct {
	*types.DescribeNotificationRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNotificationRule request.
func (r *DescribeNotificationRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
