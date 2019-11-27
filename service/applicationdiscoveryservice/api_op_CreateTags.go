// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
)

const opCreateTags = "CreateTags"

// CreateTagsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Creates one or more tags for configuration items. Tags are metadata that
// help you categorize IT assets. This API accepts a list of multiple configuration
// items.
//
//    // Example sending a request using CreateTagsRequest.
//    req := client.CreateTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/CreateTags
func (c *Client) CreateTagsRequest(input *types.CreateTagsInput) CreateTagsRequest {
	op := &aws.Operation{
		Name:       opCreateTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateTagsInput{}
	}

	req := c.newRequest(op, input, &types.CreateTagsOutput{})
	return CreateTagsRequest{Request: req, Input: input, Copy: c.CreateTagsRequest}
}

// CreateTagsRequest is the request type for the
// CreateTags API operation.
type CreateTagsRequest struct {
	*aws.Request
	Input *types.CreateTagsInput
	Copy  func(*types.CreateTagsInput) CreateTagsRequest
}

// Send marshals and sends the CreateTags API request.
func (r CreateTagsRequest) Send(ctx context.Context) (*CreateTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTagsResponse{
		CreateTagsOutput: r.Request.Data.(*types.CreateTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTagsResponse is the response type for the
// CreateTags API operation.
type CreateTagsResponse struct {
	*types.CreateTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTags request.
func (r *CreateTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
