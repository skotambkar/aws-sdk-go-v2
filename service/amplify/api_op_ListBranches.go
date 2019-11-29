// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

const opListBranches = "ListBranches"

// ListBranchesRequest returns a request value for making API operation for
// AWS Amplify.
//
// Lists branches for an Amplify App.
//
//    // Example sending a request using ListBranchesRequest.
//    req := client.ListBranchesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/ListBranches
func (c *Client) ListBranchesRequest(input *types.ListBranchesInput) ListBranchesRequest {
	op := &aws.Operation{
		Name:       opListBranches,
		HTTPMethod: "GET",
		HTTPPath:   "/apps/{appId}/branches",
	}

	if input == nil {
		input = &types.ListBranchesInput{}
	}

	req := c.newRequest(op, input, &types.ListBranchesOutput{})
	return ListBranchesRequest{Request: req, Input: input, Copy: c.ListBranchesRequest}
}

// ListBranchesRequest is the request type for the
// ListBranches API operation.
type ListBranchesRequest struct {
	*aws.Request
	Input *types.ListBranchesInput
	Copy  func(*types.ListBranchesInput) ListBranchesRequest
}

// Send marshals and sends the ListBranches API request.
func (r ListBranchesRequest) Send(ctx context.Context) (*ListBranchesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBranchesResponse{
		ListBranchesOutput: r.Request.Data.(*types.ListBranchesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBranchesResponse is the response type for the
// ListBranches API operation.
type ListBranchesResponse struct {
	*types.ListBranchesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBranches request.
func (r *ListBranchesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
