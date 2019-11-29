// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
)

const opBatchDeleteBuilds = "BatchDeleteBuilds"

// BatchDeleteBuildsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Deletes one or more builds.
//
//    // Example sending a request using BatchDeleteBuildsRequest.
//    req := client.BatchDeleteBuildsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/BatchDeleteBuilds
func (c *Client) BatchDeleteBuildsRequest(input *types.BatchDeleteBuildsInput) BatchDeleteBuildsRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteBuilds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchDeleteBuildsInput{}
	}

	req := c.newRequest(op, input, &types.BatchDeleteBuildsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.BatchDeleteBuildsMarshaler{Input: input}.GetNamedBuildHandler())

	return BatchDeleteBuildsRequest{Request: req, Input: input, Copy: c.BatchDeleteBuildsRequest}
}

// BatchDeleteBuildsRequest is the request type for the
// BatchDeleteBuilds API operation.
type BatchDeleteBuildsRequest struct {
	*aws.Request
	Input *types.BatchDeleteBuildsInput
	Copy  func(*types.BatchDeleteBuildsInput) BatchDeleteBuildsRequest
}

// Send marshals and sends the BatchDeleteBuilds API request.
func (r BatchDeleteBuildsRequest) Send(ctx context.Context) (*BatchDeleteBuildsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteBuildsResponse{
		BatchDeleteBuildsOutput: r.Request.Data.(*types.BatchDeleteBuildsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteBuildsResponse is the response type for the
// BatchDeleteBuilds API operation.
type BatchDeleteBuildsResponse struct {
	*types.BatchDeleteBuildsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteBuilds request.
func (r *BatchDeleteBuildsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
