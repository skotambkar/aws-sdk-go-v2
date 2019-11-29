// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
)

const opStopBuild = "StopBuild"

// StopBuildRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Attempts to stop running a build.
//
//    // Example sending a request using StopBuildRequest.
//    req := client.StopBuildRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/StopBuild
func (c *Client) StopBuildRequest(input *types.StopBuildInput) StopBuildRequest {
	op := &aws.Operation{
		Name:       opStopBuild,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopBuildInput{}
	}

	req := c.newRequest(op, input, &types.StopBuildOutput{})
	return StopBuildRequest{Request: req, Input: input, Copy: c.StopBuildRequest}
}

// StopBuildRequest is the request type for the
// StopBuild API operation.
type StopBuildRequest struct {
	*aws.Request
	Input *types.StopBuildInput
	Copy  func(*types.StopBuildInput) StopBuildRequest
}

// Send marshals and sends the StopBuild API request.
func (r StopBuildRequest) Send(ctx context.Context) (*StopBuildResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopBuildResponse{
		StopBuildOutput: r.Request.Data.(*types.StopBuildOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopBuildResponse is the response type for the
// StopBuild API operation.
type StopBuildResponse struct {
	*types.StopBuildOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopBuild request.
func (r *StopBuildResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
