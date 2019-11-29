// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/ses/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opCreateCustomVerificationEmailTemplate = "CreateCustomVerificationEmailTemplate"

// CreateCustomVerificationEmailTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a new custom verification email template.
//
// For more information about custom verification email templates, see Using
// Custom Verification Email Templates (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateCustomVerificationEmailTemplateRequest.
//    req := client.CreateCustomVerificationEmailTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateCustomVerificationEmailTemplate
func (c *Client) CreateCustomVerificationEmailTemplateRequest(input *types.CreateCustomVerificationEmailTemplateInput) CreateCustomVerificationEmailTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateCustomVerificationEmailTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateCustomVerificationEmailTemplateInput{}
	}

	req := c.newRequest(op, input, &types.CreateCustomVerificationEmailTemplateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateCustomVerificationEmailTemplateMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateCustomVerificationEmailTemplateRequest{Request: req, Input: input, Copy: c.CreateCustomVerificationEmailTemplateRequest}
}

// CreateCustomVerificationEmailTemplateRequest is the request type for the
// CreateCustomVerificationEmailTemplate API operation.
type CreateCustomVerificationEmailTemplateRequest struct {
	*aws.Request
	Input *types.CreateCustomVerificationEmailTemplateInput
	Copy  func(*types.CreateCustomVerificationEmailTemplateInput) CreateCustomVerificationEmailTemplateRequest
}

// Send marshals and sends the CreateCustomVerificationEmailTemplate API request.
func (r CreateCustomVerificationEmailTemplateRequest) Send(ctx context.Context) (*CreateCustomVerificationEmailTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCustomVerificationEmailTemplateResponse{
		CreateCustomVerificationEmailTemplateOutput: r.Request.Data.(*types.CreateCustomVerificationEmailTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCustomVerificationEmailTemplateResponse is the response type for the
// CreateCustomVerificationEmailTemplate API operation.
type CreateCustomVerificationEmailTemplateResponse struct {
	*types.CreateCustomVerificationEmailTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCustomVerificationEmailTemplate request.
func (r *CreateCustomVerificationEmailTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
