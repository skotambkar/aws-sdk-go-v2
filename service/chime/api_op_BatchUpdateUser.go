// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opBatchUpdateUser = "BatchUpdateUser"

// BatchUpdateUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates user details within the UpdateUserRequestItem object for up to 20
// users for the specified Amazon Chime account. Currently, only LicenseType
// updates are supported for this action.
//
//    // Example sending a request using BatchUpdateUserRequest.
//    req := client.BatchUpdateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchUpdateUser
func (c *Client) BatchUpdateUserRequest(input *types.BatchUpdateUserInput) BatchUpdateUserRequest {
	op := &aws.Operation{
		Name:       opBatchUpdateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users",
	}

	if input == nil {
		input = &types.BatchUpdateUserInput{}
	}

	req := c.newRequest(op, input, &types.BatchUpdateUserOutput{})
	return BatchUpdateUserRequest{Request: req, Input: input, Copy: c.BatchUpdateUserRequest}
}

// BatchUpdateUserRequest is the request type for the
// BatchUpdateUser API operation.
type BatchUpdateUserRequest struct {
	*aws.Request
	Input *types.BatchUpdateUserInput
	Copy  func(*types.BatchUpdateUserInput) BatchUpdateUserRequest
}

// Send marshals and sends the BatchUpdateUser API request.
func (r BatchUpdateUserRequest) Send(ctx context.Context) (*BatchUpdateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchUpdateUserResponse{
		BatchUpdateUserOutput: r.Request.Data.(*types.BatchUpdateUserOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchUpdateUserResponse is the response type for the
// BatchUpdateUser API operation.
type BatchUpdateUserResponse struct {
	*types.BatchUpdateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchUpdateUser request.
func (r *BatchUpdateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
