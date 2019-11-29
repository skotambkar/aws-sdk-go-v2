// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetBucketEncryption = "GetBucketEncryption"

// GetBucketEncryptionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the default encryption configuration for an Amazon S3 bucket. For
// information about the Amazon S3 default encryption feature, see Amazon S3
// Default Bucket Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html).
//
// To use this operation, you must have permission to perform the s3:GetEncryptionConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// The following operations are related to GetBucketEncryption:
//
//    * PutBucketEncryption
//
//    * DeleteBucketEncryption
//
//    // Example sending a request using GetBucketEncryptionRequest.
//    req := client.GetBucketEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketEncryption
func (c *Client) GetBucketEncryptionRequest(input *types.GetBucketEncryptionInput) GetBucketEncryptionRequest {
	op := &aws.Operation{
		Name:       opGetBucketEncryption,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &types.GetBucketEncryptionInput{}
	}

	req := c.newRequest(op, input, &types.GetBucketEncryptionOutput{})
	return GetBucketEncryptionRequest{Request: req, Input: input, Copy: c.GetBucketEncryptionRequest}
}

// GetBucketEncryptionRequest is the request type for the
// GetBucketEncryption API operation.
type GetBucketEncryptionRequest struct {
	*aws.Request
	Input *types.GetBucketEncryptionInput
	Copy  func(*types.GetBucketEncryptionInput) GetBucketEncryptionRequest
}

// Send marshals and sends the GetBucketEncryption API request.
func (r GetBucketEncryptionRequest) Send(ctx context.Context) (*GetBucketEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketEncryptionResponse{
		GetBucketEncryptionOutput: r.Request.Data.(*types.GetBucketEncryptionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketEncryptionResponse is the response type for the
// GetBucketEncryption API operation.
type GetBucketEncryptionResponse struct {
	*types.GetBucketEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketEncryption request.
func (r *GetBucketEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
