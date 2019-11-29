// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDescribeDataSource = "DescribeDataSource"

// DescribeDataSourceRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes a data source.
//
// The permissions resource is arn:aws:quicksight:region:aws-account-id:datasource/data-source-id
//
//    // Example sending a request using DescribeDataSourceRequest.
//    req := client.DescribeDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeDataSource
func (c *Client) DescribeDataSourceRequest(input *types.DescribeDataSourceInput) DescribeDataSourceRequest {
	op := &aws.Operation{
		Name:       opDescribeDataSource,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sources/{DataSourceId}",
	}

	if input == nil {
		input = &types.DescribeDataSourceInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDataSourceOutput{})
	return DescribeDataSourceRequest{Request: req, Input: input, Copy: c.DescribeDataSourceRequest}
}

// DescribeDataSourceRequest is the request type for the
// DescribeDataSource API operation.
type DescribeDataSourceRequest struct {
	*aws.Request
	Input *types.DescribeDataSourceInput
	Copy  func(*types.DescribeDataSourceInput) DescribeDataSourceRequest
}

// Send marshals and sends the DescribeDataSource API request.
func (r DescribeDataSourceRequest) Send(ctx context.Context) (*DescribeDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDataSourceResponse{
		DescribeDataSourceOutput: r.Request.Data.(*types.DescribeDataSourceOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDataSourceResponse is the response type for the
// DescribeDataSource API operation.
type DescribeDataSourceResponse struct {
	*types.DescribeDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataSource request.
func (r *DescribeDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
