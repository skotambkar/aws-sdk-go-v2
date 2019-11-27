// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDeleteUser = "DeleteUser"

// DeleteUserRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Deletes the Amazon QuickSight user that is associated with the identity of
// the AWS Identity and Access Management (IAM) user or role that's making the
// call. The IAM user isn't deleted as a result of this call.
//
// The permission resource is arn:aws:quicksight:us-east-1:<aws-account-id>:user/default/<user-name> .
//
// CLI Sample:
//
// aws quicksight delete-user --aws-account-id=111122223333 --namespace=default
// --user-name=Pat
//
//    // Example sending a request using DeleteUserRequest.
//    req := client.DeleteUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteUser
func (c *Client) DeleteUserRequest(input *types.DeleteUserInput) DeleteUserRequest {
	op := &aws.Operation{
		Name:       opDeleteUser,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users/{UserName}",
	}

	if input == nil {
		input = &types.DeleteUserInput{}
	}

	req := c.newRequest(op, input, &types.DeleteUserOutput{})
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
