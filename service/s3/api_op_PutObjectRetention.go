// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opPutObjectRetention = "PutObjectRetention"

// PutObjectRetentionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Places an Object Retention configuration on an object.
//
// Related Resources
//
//    * Locking Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html)
//
//    // Example sending a request using PutObjectRetentionRequest.
//    req := client.PutObjectRetentionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectRetention
func (c *Client) PutObjectRetentionRequest(input *types.PutObjectRetentionInput) PutObjectRetentionRequest {
	op := &aws.Operation{
		Name:       opPutObjectRetention,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?retention",
	}

	if input == nil {
		input = &types.PutObjectRetentionInput{}
	}

	req := c.newRequest(op, input, &types.PutObjectRetentionOutput{})
	return PutObjectRetentionRequest{Request: req, Input: input, Copy: c.PutObjectRetentionRequest}
}

// PutObjectRetentionRequest is the request type for the
// PutObjectRetention API operation.
type PutObjectRetentionRequest struct {
	*aws.Request
	Input *types.PutObjectRetentionInput
	Copy  func(*types.PutObjectRetentionInput) PutObjectRetentionRequest
}

// Send marshals and sends the PutObjectRetention API request.
func (r PutObjectRetentionRequest) Send(ctx context.Context) (*PutObjectRetentionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectRetentionResponse{
		PutObjectRetentionOutput: r.Request.Data.(*types.PutObjectRetentionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectRetentionResponse is the response type for the
// PutObjectRetention API operation.
type PutObjectRetentionResponse struct {
	*types.PutObjectRetentionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObjectRetention request.
func (r *PutObjectRetentionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
