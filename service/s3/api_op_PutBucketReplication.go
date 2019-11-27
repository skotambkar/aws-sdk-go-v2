// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opPutBucketReplication = "PutBucketReplication"

// PutBucketReplicationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a replication configuration or replaces an existing one. For more
// information, see Cross-Region Replication (CRR) (https://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html)
// in the Amazon S3 Developer Guide.
//
//    // Example sending a request using PutBucketReplicationRequest.
//    req := client.PutBucketReplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketReplication
func (c *Client) PutBucketReplicationRequest(input *types.PutBucketReplicationInput) PutBucketReplicationRequest {
	op := &aws.Operation{
		Name:       opPutBucketReplication,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?replication",
	}

	if input == nil {
		input = &types.PutBucketReplicationInput{}
	}

	req := c.newRequest(op, input, &types.PutBucketReplicationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketReplicationRequest{Request: req, Input: input, Copy: c.PutBucketReplicationRequest}
}

// PutBucketReplicationRequest is the request type for the
// PutBucketReplication API operation.
type PutBucketReplicationRequest struct {
	*aws.Request
	Input *types.PutBucketReplicationInput
	Copy  func(*types.PutBucketReplicationInput) PutBucketReplicationRequest
}

// Send marshals and sends the PutBucketReplication API request.
func (r PutBucketReplicationRequest) Send(ctx context.Context) (*PutBucketReplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketReplicationResponse{
		PutBucketReplicationOutput: r.Request.Data.(*types.PutBucketReplicationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketReplicationResponse is the response type for the
// PutBucketReplication API operation.
type PutBucketReplicationResponse struct {
	*types.PutBucketReplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketReplication request.
func (r *PutBucketReplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
