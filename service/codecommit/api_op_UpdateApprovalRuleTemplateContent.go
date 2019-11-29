// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opUpdateApprovalRuleTemplateContent = "UpdateApprovalRuleTemplateContent"

// UpdateApprovalRuleTemplateContentRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Updates the content of an approval rule template. You can change the number
// of required approvals, the membership of the approval rule, and whether an
// approval pool is defined.
//
//    // Example sending a request using UpdateApprovalRuleTemplateContentRequest.
//    req := client.UpdateApprovalRuleTemplateContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdateApprovalRuleTemplateContent
func (c *Client) UpdateApprovalRuleTemplateContentRequest(input *types.UpdateApprovalRuleTemplateContentInput) UpdateApprovalRuleTemplateContentRequest {
	op := &aws.Operation{
		Name:       opUpdateApprovalRuleTemplateContent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateApprovalRuleTemplateContentInput{}
	}

	req := c.newRequest(op, input, &types.UpdateApprovalRuleTemplateContentOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateApprovalRuleTemplateContentMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateApprovalRuleTemplateContentRequest{Request: req, Input: input, Copy: c.UpdateApprovalRuleTemplateContentRequest}
}

// UpdateApprovalRuleTemplateContentRequest is the request type for the
// UpdateApprovalRuleTemplateContent API operation.
type UpdateApprovalRuleTemplateContentRequest struct {
	*aws.Request
	Input *types.UpdateApprovalRuleTemplateContentInput
	Copy  func(*types.UpdateApprovalRuleTemplateContentInput) UpdateApprovalRuleTemplateContentRequest
}

// Send marshals and sends the UpdateApprovalRuleTemplateContent API request.
func (r UpdateApprovalRuleTemplateContentRequest) Send(ctx context.Context) (*UpdateApprovalRuleTemplateContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApprovalRuleTemplateContentResponse{
		UpdateApprovalRuleTemplateContentOutput: r.Request.Data.(*types.UpdateApprovalRuleTemplateContentOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApprovalRuleTemplateContentResponse is the response type for the
// UpdateApprovalRuleTemplateContent API operation.
type UpdateApprovalRuleTemplateContentResponse struct {
	*types.UpdateApprovalRuleTemplateContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApprovalRuleTemplateContent request.
func (r *UpdateApprovalRuleTemplateContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
