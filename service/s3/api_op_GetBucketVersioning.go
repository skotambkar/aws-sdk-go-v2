// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetBucketVersioning = "GetBucketVersioning"

// GetBucketVersioningRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the versioning state of a bucket.
//
// To retrieve the versioning state of a bucket, you must be the bucket owner.
//
// This implementation also returns the MFA Delete status of the versioning
// state, i.e., if the MFA Delete status is enabled, the bucket owner must use
// an authentication device to change the versioning state of the bucket.
//
// The following operations are related to GetBucketVersioning:
//
//    * GetObject
//
//    * PutObject
//
//    * DeleteObject
//
//    // Example sending a request using GetBucketVersioningRequest.
//    req := client.GetBucketVersioningRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketVersioning
func (c *Client) GetBucketVersioningRequest(input *types.GetBucketVersioningInput) GetBucketVersioningRequest {
	op := &aws.Operation{
		Name:       opGetBucketVersioning,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?versioning",
	}

	if input == nil {
		input = &types.GetBucketVersioningInput{}
	}

	req := c.newRequest(op, input, &types.GetBucketVersioningOutput{})
	return GetBucketVersioningRequest{Request: req, Input: input, Copy: c.GetBucketVersioningRequest}
}

// GetBucketVersioningRequest is the request type for the
// GetBucketVersioning API operation.
type GetBucketVersioningRequest struct {
	*aws.Request
	Input *types.GetBucketVersioningInput
	Copy  func(*types.GetBucketVersioningInput) GetBucketVersioningRequest
}

// Send marshals and sends the GetBucketVersioning API request.
func (r GetBucketVersioningRequest) Send(ctx context.Context) (*GetBucketVersioningResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketVersioningResponse{
		GetBucketVersioningOutput: r.Request.Data.(*types.GetBucketVersioningOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketVersioningResponse is the response type for the
// GetBucketVersioning API operation.
type GetBucketVersioningResponse struct {
	*types.GetBucketVersioningOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketVersioning request.
func (r *GetBucketVersioningResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
