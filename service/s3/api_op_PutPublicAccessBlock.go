// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opPutPublicAccessBlock = "PutPublicAccessBlock"

// PutPublicAccessBlockRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates or modifies the PublicAccessBlock configuration for an Amazon S3
// bucket.
//
//    // Example sending a request using PutPublicAccessBlockRequest.
//    req := client.PutPublicAccessBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutPublicAccessBlock
func (c *Client) PutPublicAccessBlockRequest(input *types.PutPublicAccessBlockInput) PutPublicAccessBlockRequest {
	op := &aws.Operation{
		Name:       opPutPublicAccessBlock,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?publicAccessBlock",
	}

	if input == nil {
		input = &types.PutPublicAccessBlockInput{}
	}

	req := c.newRequest(op, input, &types.PutPublicAccessBlockOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutPublicAccessBlockRequest{Request: req, Input: input, Copy: c.PutPublicAccessBlockRequest}
}

// PutPublicAccessBlockRequest is the request type for the
// PutPublicAccessBlock API operation.
type PutPublicAccessBlockRequest struct {
	*aws.Request
	Input *types.PutPublicAccessBlockInput
	Copy  func(*types.PutPublicAccessBlockInput) PutPublicAccessBlockRequest
}

// Send marshals and sends the PutPublicAccessBlock API request.
func (r PutPublicAccessBlockRequest) Send(ctx context.Context) (*PutPublicAccessBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPublicAccessBlockResponse{
		PutPublicAccessBlockOutput: r.Request.Data.(*types.PutPublicAccessBlockOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPublicAccessBlockResponse is the response type for the
// PutPublicAccessBlock API operation.
type PutPublicAccessBlockResponse struct {
	*types.PutPublicAccessBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPublicAccessBlock request.
func (r *PutPublicAccessBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
