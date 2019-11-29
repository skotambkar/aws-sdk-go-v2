// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/ses/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opGetCustomVerificationEmailTemplate = "GetCustomVerificationEmailTemplate"

// GetCustomVerificationEmailTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns the custom email verification template for the template name you
// specify.
//
// For more information about custom verification email templates, see Using
// Custom Verification Email Templates (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using GetCustomVerificationEmailTemplateRequest.
//    req := client.GetCustomVerificationEmailTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetCustomVerificationEmailTemplate
func (c *Client) GetCustomVerificationEmailTemplateRequest(input *types.GetCustomVerificationEmailTemplateInput) GetCustomVerificationEmailTemplateRequest {
	op := &aws.Operation{
		Name:       opGetCustomVerificationEmailTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetCustomVerificationEmailTemplateInput{}
	}

	req := c.newRequest(op, input, &types.GetCustomVerificationEmailTemplateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.GetCustomVerificationEmailTemplateMarshaler{Input: input}.GetNamedBuildHandler())

	return GetCustomVerificationEmailTemplateRequest{Request: req, Input: input, Copy: c.GetCustomVerificationEmailTemplateRequest}
}

// GetCustomVerificationEmailTemplateRequest is the request type for the
// GetCustomVerificationEmailTemplate API operation.
type GetCustomVerificationEmailTemplateRequest struct {
	*aws.Request
	Input *types.GetCustomVerificationEmailTemplateInput
	Copy  func(*types.GetCustomVerificationEmailTemplateInput) GetCustomVerificationEmailTemplateRequest
}

// Send marshals and sends the GetCustomVerificationEmailTemplate API request.
func (r GetCustomVerificationEmailTemplateRequest) Send(ctx context.Context) (*GetCustomVerificationEmailTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCustomVerificationEmailTemplateResponse{
		GetCustomVerificationEmailTemplateOutput: r.Request.Data.(*types.GetCustomVerificationEmailTemplateOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCustomVerificationEmailTemplateResponse is the response type for the
// GetCustomVerificationEmailTemplate API operation.
type GetCustomVerificationEmailTemplateResponse struct {
	*types.GetCustomVerificationEmailTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCustomVerificationEmailTemplate request.
func (r *GetCustomVerificationEmailTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
