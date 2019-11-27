// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opDescribeProductAsAdmin = "DescribeProductAsAdmin"

// DescribeProductAsAdminRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified product. This operation is run with
// administrator access.
//
//    // Example sending a request using DescribeProductAsAdminRequest.
//    req := client.DescribeProductAsAdminRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProductAsAdmin
func (c *Client) DescribeProductAsAdminRequest(input *types.DescribeProductAsAdminInput) DescribeProductAsAdminRequest {
	op := &aws.Operation{
		Name:       opDescribeProductAsAdmin,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeProductAsAdminInput{}
	}

	req := c.newRequest(op, input, &types.DescribeProductAsAdminOutput{})
	return DescribeProductAsAdminRequest{Request: req, Input: input, Copy: c.DescribeProductAsAdminRequest}
}

// DescribeProductAsAdminRequest is the request type for the
// DescribeProductAsAdmin API operation.
type DescribeProductAsAdminRequest struct {
	*aws.Request
	Input *types.DescribeProductAsAdminInput
	Copy  func(*types.DescribeProductAsAdminInput) DescribeProductAsAdminRequest
}

// Send marshals and sends the DescribeProductAsAdmin API request.
func (r DescribeProductAsAdminRequest) Send(ctx context.Context) (*DescribeProductAsAdminResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProductAsAdminResponse{
		DescribeProductAsAdminOutput: r.Request.Data.(*types.DescribeProductAsAdminOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProductAsAdminResponse is the response type for the
// DescribeProductAsAdmin API operation.
type DescribeProductAsAdminResponse struct {
	*types.DescribeProductAsAdminOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProductAsAdmin request.
func (r *DescribeProductAsAdminResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
