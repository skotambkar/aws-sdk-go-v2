// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opListUserGroups = "ListUserGroups"

// ListUserGroupsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists the Amazon QuickSight groups that an Amazon QuickSight user is a member
// of.
//
// The response is a one or more group objects.
//
// CLI Sample:
//
// aws quicksight list-user-groups -\-user-name=Pat -\-aws-account-id=111122223333
// -\-namespace=default -\-region=us-east-1
//
//    // Example sending a request using ListUserGroupsRequest.
//    req := client.ListUserGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUserGroups
func (c *Client) ListUserGroupsRequest(input *types.ListUserGroupsInput) ListUserGroupsRequest {
	op := &aws.Operation{
		Name:       opListUserGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users/{UserName}/groups",
	}

	if input == nil {
		input = &types.ListUserGroupsInput{}
	}

	req := c.newRequest(op, input, &types.ListUserGroupsOutput{})
	return ListUserGroupsRequest{Request: req, Input: input, Copy: c.ListUserGroupsRequest}
}

// ListUserGroupsRequest is the request type for the
// ListUserGroups API operation.
type ListUserGroupsRequest struct {
	*aws.Request
	Input *types.ListUserGroupsInput
	Copy  func(*types.ListUserGroupsInput) ListUserGroupsRequest
}

// Send marshals and sends the ListUserGroups API request.
func (r ListUserGroupsRequest) Send(ctx context.Context) (*ListUserGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUserGroupsResponse{
		ListUserGroupsOutput: r.Request.Data.(*types.ListUserGroupsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUserGroupsResponse is the response type for the
// ListUserGroups API operation.
type ListUserGroupsResponse struct {
	*types.ListUserGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUserGroups request.
func (r *ListUserGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
