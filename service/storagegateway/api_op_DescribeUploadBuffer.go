// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDescribeUploadBuffer = "DescribeUploadBuffer"

// DescribeUploadBufferRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns information about the upload buffer of a gateway. This operation
// is supported for the stored volume, cached volume and tape gateway types.
//
// The response includes disk IDs that are configured as upload buffer space,
// and it includes the amount of upload buffer space allocated and used.
//
//    // Example sending a request using DescribeUploadBufferRequest.
//    req := client.DescribeUploadBufferRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeUploadBuffer
func (c *Client) DescribeUploadBufferRequest(input *types.DescribeUploadBufferInput) DescribeUploadBufferRequest {
	op := &aws.Operation{
		Name:       opDescribeUploadBuffer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeUploadBufferInput{}
	}

	req := c.newRequest(op, input, &types.DescribeUploadBufferOutput{})
	return DescribeUploadBufferRequest{Request: req, Input: input, Copy: c.DescribeUploadBufferRequest}
}

// DescribeUploadBufferRequest is the request type for the
// DescribeUploadBuffer API operation.
type DescribeUploadBufferRequest struct {
	*aws.Request
	Input *types.DescribeUploadBufferInput
	Copy  func(*types.DescribeUploadBufferInput) DescribeUploadBufferRequest
}

// Send marshals and sends the DescribeUploadBuffer API request.
func (r DescribeUploadBufferRequest) Send(ctx context.Context) (*DescribeUploadBufferResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUploadBufferResponse{
		DescribeUploadBufferOutput: r.Request.Data.(*types.DescribeUploadBufferOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUploadBufferResponse is the response type for the
// DescribeUploadBuffer API operation.
type DescribeUploadBufferResponse struct {
	*types.DescribeUploadBufferOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUploadBuffer request.
func (r *DescribeUploadBufferResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
