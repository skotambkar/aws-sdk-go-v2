// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opGetTags = "GetTags"

// GetTagsRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the Tags for a resource.
//
//    // Example sending a request using GetTagsRequest.
//    req := client.GetTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetTags
func (c *Client) GetTagsRequest(input *types.GetTagsInput) GetTagsRequest {
	op := &aws.Operation{
		Name:       opGetTags,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/tags/{resource-arn}",
	}

	if input == nil {
		input = &types.GetTagsInput{}
	}

	req := c.newRequest(op, input, &types.GetTagsOutput{})
	return GetTagsRequest{Request: req, Input: input, Copy: c.GetTagsRequest}
}

// GetTagsRequest is the request type for the
// GetTags API operation.
type GetTagsRequest struct {
	*aws.Request
	Input *types.GetTagsInput
	Copy  func(*types.GetTagsInput) GetTagsRequest
}

// Send marshals and sends the GetTags API request.
func (r GetTagsRequest) Send(ctx context.Context) (*GetTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTagsResponse{
		GetTagsOutput: r.Request.Data.(*types.GetTagsOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTagsResponse is the response type for the
// GetTags API operation.
type GetTagsResponse struct {
	*types.GetTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTags request.
func (r *GetTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
