// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opUploadPartCopy = "UploadPartCopy"

// UploadPartCopyRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Uploads a part by copying data from an existing object as data source.
//
//    // Example sending a request using UploadPartCopyRequest.
//    req := client.UploadPartCopyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/UploadPartCopy
func (c *Client) UploadPartCopyRequest(input *types.UploadPartCopyInput) UploadPartCopyRequest {
	op := &aws.Operation{
		Name:       opUploadPartCopy,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &types.UploadPartCopyInput{}
	}

	req := c.newRequest(op, input, &types.UploadPartCopyOutput{})
	return UploadPartCopyRequest{Request: req, Input: input, Copy: c.UploadPartCopyRequest}
}

// UploadPartCopyRequest is the request type for the
// UploadPartCopy API operation.
type UploadPartCopyRequest struct {
	*aws.Request
	Input *types.UploadPartCopyInput
	Copy  func(*types.UploadPartCopyInput) UploadPartCopyRequest
}

// Send marshals and sends the UploadPartCopy API request.
func (r UploadPartCopyRequest) Send(ctx context.Context) (*UploadPartCopyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadPartCopyResponse{
		UploadPartCopyOutput: r.Request.Data.(*types.UploadPartCopyOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadPartCopyResponse is the response type for the
// UploadPartCopy API operation.
type UploadPartCopyResponse struct {
	*types.UploadPartCopyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadPartCopy request.
func (r *UploadPartCopyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
