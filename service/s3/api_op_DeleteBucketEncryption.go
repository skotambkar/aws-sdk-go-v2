// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opDeleteBucketEncryption = "DeleteBucketEncryption"

// DeleteBucketEncryptionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This implementation of the DELETE operation removes default encryption from
// the bucket. For information about the Amazon S3 default encryption feature,
// see Amazon S3 Default Bucket Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev//bucket-encryption.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// To use this operation, you must have permissions to perform the s3:PutEncryptionConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev//using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev//s3-access-control.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// Related Resources
//
//    * PutBucketEncryption
//
//    * GetBucketEncryption
//
//    // Example sending a request using DeleteBucketEncryptionRequest.
//    req := client.DeleteBucketEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketEncryption
func (c *Client) DeleteBucketEncryptionRequest(input *types.DeleteBucketEncryptionInput) DeleteBucketEncryptionRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketEncryption,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &types.DeleteBucketEncryptionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBucketEncryptionOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketEncryptionRequest{Request: req, Input: input, Copy: c.DeleteBucketEncryptionRequest}
}

// DeleteBucketEncryptionRequest is the request type for the
// DeleteBucketEncryption API operation.
type DeleteBucketEncryptionRequest struct {
	*aws.Request
	Input *types.DeleteBucketEncryptionInput
	Copy  func(*types.DeleteBucketEncryptionInput) DeleteBucketEncryptionRequest
}

// Send marshals and sends the DeleteBucketEncryption API request.
func (r DeleteBucketEncryptionRequest) Send(ctx context.Context) (*DeleteBucketEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketEncryptionResponse{
		DeleteBucketEncryptionOutput: r.Request.Data.(*types.DeleteBucketEncryptionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketEncryptionResponse is the response type for the
// DeleteBucketEncryption API operation.
type DeleteBucketEncryptionResponse struct {
	*types.DeleteBucketEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketEncryption request.
func (r *DeleteBucketEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
