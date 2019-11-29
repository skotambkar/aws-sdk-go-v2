// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
)

const opListDatasets = "ListDatasets"

// ListDatasetsRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Lists datasets for an identity. With Amazon Cognito Sync, each identity has
// access only to its own data. Thus, the credentials used to make this API
// call need to have access to the identity data.
//
// ListDatasets can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials. You should use the Cognito Identity
// credentials to make this API call.
//
//    // Example sending a request using ListDatasetsRequest.
//    req := client.ListDatasetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/ListDatasets
func (c *Client) ListDatasetsRequest(input *types.ListDatasetsInput) ListDatasetsRequest {
	op := &aws.Operation{
		Name:       opListDatasets,
		HTTPMethod: "GET",
		HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets",
	}

	if input == nil {
		input = &types.ListDatasetsInput{}
	}

	req := c.newRequest(op, input, &types.ListDatasetsOutput{})
	return ListDatasetsRequest{Request: req, Input: input, Copy: c.ListDatasetsRequest}
}

// ListDatasetsRequest is the request type for the
// ListDatasets API operation.
type ListDatasetsRequest struct {
	*aws.Request
	Input *types.ListDatasetsInput
	Copy  func(*types.ListDatasetsInput) ListDatasetsRequest
}

// Send marshals and sends the ListDatasets API request.
func (r ListDatasetsRequest) Send(ctx context.Context) (*ListDatasetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDatasetsResponse{
		ListDatasetsOutput: r.Request.Data.(*types.ListDatasetsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDatasetsResponse is the response type for the
// ListDatasets API operation.
type ListDatasetsResponse struct {
	*types.ListDatasetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDatasets request.
func (r *ListDatasetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
