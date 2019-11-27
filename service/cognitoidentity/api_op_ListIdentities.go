// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
)

const opListIdentities = "ListIdentities"

// ListIdentitiesRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Lists the identities in an identity pool.
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using ListIdentitiesRequest.
//    req := client.ListIdentitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/ListIdentities
func (c *Client) ListIdentitiesRequest(input *types.ListIdentitiesInput) ListIdentitiesRequest {
	op := &aws.Operation{
		Name:       opListIdentities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListIdentitiesInput{}
	}

	req := c.newRequest(op, input, &types.ListIdentitiesOutput{})
	return ListIdentitiesRequest{Request: req, Input: input, Copy: c.ListIdentitiesRequest}
}

// ListIdentitiesRequest is the request type for the
// ListIdentities API operation.
type ListIdentitiesRequest struct {
	*aws.Request
	Input *types.ListIdentitiesInput
	Copy  func(*types.ListIdentitiesInput) ListIdentitiesRequest
}

// Send marshals and sends the ListIdentities API request.
func (r ListIdentitiesRequest) Send(ctx context.Context) (*ListIdentitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIdentitiesResponse{
		ListIdentitiesOutput: r.Request.Data.(*types.ListIdentitiesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIdentitiesResponse is the response type for the
// ListIdentities API operation.
type ListIdentitiesResponse struct {
	*types.ListIdentitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIdentities request.
func (r *ListIdentitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
