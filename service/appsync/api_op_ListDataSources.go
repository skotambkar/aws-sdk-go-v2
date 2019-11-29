// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
)

const opListDataSources = "ListDataSources"

// ListDataSourcesRequest returns a request value for making API operation for
// AWS AppSync.
//
// Lists the data sources for a given API.
//
//    // Example sending a request using ListDataSourcesRequest.
//    req := client.ListDataSourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListDataSources
func (c *Client) ListDataSourcesRequest(input *types.ListDataSourcesInput) ListDataSourcesRequest {
	op := &aws.Operation{
		Name:       opListDataSources,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/datasources",
	}

	if input == nil {
		input = &types.ListDataSourcesInput{}
	}

	req := c.newRequest(op, input, &types.ListDataSourcesOutput{})
	return ListDataSourcesRequest{Request: req, Input: input, Copy: c.ListDataSourcesRequest}
}

// ListDataSourcesRequest is the request type for the
// ListDataSources API operation.
type ListDataSourcesRequest struct {
	*aws.Request
	Input *types.ListDataSourcesInput
	Copy  func(*types.ListDataSourcesInput) ListDataSourcesRequest
}

// Send marshals and sends the ListDataSources API request.
func (r ListDataSourcesRequest) Send(ctx context.Context) (*ListDataSourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDataSourcesResponse{
		ListDataSourcesOutput: r.Request.Data.(*types.ListDataSourcesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDataSourcesResponse is the response type for the
// ListDataSources API operation.
type ListDataSourcesResponse struct {
	*types.ListDataSourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDataSources request.
func (r *ListDataSourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
