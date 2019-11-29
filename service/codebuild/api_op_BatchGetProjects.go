// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
)

const opBatchGetProjects = "BatchGetProjects"

// BatchGetProjectsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Gets information about one or more build projects.
//
//    // Example sending a request using BatchGetProjectsRequest.
//    req := client.BatchGetProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/BatchGetProjects
func (c *Client) BatchGetProjectsRequest(input *types.BatchGetProjectsInput) BatchGetProjectsRequest {
	op := &aws.Operation{
		Name:       opBatchGetProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchGetProjectsInput{}
	}

	req := c.newRequest(op, input, &types.BatchGetProjectsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.BatchGetProjectsMarshaler{Input: input}.GetNamedBuildHandler())

	return BatchGetProjectsRequest{Request: req, Input: input, Copy: c.BatchGetProjectsRequest}
}

// BatchGetProjectsRequest is the request type for the
// BatchGetProjects API operation.
type BatchGetProjectsRequest struct {
	*aws.Request
	Input *types.BatchGetProjectsInput
	Copy  func(*types.BatchGetProjectsInput) BatchGetProjectsRequest
}

// Send marshals and sends the BatchGetProjects API request.
func (r BatchGetProjectsRequest) Send(ctx context.Context) (*BatchGetProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetProjectsResponse{
		BatchGetProjectsOutput: r.Request.Data.(*types.BatchGetProjectsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetProjectsResponse is the response type for the
// BatchGetProjects API operation.
type BatchGetProjectsResponse struct {
	*types.BatchGetProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetProjects request.
func (r *BatchGetProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
