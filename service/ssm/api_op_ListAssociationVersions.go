// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opListAssociationVersions = "ListAssociationVersions"

// ListAssociationVersionsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves all versions of an association for a specific association ID.
//
//    // Example sending a request using ListAssociationVersionsRequest.
//    req := client.ListAssociationVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListAssociationVersions
func (c *Client) ListAssociationVersionsRequest(input *types.ListAssociationVersionsInput) ListAssociationVersionsRequest {
	op := &aws.Operation{
		Name:       opListAssociationVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListAssociationVersionsInput{}
	}

	req := c.newRequest(op, input, &types.ListAssociationVersionsOutput{})
	return ListAssociationVersionsRequest{Request: req, Input: input, Copy: c.ListAssociationVersionsRequest}
}

// ListAssociationVersionsRequest is the request type for the
// ListAssociationVersions API operation.
type ListAssociationVersionsRequest struct {
	*aws.Request
	Input *types.ListAssociationVersionsInput
	Copy  func(*types.ListAssociationVersionsInput) ListAssociationVersionsRequest
}

// Send marshals and sends the ListAssociationVersions API request.
func (r ListAssociationVersionsRequest) Send(ctx context.Context) (*ListAssociationVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssociationVersionsResponse{
		ListAssociationVersionsOutput: r.Request.Data.(*types.ListAssociationVersionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAssociationVersionsResponse is the response type for the
// ListAssociationVersions API operation.
type ListAssociationVersionsResponse struct {
	*types.ListAssociationVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssociationVersions request.
func (r *ListAssociationVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
