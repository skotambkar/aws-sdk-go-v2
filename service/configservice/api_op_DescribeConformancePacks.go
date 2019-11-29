// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDescribeConformancePacks = "DescribeConformancePacks"

// DescribeConformancePacksRequest returns a request value for making API operation for
// AWS Config.
//
// Returns a list of one or more conformance packs.
//
//    // Example sending a request using DescribeConformancePacksRequest.
//    req := client.DescribeConformancePacksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConformancePacks
func (c *Client) DescribeConformancePacksRequest(input *types.DescribeConformancePacksInput) DescribeConformancePacksRequest {
	op := &aws.Operation{
		Name:       opDescribeConformancePacks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeConformancePacksInput{}
	}

	req := c.newRequest(op, input, &types.DescribeConformancePacksOutput{})
	return DescribeConformancePacksRequest{Request: req, Input: input, Copy: c.DescribeConformancePacksRequest}
}

// DescribeConformancePacksRequest is the request type for the
// DescribeConformancePacks API operation.
type DescribeConformancePacksRequest struct {
	*aws.Request
	Input *types.DescribeConformancePacksInput
	Copy  func(*types.DescribeConformancePacksInput) DescribeConformancePacksRequest
}

// Send marshals and sends the DescribeConformancePacks API request.
func (r DescribeConformancePacksRequest) Send(ctx context.Context) (*DescribeConformancePacksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConformancePacksResponse{
		DescribeConformancePacksOutput: r.Request.Data.(*types.DescribeConformancePacksOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConformancePacksResponse is the response type for the
// DescribeConformancePacks API operation.
type DescribeConformancePacksResponse struct {
	*types.DescribeConformancePacksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConformancePacks request.
func (r *DescribeConformancePacksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
