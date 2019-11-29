// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opDeletePublicAccessBlock = "DeletePublicAccessBlock"

// DeletePublicAccessBlockRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Removes the PublicAccessBlock configuration for an Amazon S3 bucket. In order
// to use this operation, you must have the s3:PutBucketPublicAccessBlock permission.
// For more information about permissions, see Permissions Related to Bucket
// Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// The following operations are related to DeleteBucketMetricsConfiguration:
//
//    * Using Amazon S3 Block Public Access (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html)
//
//    * GetPublicAccessBlock
//
//    * PutPublicAccessBlock
//
//    * GetBucketPolicyStatus
//
//    // Example sending a request using DeletePublicAccessBlockRequest.
//    req := client.DeletePublicAccessBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeletePublicAccessBlock
func (c *Client) DeletePublicAccessBlockRequest(input *types.DeletePublicAccessBlockInput) DeletePublicAccessBlockRequest {
	op := &aws.Operation{
		Name:       opDeletePublicAccessBlock,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?publicAccessBlock",
	}

	if input == nil {
		input = &types.DeletePublicAccessBlockInput{}
	}

	req := c.newRequest(op, input, &types.DeletePublicAccessBlockOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePublicAccessBlockRequest{Request: req, Input: input, Copy: c.DeletePublicAccessBlockRequest}
}

// DeletePublicAccessBlockRequest is the request type for the
// DeletePublicAccessBlock API operation.
type DeletePublicAccessBlockRequest struct {
	*aws.Request
	Input *types.DeletePublicAccessBlockInput
	Copy  func(*types.DeletePublicAccessBlockInput) DeletePublicAccessBlockRequest
}

// Send marshals and sends the DeletePublicAccessBlock API request.
func (r DeletePublicAccessBlockRequest) Send(ctx context.Context) (*DeletePublicAccessBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePublicAccessBlockResponse{
		DeletePublicAccessBlockOutput: r.Request.Data.(*types.DeletePublicAccessBlockOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePublicAccessBlockResponse is the response type for the
// DeletePublicAccessBlock API operation.
type DeletePublicAccessBlockResponse struct {
	*types.DeletePublicAccessBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePublicAccessBlock request.
func (r *DeletePublicAccessBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
