// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
)

const opImportWorkspaceImage = "ImportWorkspaceImage"

// ImportWorkspaceImageRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Imports the specified Windows 7 or Windows 10 Bring Your Own License (BYOL)
// image into Amazon WorkSpaces. The image must be an already licensed EC2 image
// that is in your AWS account, and you must own the image.
//
//    // Example sending a request using ImportWorkspaceImageRequest.
//    req := client.ImportWorkspaceImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ImportWorkspaceImage
func (c *Client) ImportWorkspaceImageRequest(input *types.ImportWorkspaceImageInput) ImportWorkspaceImageRequest {
	op := &aws.Operation{
		Name:       opImportWorkspaceImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ImportWorkspaceImageInput{}
	}

	req := c.newRequest(op, input, &types.ImportWorkspaceImageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ImportWorkspaceImageMarshaler{Input: input}.GetNamedBuildHandler())

	return ImportWorkspaceImageRequest{Request: req, Input: input, Copy: c.ImportWorkspaceImageRequest}
}

// ImportWorkspaceImageRequest is the request type for the
// ImportWorkspaceImage API operation.
type ImportWorkspaceImageRequest struct {
	*aws.Request
	Input *types.ImportWorkspaceImageInput
	Copy  func(*types.ImportWorkspaceImageInput) ImportWorkspaceImageRequest
}

// Send marshals and sends the ImportWorkspaceImage API request.
func (r ImportWorkspaceImageRequest) Send(ctx context.Context) (*ImportWorkspaceImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportWorkspaceImageResponse{
		ImportWorkspaceImageOutput: r.Request.Data.(*types.ImportWorkspaceImageOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportWorkspaceImageResponse is the response type for the
// ImportWorkspaceImage API operation.
type ImportWorkspaceImageResponse struct {
	*types.ImportWorkspaceImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportWorkspaceImage request.
func (r *ImportWorkspaceImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
