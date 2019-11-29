// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

const opGetPublicKey = "GetPublicKey2019_03_26"

// GetPublicKeyRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the public key information.
//
//    // Example sending a request using GetPublicKeyRequest.
//    req := client.GetPublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetPublicKey
func (c *Client) GetPublicKeyRequest(input *types.GetPublicKeyInput) GetPublicKeyRequest {
	op := &aws.Operation{
		Name:       opGetPublicKey,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/public-key/{Id}",
	}

	if input == nil {
		input = &types.GetPublicKeyInput{}
	}

	req := c.newRequest(op, input, &types.GetPublicKeyOutput{})
	return GetPublicKeyRequest{Request: req, Input: input, Copy: c.GetPublicKeyRequest}
}

// GetPublicKeyRequest is the request type for the
// GetPublicKey API operation.
type GetPublicKeyRequest struct {
	*aws.Request
	Input *types.GetPublicKeyInput
	Copy  func(*types.GetPublicKeyInput) GetPublicKeyRequest
}

// Send marshals and sends the GetPublicKey API request.
func (r GetPublicKeyRequest) Send(ctx context.Context) (*GetPublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPublicKeyResponse{
		GetPublicKeyOutput: r.Request.Data.(*types.GetPublicKeyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPublicKeyResponse is the response type for the
// GetPublicKey API operation.
type GetPublicKeyResponse struct {
	*types.GetPublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPublicKey request.
func (r *GetPublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
