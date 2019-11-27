// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
)

const opListIdentityPoolUsage = "ListIdentityPoolUsage"

// ListIdentityPoolUsageRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Gets a list of identity pools registered with Cognito.
//
// ListIdentityPoolUsage can only be called with developer credentials. You
// cannot make this API call with the temporary user credentials provided by
// Cognito Identity.
//
//    // Example sending a request using ListIdentityPoolUsageRequest.
//    req := client.ListIdentityPoolUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/ListIdentityPoolUsage
func (c *Client) ListIdentityPoolUsageRequest(input *types.ListIdentityPoolUsageInput) ListIdentityPoolUsageRequest {
	op := &aws.Operation{
		Name:       opListIdentityPoolUsage,
		HTTPMethod: "GET",
		HTTPPath:   "/identitypools",
	}

	if input == nil {
		input = &types.ListIdentityPoolUsageInput{}
	}

	req := c.newRequest(op, input, &types.ListIdentityPoolUsageOutput{})
	return ListIdentityPoolUsageRequest{Request: req, Input: input, Copy: c.ListIdentityPoolUsageRequest}
}

// ListIdentityPoolUsageRequest is the request type for the
// ListIdentityPoolUsage API operation.
type ListIdentityPoolUsageRequest struct {
	*aws.Request
	Input *types.ListIdentityPoolUsageInput
	Copy  func(*types.ListIdentityPoolUsageInput) ListIdentityPoolUsageRequest
}

// Send marshals and sends the ListIdentityPoolUsage API request.
func (r ListIdentityPoolUsageRequest) Send(ctx context.Context) (*ListIdentityPoolUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIdentityPoolUsageResponse{
		ListIdentityPoolUsageOutput: r.Request.Data.(*types.ListIdentityPoolUsageOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIdentityPoolUsageResponse is the response type for the
// ListIdentityPoolUsage API operation.
type ListIdentityPoolUsageResponse struct {
	*types.ListIdentityPoolUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIdentityPoolUsage request.
func (r *ListIdentityPoolUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
