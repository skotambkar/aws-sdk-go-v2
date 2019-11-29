// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opListUsers = "ListUsers"

// ListUsersRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Returns a list of all of the Amazon QuickSight users belonging to this account.
//
// The response is a list of user objects, containing each user's Amazon Resource
// Name (ARN), AWS Identity and Access Management (IAM) role, and email address.
//
// CLI Sample:
//
// aws quicksight list-users --aws-account-id=111122223333 --namespace=default
//
//    // Example sending a request using ListUsersRequest.
//    req := client.ListUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUsers
func (c *Client) ListUsersRequest(input *types.ListUsersInput) ListUsersRequest {
	op := &aws.Operation{
		Name:       opListUsers,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users",
	}

	if input == nil {
		input = &types.ListUsersInput{}
	}

	req := c.newRequest(op, input, &types.ListUsersOutput{})
	return ListUsersRequest{Request: req, Input: input, Copy: c.ListUsersRequest}
}

// ListUsersRequest is the request type for the
// ListUsers API operation.
type ListUsersRequest struct {
	*aws.Request
	Input *types.ListUsersInput
	Copy  func(*types.ListUsersInput) ListUsersRequest
}

// Send marshals and sends the ListUsers API request.
func (r ListUsersRequest) Send(ctx context.Context) (*ListUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersResponse{
		ListUsersOutput: r.Request.Data.(*types.ListUsersOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUsersResponse is the response type for the
// ListUsers API operation.
type ListUsersResponse struct {
	*types.ListUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsers request.
func (r *ListUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
