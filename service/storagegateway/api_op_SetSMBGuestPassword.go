// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opSetSMBGuestPassword = "SetSMBGuestPassword"

// SetSMBGuestPasswordRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Sets the password for the guest user smbguest. The smbguest user is the user
// when the authentication method for the file share is set to GuestAccess.
//
//    // Example sending a request using SetSMBGuestPasswordRequest.
//    req := client.SetSMBGuestPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/SetSMBGuestPassword
func (c *Client) SetSMBGuestPasswordRequest(input *types.SetSMBGuestPasswordInput) SetSMBGuestPasswordRequest {
	op := &aws.Operation{
		Name:       opSetSMBGuestPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetSMBGuestPasswordInput{}
	}

	req := c.newRequest(op, input, &types.SetSMBGuestPasswordOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.SetSMBGuestPasswordMarshaler{Input: input}.GetNamedBuildHandler())

	return SetSMBGuestPasswordRequest{Request: req, Input: input, Copy: c.SetSMBGuestPasswordRequest}
}

// SetSMBGuestPasswordRequest is the request type for the
// SetSMBGuestPassword API operation.
type SetSMBGuestPasswordRequest struct {
	*aws.Request
	Input *types.SetSMBGuestPasswordInput
	Copy  func(*types.SetSMBGuestPasswordInput) SetSMBGuestPasswordRequest
}

// Send marshals and sends the SetSMBGuestPassword API request.
func (r SetSMBGuestPasswordRequest) Send(ctx context.Context) (*SetSMBGuestPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetSMBGuestPasswordResponse{
		SetSMBGuestPasswordOutput: r.Request.Data.(*types.SetSMBGuestPasswordOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetSMBGuestPasswordResponse is the response type for the
// SetSMBGuestPassword API operation.
type SetSMBGuestPasswordResponse struct {
	*types.SetSMBGuestPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetSMBGuestPassword request.
func (r *SetSMBGuestPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
