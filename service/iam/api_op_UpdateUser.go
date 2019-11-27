// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opUpdateUser = "UpdateUser"

// UpdateUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Updates the name and/or the path of the specified IAM user.
//
// You should understand the implications of changing an IAM user's path or
// name. For more information, see Renaming an IAM User (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_renaming)
// and Renaming an IAM Group (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups_manage_rename.html)
// in the IAM User Guide.
//
// To change a user name, the requester must have appropriate permissions on
// both the source object and the target object. For example, to change Bob
// to Robert, the entity making the request must have permission on Bob and
// Robert, or must have permission on all (*). For more information about permissions,
// see Permissions and Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/PermissionsAndPolicies.html).
//
//    // Example sending a request using UpdateUserRequest.
//    req := client.UpdateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateUser
func (c *Client) UpdateUserRequest(input *types.UpdateUserInput) UpdateUserRequest {
	op := &aws.Operation{
		Name:       opUpdateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateUserInput{}
	}

	req := c.newRequest(op, input, &types.UpdateUserOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateUserRequest{Request: req, Input: input, Copy: c.UpdateUserRequest}
}

// UpdateUserRequest is the request type for the
// UpdateUser API operation.
type UpdateUserRequest struct {
	*aws.Request
	Input *types.UpdateUserInput
	Copy  func(*types.UpdateUserInput) UpdateUserRequest
}

// Send marshals and sends the UpdateUser API request.
func (r UpdateUserRequest) Send(ctx context.Context) (*UpdateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserResponse{
		UpdateUserOutput: r.Request.Data.(*types.UpdateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserResponse is the response type for the
// UpdateUser API operation.
type UpdateUserResponse struct {
	*types.UpdateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUser request.
func (r *UpdateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
