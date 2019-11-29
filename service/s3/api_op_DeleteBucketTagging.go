// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opDeleteBucketTagging = "DeleteBucketTagging"

// DeleteBucketTaggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the tags from the bucket.
//
// To use this operation, you must have permission to perform the s3:PutBucketTagging
// action. By default, the bucket owner has this permission and can grant this
// permission to others.
//
// The following operations are related to DeleteBucketTagging
//
//    * GetBucketTagging
//
//    * PutBucketTagging
//
//    // Example sending a request using DeleteBucketTaggingRequest.
//    req := client.DeleteBucketTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketTagging
func (c *Client) DeleteBucketTaggingRequest(input *types.DeleteBucketTaggingInput) DeleteBucketTaggingRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketTagging,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?tagging",
	}

	if input == nil {
		input = &types.DeleteBucketTaggingInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBucketTaggingOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketTaggingRequest{Request: req, Input: input, Copy: c.DeleteBucketTaggingRequest}
}

// DeleteBucketTaggingRequest is the request type for the
// DeleteBucketTagging API operation.
type DeleteBucketTaggingRequest struct {
	*aws.Request
	Input *types.DeleteBucketTaggingInput
	Copy  func(*types.DeleteBucketTaggingInput) DeleteBucketTaggingRequest
}

// Send marshals and sends the DeleteBucketTagging API request.
func (r DeleteBucketTaggingRequest) Send(ctx context.Context) (*DeleteBucketTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketTaggingResponse{
		DeleteBucketTaggingOutput: r.Request.Data.(*types.DeleteBucketTaggingOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketTaggingResponse is the response type for the
// DeleteBucketTagging API operation.
type DeleteBucketTaggingResponse struct {
	*types.DeleteBucketTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketTagging request.
func (r *DeleteBucketTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
