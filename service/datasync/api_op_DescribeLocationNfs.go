// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
)

const opDescribeLocationNfs = "DescribeLocationNfs"

// DescribeLocationNfsRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns metadata, such as the path information, about a NFS location.
//
//    // Example sending a request using DescribeLocationNfsRequest.
//    req := client.DescribeLocationNfsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationNfs
func (c *Client) DescribeLocationNfsRequest(input *types.DescribeLocationNfsInput) DescribeLocationNfsRequest {
	op := &aws.Operation{
		Name:       opDescribeLocationNfs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeLocationNfsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeLocationNfsOutput{})
	return DescribeLocationNfsRequest{Request: req, Input: input, Copy: c.DescribeLocationNfsRequest}
}

// DescribeLocationNfsRequest is the request type for the
// DescribeLocationNfs API operation.
type DescribeLocationNfsRequest struct {
	*aws.Request
	Input *types.DescribeLocationNfsInput
	Copy  func(*types.DescribeLocationNfsInput) DescribeLocationNfsRequest
}

// Send marshals and sends the DescribeLocationNfs API request.
func (r DescribeLocationNfsRequest) Send(ctx context.Context) (*DescribeLocationNfsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLocationNfsResponse{
		DescribeLocationNfsOutput: r.Request.Data.(*types.DescribeLocationNfsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLocationNfsResponse is the response type for the
// DescribeLocationNfs API operation.
type DescribeLocationNfsResponse struct {
	*types.DescribeLocationNfsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLocationNfs request.
func (r *DescribeLocationNfsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
