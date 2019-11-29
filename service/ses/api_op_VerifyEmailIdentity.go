// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/ses/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opVerifyEmailIdentity = "VerifyEmailIdentity"

// VerifyEmailIdentityRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Adds an email address to the list of identities for your Amazon SES account
// in the current AWS region and attempts to verify it. As a result of executing
// this operation, a verification email is sent to the specified address.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using VerifyEmailIdentityRequest.
//    req := client.VerifyEmailIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/VerifyEmailIdentity
func (c *Client) VerifyEmailIdentityRequest(input *types.VerifyEmailIdentityInput) VerifyEmailIdentityRequest {
	op := &aws.Operation{
		Name:       opVerifyEmailIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.VerifyEmailIdentityInput{}
	}

	req := c.newRequest(op, input, &types.VerifyEmailIdentityOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.VerifyEmailIdentityMarshaler{Input: input}.GetNamedBuildHandler())

	return VerifyEmailIdentityRequest{Request: req, Input: input, Copy: c.VerifyEmailIdentityRequest}
}

// VerifyEmailIdentityRequest is the request type for the
// VerifyEmailIdentity API operation.
type VerifyEmailIdentityRequest struct {
	*aws.Request
	Input *types.VerifyEmailIdentityInput
	Copy  func(*types.VerifyEmailIdentityInput) VerifyEmailIdentityRequest
}

// Send marshals and sends the VerifyEmailIdentity API request.
func (r VerifyEmailIdentityRequest) Send(ctx context.Context) (*VerifyEmailIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &VerifyEmailIdentityResponse{
		VerifyEmailIdentityOutput: r.Request.Data.(*types.VerifyEmailIdentityOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// VerifyEmailIdentityResponse is the response type for the
// VerifyEmailIdentity API operation.
type VerifyEmailIdentityResponse struct {
	*types.VerifyEmailIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// VerifyEmailIdentity request.
func (r *VerifyEmailIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
