// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opDeleteBucketReplication = "DeleteBucketReplication"

// DeleteBucketReplicationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the replication configuration from the bucket. For information about
// replication configuration, see Cross-Region Replication (CRR) (https://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html)
// in the Amazon S3 Developer Guide.
//
//    // Example sending a request using DeleteBucketReplicationRequest.
//    req := client.DeleteBucketReplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketReplication
func (c *Client) DeleteBucketReplicationRequest(input *types.DeleteBucketReplicationInput) DeleteBucketReplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketReplication,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?replication",
	}

	if input == nil {
		input = &types.DeleteBucketReplicationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBucketReplicationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketReplicationRequest{Request: req, Input: input, Copy: c.DeleteBucketReplicationRequest}
}

// DeleteBucketReplicationRequest is the request type for the
// DeleteBucketReplication API operation.
type DeleteBucketReplicationRequest struct {
	*aws.Request
	Input *types.DeleteBucketReplicationInput
	Copy  func(*types.DeleteBucketReplicationInput) DeleteBucketReplicationRequest
}

// Send marshals and sends the DeleteBucketReplication API request.
func (r DeleteBucketReplicationRequest) Send(ctx context.Context) (*DeleteBucketReplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketReplicationResponse{
		DeleteBucketReplicationOutput: r.Request.Data.(*types.DeleteBucketReplicationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketReplicationResponse is the response type for the
// DeleteBucketReplication API operation.
type DeleteBucketReplicationResponse struct {
	*types.DeleteBucketReplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketReplication request.
func (r *DeleteBucketReplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
