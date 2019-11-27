// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

const opCreateInvalidation = "CreateInvalidation2019_03_26"

// CreateInvalidationRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Create a new invalidation.
//
//    // Example sending a request using CreateInvalidationRequest.
//    req := client.CreateInvalidationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateInvalidation
func (c *Client) CreateInvalidationRequest(input *types.CreateInvalidationInput) CreateInvalidationRequest {
	op := &aws.Operation{
		Name:       opCreateInvalidation,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-03-26/distribution/{DistributionId}/invalidation",
	}

	if input == nil {
		input = &types.CreateInvalidationInput{}
	}

	req := c.newRequest(op, input, &types.CreateInvalidationOutput{})
	return CreateInvalidationRequest{Request: req, Input: input, Copy: c.CreateInvalidationRequest}
}

// CreateInvalidationRequest is the request type for the
// CreateInvalidation API operation.
type CreateInvalidationRequest struct {
	*aws.Request
	Input *types.CreateInvalidationInput
	Copy  func(*types.CreateInvalidationInput) CreateInvalidationRequest
}

// Send marshals and sends the CreateInvalidation API request.
func (r CreateInvalidationRequest) Send(ctx context.Context) (*CreateInvalidationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInvalidationResponse{
		CreateInvalidationOutput: r.Request.Data.(*types.CreateInvalidationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInvalidationResponse is the response type for the
// CreateInvalidation API operation.
type CreateInvalidationResponse struct {
	*types.CreateInvalidationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInvalidation request.
func (r *CreateInvalidationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
