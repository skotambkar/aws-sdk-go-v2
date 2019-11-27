// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opDeleteGitHubAccountToken = "DeleteGitHubAccountToken"

// DeleteGitHubAccountTokenRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Deletes a GitHub account connection.
//
//    // Example sending a request using DeleteGitHubAccountTokenRequest.
//    req := client.DeleteGitHubAccountTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeleteGitHubAccountToken
func (c *Client) DeleteGitHubAccountTokenRequest(input *types.DeleteGitHubAccountTokenInput) DeleteGitHubAccountTokenRequest {
	op := &aws.Operation{
		Name:       opDeleteGitHubAccountToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteGitHubAccountTokenInput{}
	}

	req := c.newRequest(op, input, &types.DeleteGitHubAccountTokenOutput{})
	return DeleteGitHubAccountTokenRequest{Request: req, Input: input, Copy: c.DeleteGitHubAccountTokenRequest}
}

// DeleteGitHubAccountTokenRequest is the request type for the
// DeleteGitHubAccountToken API operation.
type DeleteGitHubAccountTokenRequest struct {
	*aws.Request
	Input *types.DeleteGitHubAccountTokenInput
	Copy  func(*types.DeleteGitHubAccountTokenInput) DeleteGitHubAccountTokenRequest
}

// Send marshals and sends the DeleteGitHubAccountToken API request.
func (r DeleteGitHubAccountTokenRequest) Send(ctx context.Context) (*DeleteGitHubAccountTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGitHubAccountTokenResponse{
		DeleteGitHubAccountTokenOutput: r.Request.Data.(*types.DeleteGitHubAccountTokenOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGitHubAccountTokenResponse is the response type for the
// DeleteGitHubAccountToken API operation.
type DeleteGitHubAccountTokenResponse struct {
	*types.DeleteGitHubAccountTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGitHubAccountToken request.
func (r *DeleteGitHubAccountTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
