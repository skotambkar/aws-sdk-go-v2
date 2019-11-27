// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opListGitHubAccountTokenNames = "ListGitHubAccountTokenNames"

// ListGitHubAccountTokenNamesRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Lists the names of stored connections to GitHub accounts.
//
//    // Example sending a request using ListGitHubAccountTokenNamesRequest.
//    req := client.ListGitHubAccountTokenNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListGitHubAccountTokenNames
func (c *Client) ListGitHubAccountTokenNamesRequest(input *types.ListGitHubAccountTokenNamesInput) ListGitHubAccountTokenNamesRequest {
	op := &aws.Operation{
		Name:       opListGitHubAccountTokenNames,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListGitHubAccountTokenNamesInput{}
	}

	req := c.newRequest(op, input, &types.ListGitHubAccountTokenNamesOutput{})
	return ListGitHubAccountTokenNamesRequest{Request: req, Input: input, Copy: c.ListGitHubAccountTokenNamesRequest}
}

// ListGitHubAccountTokenNamesRequest is the request type for the
// ListGitHubAccountTokenNames API operation.
type ListGitHubAccountTokenNamesRequest struct {
	*aws.Request
	Input *types.ListGitHubAccountTokenNamesInput
	Copy  func(*types.ListGitHubAccountTokenNamesInput) ListGitHubAccountTokenNamesRequest
}

// Send marshals and sends the ListGitHubAccountTokenNames API request.
func (r ListGitHubAccountTokenNamesRequest) Send(ctx context.Context) (*ListGitHubAccountTokenNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGitHubAccountTokenNamesResponse{
		ListGitHubAccountTokenNamesOutput: r.Request.Data.(*types.ListGitHubAccountTokenNamesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGitHubAccountTokenNamesResponse is the response type for the
// ListGitHubAccountTokenNames API operation.
type ListGitHubAccountTokenNamesResponse struct {
	*types.ListGitHubAccountTokenNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGitHubAccountTokenNames request.
func (r *ListGitHubAccountTokenNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
