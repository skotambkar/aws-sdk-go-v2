// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opPutBucketPolicy = "PutBucketPolicy"

// PutBucketPolicyRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Applies an Amazon S3 bucket policy to an Amazon S3 bucket.
//
//    // Example sending a request using PutBucketPolicyRequest.
//    req := client.PutBucketPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketPolicy
func (c *Client) PutBucketPolicyRequest(input *types.PutBucketPolicyInput) PutBucketPolicyRequest {
	op := &aws.Operation{
		Name:       opPutBucketPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?policy",
	}

	if input == nil {
		input = &types.PutBucketPolicyInput{}
	}

	req := c.newRequest(op, input, &types.PutBucketPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketPolicyRequest{Request: req, Input: input, Copy: c.PutBucketPolicyRequest}
}

// PutBucketPolicyRequest is the request type for the
// PutBucketPolicy API operation.
type PutBucketPolicyRequest struct {
	*aws.Request
	Input *types.PutBucketPolicyInput
	Copy  func(*types.PutBucketPolicyInput) PutBucketPolicyRequest
}

// Send marshals and sends the PutBucketPolicy API request.
func (r PutBucketPolicyRequest) Send(ctx context.Context) (*PutBucketPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketPolicyResponse{
		PutBucketPolicyOutput: r.Request.Data.(*types.PutBucketPolicyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketPolicyResponse is the response type for the
// PutBucketPolicy API operation.
type PutBucketPolicyResponse struct {
	*types.PutBucketPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketPolicy request.
func (r *PutBucketPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
