// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/transfer/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
)

const opDeleteUser = "DeleteUser"

// DeleteUserRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Deletes the user belonging to the server you specify.
//
// No response returns from this operation.
//
// When you delete a user from a server, the user's information is lost.
//
//    // Example sending a request using DeleteUserRequest.
//    req := client.DeleteUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DeleteUser
func (c *Client) DeleteUserRequest(input *types.DeleteUserInput) DeleteUserRequest {
	op := &aws.Operation{
		Name:       opDeleteUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteUserInput{}
	}

	req := c.newRequest(op, input, &types.DeleteUserOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteUserMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteUserRequest{Request: req, Input: input, Copy: c.DeleteUserRequest}
}

// DeleteUserRequest is the request type for the
// DeleteUser API operation.
type DeleteUserRequest struct {
	*aws.Request
	Input *types.DeleteUserInput
	Copy  func(*types.DeleteUserInput) DeleteUserRequest
}

// Send marshals and sends the DeleteUser API request.
func (r DeleteUserRequest) Send(ctx context.Context) (*DeleteUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserResponse{
		DeleteUserOutput: r.Request.Data.(*types.DeleteUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserResponse is the response type for the
// DeleteUser API operation.
type DeleteUserResponse struct {
	*types.DeleteUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUser request.
func (r *DeleteUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
