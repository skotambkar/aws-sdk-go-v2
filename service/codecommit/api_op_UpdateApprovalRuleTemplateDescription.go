// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opUpdateApprovalRuleTemplateDescription = "UpdateApprovalRuleTemplateDescription"

// UpdateApprovalRuleTemplateDescriptionRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Updates the description for a specified approval rule template.
//
//    // Example sending a request using UpdateApprovalRuleTemplateDescriptionRequest.
//    req := client.UpdateApprovalRuleTemplateDescriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdateApprovalRuleTemplateDescription
func (c *Client) UpdateApprovalRuleTemplateDescriptionRequest(input *types.UpdateApprovalRuleTemplateDescriptionInput) UpdateApprovalRuleTemplateDescriptionRequest {
	op := &aws.Operation{
		Name:       opUpdateApprovalRuleTemplateDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateApprovalRuleTemplateDescriptionInput{}
	}

	req := c.newRequest(op, input, &types.UpdateApprovalRuleTemplateDescriptionOutput{})
	return UpdateApprovalRuleTemplateDescriptionRequest{Request: req, Input: input, Copy: c.UpdateApprovalRuleTemplateDescriptionRequest}
}

// UpdateApprovalRuleTemplateDescriptionRequest is the request type for the
// UpdateApprovalRuleTemplateDescription API operation.
type UpdateApprovalRuleTemplateDescriptionRequest struct {
	*aws.Request
	Input *types.UpdateApprovalRuleTemplateDescriptionInput
	Copy  func(*types.UpdateApprovalRuleTemplateDescriptionInput) UpdateApprovalRuleTemplateDescriptionRequest
}

// Send marshals and sends the UpdateApprovalRuleTemplateDescription API request.
func (r UpdateApprovalRuleTemplateDescriptionRequest) Send(ctx context.Context) (*UpdateApprovalRuleTemplateDescriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApprovalRuleTemplateDescriptionResponse{
		UpdateApprovalRuleTemplateDescriptionOutput: r.Request.Data.(*types.UpdateApprovalRuleTemplateDescriptionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApprovalRuleTemplateDescriptionResponse is the response type for the
// UpdateApprovalRuleTemplateDescription API operation.
type UpdateApprovalRuleTemplateDescriptionResponse struct {
	*types.UpdateApprovalRuleTemplateDescriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApprovalRuleTemplateDescription request.
func (r *UpdateApprovalRuleTemplateDescriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
