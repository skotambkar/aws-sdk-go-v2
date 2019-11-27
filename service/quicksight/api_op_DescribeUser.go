// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDescribeUser = "DescribeUser"

// DescribeUserRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Returns information about a user, given the user name.
//
// The permission resource is arn:aws:quicksight:us-east-1:<aws-account-id>:user/default/<user-name> .
//
// The response is a user object that contains the user's Amazon Resource Name
// (ARN), AWS Identity and Access Management (IAM) role, and email address.
//
// CLI Sample:
//
// aws quicksight describe-user --aws-account-id=111122223333 --namespace=default
// --user-name=Pat
//
//    // Example sending a request using DescribeUserRequest.
//    req := client.DescribeUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeUser
func (c *Client) DescribeUserRequest(input *types.DescribeUserInput) DescribeUserRequest {
	op := &aws.Operation{
		Name:       opDescribeUser,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users/{UserName}",
	}

	if input == nil {
		input = &types.DescribeUserInput{}
	}

	req := c.newRequest(op, input, &types.DescribeUserOutput{})
	return DescribeUserRequest{Request: req, Input: input, Copy: c.DescribeUserRequest}
}

// DescribeUserRequest is the request type for the
// DescribeUser API operation.
type DescribeUserRequest struct {
	*aws.Request
	Input *types.DescribeUserInput
	Copy  func(*types.DescribeUserInput) DescribeUserRequest
}

// Send marshals and sends the DescribeUser API request.
func (r DescribeUserRequest) Send(ctx context.Context) (*DescribeUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserResponse{
		DescribeUserOutput: r.Request.Data.(*types.DescribeUserOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserResponse is the response type for the
// DescribeUser API operation.
type DescribeUserResponse struct {
	*types.DescribeUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUser request.
func (r *DescribeUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
