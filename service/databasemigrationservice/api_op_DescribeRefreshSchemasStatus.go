// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDescribeRefreshSchemasStatus = "DescribeRefreshSchemasStatus"

// DescribeRefreshSchemasStatusRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns the status of the RefreshSchemas operation.
//
//    // Example sending a request using DescribeRefreshSchemasStatusRequest.
//    req := client.DescribeRefreshSchemasStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeRefreshSchemasStatus
func (c *Client) DescribeRefreshSchemasStatusRequest(input *types.DescribeRefreshSchemasStatusInput) DescribeRefreshSchemasStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeRefreshSchemasStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeRefreshSchemasStatusInput{}
	}

	req := c.newRequest(op, input, &types.DescribeRefreshSchemasStatusOutput{})
	return DescribeRefreshSchemasStatusRequest{Request: req, Input: input, Copy: c.DescribeRefreshSchemasStatusRequest}
}

// DescribeRefreshSchemasStatusRequest is the request type for the
// DescribeRefreshSchemasStatus API operation.
type DescribeRefreshSchemasStatusRequest struct {
	*aws.Request
	Input *types.DescribeRefreshSchemasStatusInput
	Copy  func(*types.DescribeRefreshSchemasStatusInput) DescribeRefreshSchemasStatusRequest
}

// Send marshals and sends the DescribeRefreshSchemasStatus API request.
func (r DescribeRefreshSchemasStatusRequest) Send(ctx context.Context) (*DescribeRefreshSchemasStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRefreshSchemasStatusResponse{
		DescribeRefreshSchemasStatusOutput: r.Request.Data.(*types.DescribeRefreshSchemasStatusOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRefreshSchemasStatusResponse is the response type for the
// DescribeRefreshSchemasStatus API operation.
type DescribeRefreshSchemasStatusResponse struct {
	*types.DescribeRefreshSchemasStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRefreshSchemasStatus request.
func (r *DescribeRefreshSchemasStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
