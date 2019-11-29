// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
)

const opUpdateFileSystem = "UpdateFileSystem"

// UpdateFileSystemRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Updates the throughput mode or the amount of provisioned throughput of an
// existing file system.
//
//    // Example sending a request using UpdateFileSystemRequest.
//    req := client.UpdateFileSystemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/UpdateFileSystem
func (c *Client) UpdateFileSystemRequest(input *types.UpdateFileSystemInput) UpdateFileSystemRequest {
	op := &aws.Operation{
		Name:       opUpdateFileSystem,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}",
	}

	if input == nil {
		input = &types.UpdateFileSystemInput{}
	}

	req := c.newRequest(op, input, &types.UpdateFileSystemOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateFileSystemMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateFileSystemRequest{Request: req, Input: input, Copy: c.UpdateFileSystemRequest}
}

// UpdateFileSystemRequest is the request type for the
// UpdateFileSystem API operation.
type UpdateFileSystemRequest struct {
	*aws.Request
	Input *types.UpdateFileSystemInput
	Copy  func(*types.UpdateFileSystemInput) UpdateFileSystemRequest
}

// Send marshals and sends the UpdateFileSystem API request.
func (r UpdateFileSystemRequest) Send(ctx context.Context) (*UpdateFileSystemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFileSystemResponse{
		UpdateFileSystemOutput: r.Request.Data.(*types.UpdateFileSystemOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFileSystemResponse is the response type for the
// UpdateFileSystem API operation.
type UpdateFileSystemResponse struct {
	*types.UpdateFileSystemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFileSystem request.
func (r *UpdateFileSystemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
