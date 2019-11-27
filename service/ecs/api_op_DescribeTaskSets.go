// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opDescribeTaskSets = "DescribeTaskSets"

// DescribeTaskSetsRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Describes the task sets in the specified cluster and service. This is used
// when a service uses the EXTERNAL deployment controller type. For more information,
// see Amazon ECS Deployment Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using DescribeTaskSetsRequest.
//    req := client.DescribeTaskSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DescribeTaskSets
func (c *Client) DescribeTaskSetsRequest(input *types.DescribeTaskSetsInput) DescribeTaskSetsRequest {
	op := &aws.Operation{
		Name:       opDescribeTaskSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeTaskSetsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTaskSetsOutput{})
	return DescribeTaskSetsRequest{Request: req, Input: input, Copy: c.DescribeTaskSetsRequest}
}

// DescribeTaskSetsRequest is the request type for the
// DescribeTaskSets API operation.
type DescribeTaskSetsRequest struct {
	*aws.Request
	Input *types.DescribeTaskSetsInput
	Copy  func(*types.DescribeTaskSetsInput) DescribeTaskSetsRequest
}

// Send marshals and sends the DescribeTaskSets API request.
func (r DescribeTaskSetsRequest) Send(ctx context.Context) (*DescribeTaskSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTaskSetsResponse{
		DescribeTaskSetsOutput: r.Request.Data.(*types.DescribeTaskSetsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTaskSetsResponse is the response type for the
// DescribeTaskSets API operation.
type DescribeTaskSetsResponse struct {
	*types.DescribeTaskSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTaskSets request.
func (r *DescribeTaskSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
