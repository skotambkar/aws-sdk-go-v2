// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opUpdateCustomVerificationEmailTemplate = "UpdateCustomVerificationEmailTemplate"

// UpdateCustomVerificationEmailTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Updates an existing custom verification email template.
//
// For more information about custom verification email templates, see Using
// Custom Verification Email Templates (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using UpdateCustomVerificationEmailTemplateRequest.
//    req := client.UpdateCustomVerificationEmailTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/UpdateCustomVerificationEmailTemplate
func (c *Client) UpdateCustomVerificationEmailTemplateRequest(input *types.UpdateCustomVerificationEmailTemplateInput) UpdateCustomVerificationEmailTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateCustomVerificationEmailTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateCustomVerificationEmailTemplateInput{}
	}

	req := c.newRequest(op, input, &types.UpdateCustomVerificationEmailTemplateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateCustomVerificationEmailTemplateRequest{Request: req, Input: input, Copy: c.UpdateCustomVerificationEmailTemplateRequest}
}

// UpdateCustomVerificationEmailTemplateRequest is the request type for the
// UpdateCustomVerificationEmailTemplate API operation.
type UpdateCustomVerificationEmailTemplateRequest struct {
	*aws.Request
	Input *types.UpdateCustomVerificationEmailTemplateInput
	Copy  func(*types.UpdateCustomVerificationEmailTemplateInput) UpdateCustomVerificationEmailTemplateRequest
}

// Send marshals and sends the UpdateCustomVerificationEmailTemplate API request.
func (r UpdateCustomVerificationEmailTemplateRequest) Send(ctx context.Context) (*UpdateCustomVerificationEmailTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCustomVerificationEmailTemplateResponse{
		UpdateCustomVerificationEmailTemplateOutput: r.Request.Data.(*types.UpdateCustomVerificationEmailTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCustomVerificationEmailTemplateResponse is the response type for the
// UpdateCustomVerificationEmailTemplate API operation.
type UpdateCustomVerificationEmailTemplateResponse struct {
	*types.UpdateCustomVerificationEmailTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCustomVerificationEmailTemplate request.
func (r *UpdateCustomVerificationEmailTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
