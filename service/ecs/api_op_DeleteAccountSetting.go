// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opDeleteAccountSetting = "DeleteAccountSetting"

// DeleteAccountSettingRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Disables an account setting for a specified IAM user, IAM role, or the root
// user for an account.
//
//    // Example sending a request using DeleteAccountSettingRequest.
//    req := client.DeleteAccountSettingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeleteAccountSetting
func (c *Client) DeleteAccountSettingRequest(input *types.DeleteAccountSettingInput) DeleteAccountSettingRequest {
	op := &aws.Operation{
		Name:       opDeleteAccountSetting,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAccountSettingInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAccountSettingOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteAccountSettingMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteAccountSettingRequest{Request: req, Input: input, Copy: c.DeleteAccountSettingRequest}
}

// DeleteAccountSettingRequest is the request type for the
// DeleteAccountSetting API operation.
type DeleteAccountSettingRequest struct {
	*aws.Request
	Input *types.DeleteAccountSettingInput
	Copy  func(*types.DeleteAccountSettingInput) DeleteAccountSettingRequest
}

// Send marshals and sends the DeleteAccountSetting API request.
func (r DeleteAccountSettingRequest) Send(ctx context.Context) (*DeleteAccountSettingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccountSettingResponse{
		DeleteAccountSettingOutput: r.Request.Data.(*types.DeleteAccountSettingOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccountSettingResponse is the response type for the
// DeleteAccountSetting API operation.
type DeleteAccountSettingResponse struct {
	*types.DeleteAccountSettingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccountSetting request.
func (r *DeleteAccountSettingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
