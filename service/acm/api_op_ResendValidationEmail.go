// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/acm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
)

const opResendValidationEmail = "ResendValidationEmail"

// ResendValidationEmailRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Resends the email that requests domain ownership validation. The domain owner
// or an authorized representative must approve the ACM certificate before it
// can be issued. The certificate can be approved by clicking a link in the
// mail to navigate to the Amazon certificate approval website and then clicking
// I Approve. However, the validation email can be blocked by spam filters.
// Therefore, if you do not receive the original mail, you can request that
// the mail be resent within 72 hours of requesting the ACM certificate. If
// more than 72 hours have elapsed since your original request or since your
// last attempt to resend validation mail, you must request a new certificate.
// For more information about setting up your contact email addresses, see Configure
// Email for your Domain (https://docs.aws.amazon.com/acm/latest/userguide/setup-email.html).
//
//    // Example sending a request using ResendValidationEmailRequest.
//    req := client.ResendValidationEmailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ResendValidationEmail
func (c *Client) ResendValidationEmailRequest(input *types.ResendValidationEmailInput) ResendValidationEmailRequest {
	op := &aws.Operation{
		Name:       opResendValidationEmail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResendValidationEmailInput{}
	}

	req := c.newRequest(op, input, &types.ResendValidationEmailOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ResendValidationEmailMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ResendValidationEmailRequest{Request: req, Input: input, Copy: c.ResendValidationEmailRequest}
}

// ResendValidationEmailRequest is the request type for the
// ResendValidationEmail API operation.
type ResendValidationEmailRequest struct {
	*aws.Request
	Input *types.ResendValidationEmailInput
	Copy  func(*types.ResendValidationEmailInput) ResendValidationEmailRequest
}

// Send marshals and sends the ResendValidationEmail API request.
func (r ResendValidationEmailRequest) Send(ctx context.Context) (*ResendValidationEmailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResendValidationEmailResponse{
		ResendValidationEmailOutput: r.Request.Data.(*types.ResendValidationEmailOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResendValidationEmailResponse is the response type for the
// ResendValidationEmail API operation.
type ResendValidationEmailResponse struct {
	*types.ResendValidationEmailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResendValidationEmail request.
func (r *ResendValidationEmailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
