// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opListGroupVersions = "ListGroupVersions"

// ListGroupVersionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Lists the versions of a group.
//
//    // Example sending a request using ListGroupVersionsRequest.
//    req := client.ListGroupVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupVersions
func (c *Client) ListGroupVersionsRequest(input *types.ListGroupVersionsInput) ListGroupVersionsRequest {
	op := &aws.Operation{
		Name:       opListGroupVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/versions",
	}

	if input == nil {
		input = &types.ListGroupVersionsInput{}
	}

	req := c.newRequest(op, input, &types.ListGroupVersionsOutput{})
	return ListGroupVersionsRequest{Request: req, Input: input, Copy: c.ListGroupVersionsRequest}
}

// ListGroupVersionsRequest is the request type for the
// ListGroupVersions API operation.
type ListGroupVersionsRequest struct {
	*aws.Request
	Input *types.ListGroupVersionsInput
	Copy  func(*types.ListGroupVersionsInput) ListGroupVersionsRequest
}

// Send marshals and sends the ListGroupVersions API request.
func (r ListGroupVersionsRequest) Send(ctx context.Context) (*ListGroupVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupVersionsResponse{
		ListGroupVersionsOutput: r.Request.Data.(*types.ListGroupVersionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGroupVersionsResponse is the response type for the
// ListGroupVersions API operation.
type ListGroupVersionsResponse struct {
	*types.ListGroupVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroupVersions request.
func (r *ListGroupVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
