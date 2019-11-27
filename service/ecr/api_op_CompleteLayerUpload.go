// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

const opCompleteLayerUpload = "CompleteLayerUpload"

// CompleteLayerUploadRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Informs Amazon ECR that the image layer upload has completed for a specified
// registry, repository name, and upload ID. You can optionally provide a sha256
// digest of the image layer for data validation purposes.
//
// This operation is used by the Amazon ECR proxy, and it is not intended for
// general use by customers for pulling and pushing images. In most cases, you
// should use the docker CLI to pull, tag, and push images.
//
//    // Example sending a request using CompleteLayerUploadRequest.
//    req := client.CompleteLayerUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/CompleteLayerUpload
func (c *Client) CompleteLayerUploadRequest(input *types.CompleteLayerUploadInput) CompleteLayerUploadRequest {
	op := &aws.Operation{
		Name:       opCompleteLayerUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CompleteLayerUploadInput{}
	}

	req := c.newRequest(op, input, &types.CompleteLayerUploadOutput{})
	return CompleteLayerUploadRequest{Request: req, Input: input, Copy: c.CompleteLayerUploadRequest}
}

// CompleteLayerUploadRequest is the request type for the
// CompleteLayerUpload API operation.
type CompleteLayerUploadRequest struct {
	*aws.Request
	Input *types.CompleteLayerUploadInput
	Copy  func(*types.CompleteLayerUploadInput) CompleteLayerUploadRequest
}

// Send marshals and sends the CompleteLayerUpload API request.
func (r CompleteLayerUploadRequest) Send(ctx context.Context) (*CompleteLayerUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CompleteLayerUploadResponse{
		CompleteLayerUploadOutput: r.Request.Data.(*types.CompleteLayerUploadOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CompleteLayerUploadResponse is the response type for the
// CompleteLayerUpload API operation.
type CompleteLayerUploadResponse struct {
	*types.CompleteLayerUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CompleteLayerUpload request.
func (r *CompleteLayerUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
