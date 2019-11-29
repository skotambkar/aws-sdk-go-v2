// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

const opGetRandomPassword = "GetRandomPassword"

// GetRandomPasswordRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Generates a random password of the specified complexity. This operation is
// intended for use in the Lambda rotation function. Per best practice, we recommend
// that you specify the maximum length and include every character type that
// the system you are generating a password for can support.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:GetRandomPassword
//
//    // Example sending a request using GetRandomPasswordRequest.
//    req := client.GetRandomPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/GetRandomPassword
func (c *Client) GetRandomPasswordRequest(input *types.GetRandomPasswordInput) GetRandomPasswordRequest {
	op := &aws.Operation{
		Name:       opGetRandomPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRandomPasswordInput{}
	}

	req := c.newRequest(op, input, &types.GetRandomPasswordOutput{})
	return GetRandomPasswordRequest{Request: req, Input: input, Copy: c.GetRandomPasswordRequest}
}

// GetRandomPasswordRequest is the request type for the
// GetRandomPassword API operation.
type GetRandomPasswordRequest struct {
	*aws.Request
	Input *types.GetRandomPasswordInput
	Copy  func(*types.GetRandomPasswordInput) GetRandomPasswordRequest
}

// Send marshals and sends the GetRandomPassword API request.
func (r GetRandomPasswordRequest) Send(ctx context.Context) (*GetRandomPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRandomPasswordResponse{
		GetRandomPasswordOutput: r.Request.Data.(*types.GetRandomPasswordOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRandomPasswordResponse is the response type for the
// GetRandomPassword API operation.
type GetRandomPasswordResponse struct {
	*types.GetRandomPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRandomPassword request.
func (r *GetRandomPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
