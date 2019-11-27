// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opDescribeProduct = "DescribeProduct"

// DescribeProductRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified product.
//
//    // Example sending a request using DescribeProductRequest.
//    req := client.DescribeProductRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProduct
func (c *Client) DescribeProductRequest(input *types.DescribeProductInput) DescribeProductRequest {
	op := &aws.Operation{
		Name:       opDescribeProduct,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeProductInput{}
	}

	req := c.newRequest(op, input, &types.DescribeProductOutput{})
	return DescribeProductRequest{Request: req, Input: input, Copy: c.DescribeProductRequest}
}

// DescribeProductRequest is the request type for the
// DescribeProduct API operation.
type DescribeProductRequest struct {
	*aws.Request
	Input *types.DescribeProductInput
	Copy  func(*types.DescribeProductInput) DescribeProductRequest
}

// Send marshals and sends the DescribeProduct API request.
func (r DescribeProductRequest) Send(ctx context.Context) (*DescribeProductResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProductResponse{
		DescribeProductOutput: r.Request.Data.(*types.DescribeProductOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProductResponse is the response type for the
// DescribeProduct API operation.
type DescribeProductResponse struct {
	*types.DescribeProductOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProduct request.
func (r *DescribeProductResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
