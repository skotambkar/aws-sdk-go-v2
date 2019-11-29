// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opResetPersonalPIN = "ResetPersonalPIN"

// ResetPersonalPINRequest returns a request value for making API operation for
// Amazon Chime.
//
// Resets the personal meeting PIN for the specified user on an Amazon Chime
// account. Returns the User object with the updated personal meeting PIN.
//
//    // Example sending a request using ResetPersonalPINRequest.
//    req := client.ResetPersonalPINRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ResetPersonalPIN
func (c *Client) ResetPersonalPINRequest(input *types.ResetPersonalPINInput) ResetPersonalPINRequest {
	op := &aws.Operation{
		Name:       opResetPersonalPIN,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users/{userId}?operation=reset-personal-pin",
	}

	if input == nil {
		input = &types.ResetPersonalPINInput{}
	}

	req := c.newRequest(op, input, &types.ResetPersonalPINOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ResetPersonalPINMarshaler{Input: input}.GetNamedBuildHandler())

	return ResetPersonalPINRequest{Request: req, Input: input, Copy: c.ResetPersonalPINRequest}
}

// ResetPersonalPINRequest is the request type for the
// ResetPersonalPIN API operation.
type ResetPersonalPINRequest struct {
	*aws.Request
	Input *types.ResetPersonalPINInput
	Copy  func(*types.ResetPersonalPINInput) ResetPersonalPINRequest
}

// Send marshals and sends the ResetPersonalPIN API request.
func (r ResetPersonalPINRequest) Send(ctx context.Context) (*ResetPersonalPINResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetPersonalPINResponse{
		ResetPersonalPINOutput: r.Request.Data.(*types.ResetPersonalPINOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetPersonalPINResponse is the response type for the
// ResetPersonalPIN API operation.
type ResetPersonalPINResponse struct {
	*types.ResetPersonalPINOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetPersonalPIN request.
func (r *ResetPersonalPINResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
