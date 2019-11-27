// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListThingTypes = "ListThingTypes"

// ListThingTypesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the existing thing types.
//
//    // Example sending a request using ListThingTypesRequest.
//    req := client.ListThingTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListThingTypesRequest(input *types.ListThingTypesInput) ListThingTypesRequest {
	op := &aws.Operation{
		Name:       opListThingTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/thing-types",
	}

	if input == nil {
		input = &types.ListThingTypesInput{}
	}

	req := c.newRequest(op, input, &types.ListThingTypesOutput{})
	return ListThingTypesRequest{Request: req, Input: input, Copy: c.ListThingTypesRequest}
}

// ListThingTypesRequest is the request type for the
// ListThingTypes API operation.
type ListThingTypesRequest struct {
	*aws.Request
	Input *types.ListThingTypesInput
	Copy  func(*types.ListThingTypesInput) ListThingTypesRequest
}

// Send marshals and sends the ListThingTypes API request.
func (r ListThingTypesRequest) Send(ctx context.Context) (*ListThingTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThingTypesResponse{
		ListThingTypesOutput: r.Request.Data.(*types.ListThingTypesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListThingTypesResponse is the response type for the
// ListThingTypes API operation.
type ListThingTypesResponse struct {
	*types.ListThingTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThingTypes request.
func (r *ListThingTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
