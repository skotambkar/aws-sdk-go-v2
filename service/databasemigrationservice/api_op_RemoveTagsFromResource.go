// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opRemoveTagsFromResource = "RemoveTagsFromResource"

// RemoveTagsFromResourceRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Removes metadata tags from a DMS resource.
//
//    // Example sending a request using RemoveTagsFromResourceRequest.
//    req := client.RemoveTagsFromResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/RemoveTagsFromResource
func (c *Client) RemoveTagsFromResourceRequest(input *types.RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest {
	op := &aws.Operation{
		Name:       opRemoveTagsFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RemoveTagsFromResourceInput{}
	}

	req := c.newRequest(op, input, &types.RemoveTagsFromResourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RemoveTagsFromResourceMarshaler{Input: input}.GetNamedBuildHandler())

	return RemoveTagsFromResourceRequest{Request: req, Input: input, Copy: c.RemoveTagsFromResourceRequest}
}

// RemoveTagsFromResourceRequest is the request type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceRequest struct {
	*aws.Request
	Input *types.RemoveTagsFromResourceInput
	Copy  func(*types.RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest
}

// Send marshals and sends the RemoveTagsFromResource API request.
func (r RemoveTagsFromResourceRequest) Send(ctx context.Context) (*RemoveTagsFromResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsFromResourceResponse{
		RemoveTagsFromResourceOutput: r.Request.Data.(*types.RemoveTagsFromResourceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsFromResourceResponse is the response type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceResponse struct {
	*types.RemoveTagsFromResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTagsFromResource request.
func (r *RemoveTagsFromResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
