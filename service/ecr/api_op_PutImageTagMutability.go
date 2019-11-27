// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

const opPutImageTagMutability = "PutImageTagMutability"

// PutImageTagMutabilityRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Updates the image tag mutability settings for a repository.
//
//    // Example sending a request using PutImageTagMutabilityRequest.
//    req := client.PutImageTagMutabilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutImageTagMutability
func (c *Client) PutImageTagMutabilityRequest(input *types.PutImageTagMutabilityInput) PutImageTagMutabilityRequest {
	op := &aws.Operation{
		Name:       opPutImageTagMutability,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutImageTagMutabilityInput{}
	}

	req := c.newRequest(op, input, &types.PutImageTagMutabilityOutput{})
	return PutImageTagMutabilityRequest{Request: req, Input: input, Copy: c.PutImageTagMutabilityRequest}
}

// PutImageTagMutabilityRequest is the request type for the
// PutImageTagMutability API operation.
type PutImageTagMutabilityRequest struct {
	*aws.Request
	Input *types.PutImageTagMutabilityInput
	Copy  func(*types.PutImageTagMutabilityInput) PutImageTagMutabilityRequest
}

// Send marshals and sends the PutImageTagMutability API request.
func (r PutImageTagMutabilityRequest) Send(ctx context.Context) (*PutImageTagMutabilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutImageTagMutabilityResponse{
		PutImageTagMutabilityOutput: r.Request.Data.(*types.PutImageTagMutabilityOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutImageTagMutabilityResponse is the response type for the
// PutImageTagMutability API operation.
type PutImageTagMutabilityResponse struct {
	*types.PutImageTagMutabilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutImageTagMutability request.
func (r *PutImageTagMutabilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
