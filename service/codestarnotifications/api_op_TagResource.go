// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
)

const opTagResource = "TagResource"

// TagResourceRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Associates a set of provided tags with a notification rule.
//
//    // Example sending a request using TagResourceRequest.
//    req := client.TagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/TagResource
func (c *Client) TagResourceRequest(input *types.TagResourceInput) TagResourceRequest {
	op := &aws.Operation{
		Name:       opTagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/tagResource",
	}

	if input == nil {
		input = &types.TagResourceInput{}
	}

	req := c.newRequest(op, input, &types.TagResourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.TagResourceMarshaler{Input: input}.GetNamedBuildHandler())

	return TagResourceRequest{Request: req, Input: input, Copy: c.TagResourceRequest}
}

// TagResourceRequest is the request type for the
// TagResource API operation.
type TagResourceRequest struct {
	*aws.Request
	Input *types.TagResourceInput
	Copy  func(*types.TagResourceInput) TagResourceRequest
}

// Send marshals and sends the TagResource API request.
func (r TagResourceRequest) Send(ctx context.Context) (*TagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourceResponse{
		TagResourceOutput: r.Request.Data.(*types.TagResourceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourceResponse is the response type for the
// TagResource API operation.
type TagResourceResponse struct {
	*types.TagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResource request.
func (r *TagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
