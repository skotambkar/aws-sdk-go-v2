// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
)

const opCreateNotificationRule = "CreateNotificationRule"

// CreateNotificationRuleRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Creates a notification rule for a resource. The rule specifies the events
// you want notifications about and the targets (such as SNS topics) where you
// want to receive them.
//
//    // Example sending a request using CreateNotificationRuleRequest.
//    req := client.CreateNotificationRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/CreateNotificationRule
func (c *Client) CreateNotificationRuleRequest(input *types.CreateNotificationRuleInput) CreateNotificationRuleRequest {
	op := &aws.Operation{
		Name:       opCreateNotificationRule,
		HTTPMethod: "POST",
		HTTPPath:   "/createNotificationRule",
	}

	if input == nil {
		input = &types.CreateNotificationRuleInput{}
	}

	req := c.newRequest(op, input, &types.CreateNotificationRuleOutput{})
	return CreateNotificationRuleRequest{Request: req, Input: input, Copy: c.CreateNotificationRuleRequest}
}

// CreateNotificationRuleRequest is the request type for the
// CreateNotificationRule API operation.
type CreateNotificationRuleRequest struct {
	*aws.Request
	Input *types.CreateNotificationRuleInput
	Copy  func(*types.CreateNotificationRuleInput) CreateNotificationRuleRequest
}

// Send marshals and sends the CreateNotificationRule API request.
func (r CreateNotificationRuleRequest) Send(ctx context.Context) (*CreateNotificationRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNotificationRuleResponse{
		CreateNotificationRuleOutput: r.Request.Data.(*types.CreateNotificationRuleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNotificationRuleResponse is the response type for the
// CreateNotificationRule API operation.
type CreateNotificationRuleResponse struct {
	*types.CreateNotificationRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNotificationRule request.
func (r *CreateNotificationRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
