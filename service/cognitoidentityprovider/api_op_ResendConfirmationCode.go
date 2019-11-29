// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opResendConfirmationCode = "ResendConfirmationCode"

// ResendConfirmationCodeRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Resends the confirmation (for confirmation of registration) to a specific
// user in the user pool.
//
//    // Example sending a request using ResendConfirmationCodeRequest.
//    req := client.ResendConfirmationCodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ResendConfirmationCode
func (c *Client) ResendConfirmationCodeRequest(input *types.ResendConfirmationCodeInput) ResendConfirmationCodeRequest {
	op := &aws.Operation{
		Name:       opResendConfirmationCode,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResendConfirmationCodeInput{}
	}

	req := c.newRequest(op, input, &types.ResendConfirmationCodeOutput{})
	req.Config.Credentials = aws.AnonymousCredentials
	return ResendConfirmationCodeRequest{Request: req, Input: input, Copy: c.ResendConfirmationCodeRequest}
}

// ResendConfirmationCodeRequest is the request type for the
// ResendConfirmationCode API operation.
type ResendConfirmationCodeRequest struct {
	*aws.Request
	Input *types.ResendConfirmationCodeInput
	Copy  func(*types.ResendConfirmationCodeInput) ResendConfirmationCodeRequest
}

// Send marshals and sends the ResendConfirmationCode API request.
func (r ResendConfirmationCodeRequest) Send(ctx context.Context) (*ResendConfirmationCodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResendConfirmationCodeResponse{
		ResendConfirmationCodeOutput: r.Request.Data.(*types.ResendConfirmationCodeOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResendConfirmationCodeResponse is the response type for the
// ResendConfirmationCode API operation.
type ResendConfirmationCodeResponse struct {
	*types.ResendConfirmationCodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResendConfirmationCode request.
func (r *ResendConfirmationCodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
