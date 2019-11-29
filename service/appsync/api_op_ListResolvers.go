// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
)

const opListResolvers = "ListResolvers"

// ListResolversRequest returns a request value for making API operation for
// AWS AppSync.
//
// Lists the resolvers for a given API and type.
//
//    // Example sending a request using ListResolversRequest.
//    req := client.ListResolversRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListResolvers
func (c *Client) ListResolversRequest(input *types.ListResolversInput) ListResolversRequest {
	op := &aws.Operation{
		Name:       opListResolvers,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/types/{typeName}/resolvers",
	}

	if input == nil {
		input = &types.ListResolversInput{}
	}

	req := c.newRequest(op, input, &types.ListResolversOutput{})
	return ListResolversRequest{Request: req, Input: input, Copy: c.ListResolversRequest}
}

// ListResolversRequest is the request type for the
// ListResolvers API operation.
type ListResolversRequest struct {
	*aws.Request
	Input *types.ListResolversInput
	Copy  func(*types.ListResolversInput) ListResolversRequest
}

// Send marshals and sends the ListResolvers API request.
func (r ListResolversRequest) Send(ctx context.Context) (*ListResolversResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResolversResponse{
		ListResolversOutput: r.Request.Data.(*types.ListResolversOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResolversResponse is the response type for the
// ListResolvers API operation.
type ListResolversResponse struct {
	*types.ListResolversOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResolvers request.
func (r *ListResolversResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
