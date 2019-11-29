// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
)

const opDescribeWorkspaceImages = "DescribeWorkspaceImages"

// DescribeWorkspaceImagesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Retrieves a list that describes one or more specified images, if the image
// identifiers are provided. Otherwise, all images in the account are described.
//
//    // Example sending a request using DescribeWorkspaceImagesRequest.
//    req := client.DescribeWorkspaceImagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeWorkspaceImages
func (c *Client) DescribeWorkspaceImagesRequest(input *types.DescribeWorkspaceImagesInput) DescribeWorkspaceImagesRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkspaceImages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeWorkspaceImagesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeWorkspaceImagesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeWorkspaceImagesMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeWorkspaceImagesRequest{Request: req, Input: input, Copy: c.DescribeWorkspaceImagesRequest}
}

// DescribeWorkspaceImagesRequest is the request type for the
// DescribeWorkspaceImages API operation.
type DescribeWorkspaceImagesRequest struct {
	*aws.Request
	Input *types.DescribeWorkspaceImagesInput
	Copy  func(*types.DescribeWorkspaceImagesInput) DescribeWorkspaceImagesRequest
}

// Send marshals and sends the DescribeWorkspaceImages API request.
func (r DescribeWorkspaceImagesRequest) Send(ctx context.Context) (*DescribeWorkspaceImagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkspaceImagesResponse{
		DescribeWorkspaceImagesOutput: r.Request.Data.(*types.DescribeWorkspaceImagesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWorkspaceImagesResponse is the response type for the
// DescribeWorkspaceImages API operation.
type DescribeWorkspaceImagesResponse struct {
	*types.DescribeWorkspaceImagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkspaceImages request.
func (r *DescribeWorkspaceImagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
