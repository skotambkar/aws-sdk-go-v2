// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastoredata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata/types"
)

const opPutObject = "PutObject"

// PutObjectRequest returns a request value for making API operation for
// AWS Elemental MediaStore Data Plane.
//
// Uploads an object to the specified path. Object sizes are limited to 25 MB
// for standard upload availability and 10 MB for streaming upload availability.
//
//    // Example sending a request using PutObjectRequest.
//    req := client.PutObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-data-2017-09-01/PutObject
func (c *Client) PutObjectRequest(input *types.PutObjectInput) PutObjectRequest {
	op := &aws.Operation{
		Name:       opPutObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Path+}",
	}

	if input == nil {
		input = &types.PutObjectInput{}
	}

	req := c.newRequest(op, input, &types.PutObjectOutput{})
	req.Handlers.Sign.Remove(v4.SignRequestHandler)
	handler := v4.BuildNamedHandler("v4.CustomSignerHandler", v4.WithUnsignedPayload)
	req.Handlers.Sign.PushFrontNamed(handler)
	return PutObjectRequest{Request: req, Input: input, Copy: c.PutObjectRequest}
}

// PutObjectRequest is the request type for the
// PutObject API operation.
type PutObjectRequest struct {
	*aws.Request
	Input *types.PutObjectInput
	Copy  func(*types.PutObjectInput) PutObjectRequest
}

// Send marshals and sends the PutObject API request.
func (r PutObjectRequest) Send(ctx context.Context) (*PutObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectResponse{
		PutObjectOutput: r.Request.Data.(*types.PutObjectOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectResponse is the response type for the
// PutObject API operation.
type PutObjectResponse struct {
	*types.PutObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObject request.
func (r *PutObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
