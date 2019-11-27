// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opDeleteBucketCors = "DeleteBucketCors"

// DeleteBucketCorsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the CORS configuration information set for the bucket.
//
//    // Example sending a request using DeleteBucketCorsRequest.
//    req := client.DeleteBucketCorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketCors
func (c *Client) DeleteBucketCorsRequest(input *types.DeleteBucketCorsInput) DeleteBucketCorsRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketCors,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?cors",
	}

	if input == nil {
		input = &types.DeleteBucketCorsInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBucketCorsOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketCorsRequest{Request: req, Input: input, Copy: c.DeleteBucketCorsRequest}
}

// DeleteBucketCorsRequest is the request type for the
// DeleteBucketCors API operation.
type DeleteBucketCorsRequest struct {
	*aws.Request
	Input *types.DeleteBucketCorsInput
	Copy  func(*types.DeleteBucketCorsInput) DeleteBucketCorsRequest
}

// Send marshals and sends the DeleteBucketCors API request.
func (r DeleteBucketCorsRequest) Send(ctx context.Context) (*DeleteBucketCorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketCorsResponse{
		DeleteBucketCorsOutput: r.Request.Data.(*types.DeleteBucketCorsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketCorsResponse is the response type for the
// DeleteBucketCors API operation.
type DeleteBucketCorsResponse struct {
	*types.DeleteBucketCorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketCors request.
func (r *DeleteBucketCorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
