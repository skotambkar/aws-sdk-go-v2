// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/signer/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
)

const opUntagResource = "UntagResource"

// UntagResourceRequest returns a request value for making API operation for
// AWS Signer.
//
// Remove one or more tags from a signing profile. Specify a list of tag keys
// to remove the tags.
//
//    // Example sending a request using UntagResourceRequest.
//    req := client.UntagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/UntagResource
func (c *Client) UntagResourceRequest(input *types.UntagResourceInput) UntagResourceRequest {
	op := &aws.Operation{
		Name:       opUntagResource,
		HTTPMethod: "DELETE",
		HTTPPath:   "/tags/{resourceArn}",
	}

	if input == nil {
		input = &types.UntagResourceInput{}
	}

	req := c.newRequest(op, input, &types.UntagResourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UntagResourceMarshaler{Input: input}.GetNamedBuildHandler())

	return UntagResourceRequest{Request: req, Input: input, Copy: c.UntagResourceRequest}
}

// UntagResourceRequest is the request type for the
// UntagResource API operation.
type UntagResourceRequest struct {
	*aws.Request
	Input *types.UntagResourceInput
	Copy  func(*types.UntagResourceInput) UntagResourceRequest
}

// Send marshals and sends the UntagResource API request.
func (r UntagResourceRequest) Send(ctx context.Context) (*UntagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagResourceResponse{
		UntagResourceOutput: r.Request.Data.(*types.UntagResourceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagResourceResponse is the response type for the
// UntagResource API operation.
type UntagResourceResponse struct {
	*types.UntagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagResource request.
func (r *UntagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
