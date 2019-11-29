// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opListGroups = "ListGroups"

// ListGroupsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists all user groups in Amazon QuickSight.
//
// The permissions resource is arn:aws:quicksight:us-east-1:<aws-account-id>:group/default/*.
//
// The response is a list of group objects.
//
// CLI Sample:
//
// aws quicksight list-groups -\-aws-account-id=111122223333 -\-namespace=default
//
//    // Example sending a request using ListGroupsRequest.
//    req := client.ListGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListGroups
func (c *Client) ListGroupsRequest(input *types.ListGroupsInput) ListGroupsRequest {
	op := &aws.Operation{
		Name:       opListGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/groups",
	}

	if input == nil {
		input = &types.ListGroupsInput{}
	}

	req := c.newRequest(op, input, &types.ListGroupsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListGroupsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListGroupsRequest{Request: req, Input: input, Copy: c.ListGroupsRequest}
}

// ListGroupsRequest is the request type for the
// ListGroups API operation.
type ListGroupsRequest struct {
	*aws.Request
	Input *types.ListGroupsInput
	Copy  func(*types.ListGroupsInput) ListGroupsRequest
}

// Send marshals and sends the ListGroups API request.
func (r ListGroupsRequest) Send(ctx context.Context) (*ListGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupsResponse{
		ListGroupsOutput: r.Request.Data.(*types.ListGroupsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGroupsResponse is the response type for the
// ListGroups API operation.
type ListGroupsResponse struct {
	*types.ListGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroups request.
func (r *ListGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
