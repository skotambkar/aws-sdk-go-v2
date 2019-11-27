// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opDescribeStorage = "DescribeStorage"

// DescribeStorageRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns the total amount of snapshot usage and provisioned storage in megabytes.
//
//    // Example sending a request using DescribeStorageRequest.
//    req := client.DescribeStorageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeStorage
func (c *Client) DescribeStorageRequest(input *types.DescribeStorageInput) DescribeStorageRequest {
	op := &aws.Operation{
		Name:       opDescribeStorage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeStorageInput{}
	}

	req := c.newRequest(op, input, &types.DescribeStorageOutput{})
	return DescribeStorageRequest{Request: req, Input: input, Copy: c.DescribeStorageRequest}
}

// DescribeStorageRequest is the request type for the
// DescribeStorage API operation.
type DescribeStorageRequest struct {
	*aws.Request
	Input *types.DescribeStorageInput
	Copy  func(*types.DescribeStorageInput) DescribeStorageRequest
}

// Send marshals and sends the DescribeStorage API request.
func (r DescribeStorageRequest) Send(ctx context.Context) (*DescribeStorageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStorageResponse{
		DescribeStorageOutput: r.Request.Data.(*types.DescribeStorageOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStorageResponse is the response type for the
// DescribeStorage API operation.
type DescribeStorageResponse struct {
	*types.DescribeStorageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStorage request.
func (r *DescribeStorageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
