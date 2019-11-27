// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opUpdateChapCredentials = "UpdateChapCredentials"

// UpdateChapCredentialsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates the Challenge-Handshake Authentication Protocol (CHAP) credentials
// for a specified iSCSI target. By default, a gateway does not have CHAP enabled;
// however, for added security, you might use it.
//
// When you update CHAP credentials, all existing connections on the target
// are closed and initiators must reconnect with the new credentials.
//
//    // Example sending a request using UpdateChapCredentialsRequest.
//    req := client.UpdateChapCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateChapCredentials
func (c *Client) UpdateChapCredentialsRequest(input *types.UpdateChapCredentialsInput) UpdateChapCredentialsRequest {
	op := &aws.Operation{
		Name:       opUpdateChapCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateChapCredentialsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateChapCredentialsOutput{})
	return UpdateChapCredentialsRequest{Request: req, Input: input, Copy: c.UpdateChapCredentialsRequest}
}

// UpdateChapCredentialsRequest is the request type for the
// UpdateChapCredentials API operation.
type UpdateChapCredentialsRequest struct {
	*aws.Request
	Input *types.UpdateChapCredentialsInput
	Copy  func(*types.UpdateChapCredentialsInput) UpdateChapCredentialsRequest
}

// Send marshals and sends the UpdateChapCredentials API request.
func (r UpdateChapCredentialsRequest) Send(ctx context.Context) (*UpdateChapCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateChapCredentialsResponse{
		UpdateChapCredentialsOutput: r.Request.Data.(*types.UpdateChapCredentialsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateChapCredentialsResponse is the response type for the
// UpdateChapCredentials API operation.
type UpdateChapCredentialsResponse struct {
	*types.UpdateChapCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateChapCredentials request.
func (r *UpdateChapCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
