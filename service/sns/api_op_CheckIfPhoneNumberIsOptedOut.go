// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opCheckIfPhoneNumberIsOptedOut = "CheckIfPhoneNumberIsOptedOut"

// CheckIfPhoneNumberIsOptedOutRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Accepts a phone number and indicates whether the phone holder has opted out
// of receiving SMS messages from your account. You cannot send SMS messages
// to a number that is opted out.
//
// To resume sending messages, you can opt in the number by using the OptInPhoneNumber
// action.
//
//    // Example sending a request using CheckIfPhoneNumberIsOptedOutRequest.
//    req := client.CheckIfPhoneNumberIsOptedOutRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/CheckIfPhoneNumberIsOptedOut
func (c *Client) CheckIfPhoneNumberIsOptedOutRequest(input *types.CheckIfPhoneNumberIsOptedOutInput) CheckIfPhoneNumberIsOptedOutRequest {
	op := &aws.Operation{
		Name:       opCheckIfPhoneNumberIsOptedOut,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CheckIfPhoneNumberIsOptedOutInput{}
	}

	req := c.newRequest(op, input, &types.CheckIfPhoneNumberIsOptedOutOutput{})
	return CheckIfPhoneNumberIsOptedOutRequest{Request: req, Input: input, Copy: c.CheckIfPhoneNumberIsOptedOutRequest}
}

// CheckIfPhoneNumberIsOptedOutRequest is the request type for the
// CheckIfPhoneNumberIsOptedOut API operation.
type CheckIfPhoneNumberIsOptedOutRequest struct {
	*aws.Request
	Input *types.CheckIfPhoneNumberIsOptedOutInput
	Copy  func(*types.CheckIfPhoneNumberIsOptedOutInput) CheckIfPhoneNumberIsOptedOutRequest
}

// Send marshals and sends the CheckIfPhoneNumberIsOptedOut API request.
func (r CheckIfPhoneNumberIsOptedOutRequest) Send(ctx context.Context) (*CheckIfPhoneNumberIsOptedOutResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CheckIfPhoneNumberIsOptedOutResponse{
		CheckIfPhoneNumberIsOptedOutOutput: r.Request.Data.(*types.CheckIfPhoneNumberIsOptedOutOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CheckIfPhoneNumberIsOptedOutResponse is the response type for the
// CheckIfPhoneNumberIsOptedOut API operation.
type CheckIfPhoneNumberIsOptedOutResponse struct {
	*types.CheckIfPhoneNumberIsOptedOutOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CheckIfPhoneNumberIsOptedOut request.
func (r *CheckIfPhoneNumberIsOptedOutResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
