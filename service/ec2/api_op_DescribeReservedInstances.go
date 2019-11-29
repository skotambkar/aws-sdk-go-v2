// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeReservedInstances = "DescribeReservedInstances"

// DescribeReservedInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of the Reserved Instances that you purchased.
//
// For more information about Reserved Instances, see Reserved Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts-on-demand-reserved-instances.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeReservedInstancesRequest.
//    req := client.DescribeReservedInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeReservedInstances
func (c *Client) DescribeReservedInstancesRequest(input *types.DescribeReservedInstancesInput) DescribeReservedInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeReservedInstancesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeReservedInstancesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DescribeReservedInstancesMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeReservedInstancesRequest{Request: req, Input: input, Copy: c.DescribeReservedInstancesRequest}
}

// DescribeReservedInstancesRequest is the request type for the
// DescribeReservedInstances API operation.
type DescribeReservedInstancesRequest struct {
	*aws.Request
	Input *types.DescribeReservedInstancesInput
	Copy  func(*types.DescribeReservedInstancesInput) DescribeReservedInstancesRequest
}

// Send marshals and sends the DescribeReservedInstances API request.
func (r DescribeReservedInstancesRequest) Send(ctx context.Context) (*DescribeReservedInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedInstancesResponse{
		DescribeReservedInstancesOutput: r.Request.Data.(*types.DescribeReservedInstancesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeReservedInstancesResponse is the response type for the
// DescribeReservedInstances API operation.
type DescribeReservedInstancesResponse struct {
	*types.DescribeReservedInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedInstances request.
func (r *DescribeReservedInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
