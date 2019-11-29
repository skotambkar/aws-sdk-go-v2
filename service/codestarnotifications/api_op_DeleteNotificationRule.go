// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
)

const opDeleteNotificationRule = "DeleteNotificationRule"

// DeleteNotificationRuleRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Deletes a notification rule for a resource.
//
//    // Example sending a request using DeleteNotificationRuleRequest.
//    req := client.DeleteNotificationRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/DeleteNotificationRule
func (c *Client) DeleteNotificationRuleRequest(input *types.DeleteNotificationRuleInput) DeleteNotificationRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteNotificationRule,
		HTTPMethod: "POST",
		HTTPPath:   "/deleteNotificationRule",
	}

	if input == nil {
		input = &types.DeleteNotificationRuleInput{}
	}

	req := c.newRequest(op, input, &types.DeleteNotificationRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteNotificationRuleMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteNotificationRuleRequest{Request: req, Input: input, Copy: c.DeleteNotificationRuleRequest}
}

// DeleteNotificationRuleRequest is the request type for the
// DeleteNotificationRule API operation.
type DeleteNotificationRuleRequest struct {
	*aws.Request
	Input *types.DeleteNotificationRuleInput
	Copy  func(*types.DeleteNotificationRuleInput) DeleteNotificationRuleRequest
}

// Send marshals and sends the DeleteNotificationRule API request.
func (r DeleteNotificationRuleRequest) Send(ctx context.Context) (*DeleteNotificationRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNotificationRuleResponse{
		DeleteNotificationRuleOutput: r.Request.Data.(*types.DeleteNotificationRuleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNotificationRuleResponse is the response type for the
// DeleteNotificationRule API operation.
type DeleteNotificationRuleResponse struct {
	*types.DeleteNotificationRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNotificationRule request.
func (r *DeleteNotificationRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
