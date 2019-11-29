// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeFleetHistory = "DescribeFleetHistory"

// DescribeFleetHistoryRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the events for the specified EC2 Fleet during the specified time.
//
//    // Example sending a request using DescribeFleetHistoryRequest.
//    req := client.DescribeFleetHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleetHistory
func (c *Client) DescribeFleetHistoryRequest(input *types.DescribeFleetHistoryInput) DescribeFleetHistoryRequest {
	op := &aws.Operation{
		Name:       opDescribeFleetHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeFleetHistoryInput{}
	}

	req := c.newRequest(op, input, &types.DescribeFleetHistoryOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DescribeFleetHistoryMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeFleetHistoryRequest{Request: req, Input: input, Copy: c.DescribeFleetHistoryRequest}
}

// DescribeFleetHistoryRequest is the request type for the
// DescribeFleetHistory API operation.
type DescribeFleetHistoryRequest struct {
	*aws.Request
	Input *types.DescribeFleetHistoryInput
	Copy  func(*types.DescribeFleetHistoryInput) DescribeFleetHistoryRequest
}

// Send marshals and sends the DescribeFleetHistory API request.
func (r DescribeFleetHistoryRequest) Send(ctx context.Context) (*DescribeFleetHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetHistoryResponse{
		DescribeFleetHistoryOutput: r.Request.Data.(*types.DescribeFleetHistoryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFleetHistoryResponse is the response type for the
// DescribeFleetHistory API operation.
type DescribeFleetHistoryResponse struct {
	*types.DescribeFleetHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleetHistory request.
func (r *DescribeFleetHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
