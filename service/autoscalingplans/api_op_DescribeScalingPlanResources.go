// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscalingplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
)

const opDescribeScalingPlanResources = "DescribeScalingPlanResources"

// DescribeScalingPlanResourcesRequest returns a request value for making API operation for
// AWS Auto Scaling Plans.
//
// Describes the scalable resources in the specified scaling plan.
//
//    // Example sending a request using DescribeScalingPlanResourcesRequest.
//    req := client.DescribeScalingPlanResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/DescribeScalingPlanResources
func (c *Client) DescribeScalingPlanResourcesRequest(input *types.DescribeScalingPlanResourcesInput) DescribeScalingPlanResourcesRequest {
	op := &aws.Operation{
		Name:       opDescribeScalingPlanResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeScalingPlanResourcesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeScalingPlanResourcesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeScalingPlanResourcesMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeScalingPlanResourcesRequest{Request: req, Input: input, Copy: c.DescribeScalingPlanResourcesRequest}
}

// DescribeScalingPlanResourcesRequest is the request type for the
// DescribeScalingPlanResources API operation.
type DescribeScalingPlanResourcesRequest struct {
	*aws.Request
	Input *types.DescribeScalingPlanResourcesInput
	Copy  func(*types.DescribeScalingPlanResourcesInput) DescribeScalingPlanResourcesRequest
}

// Send marshals and sends the DescribeScalingPlanResources API request.
func (r DescribeScalingPlanResourcesRequest) Send(ctx context.Context) (*DescribeScalingPlanResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScalingPlanResourcesResponse{
		DescribeScalingPlanResourcesOutput: r.Request.Data.(*types.DescribeScalingPlanResourcesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeScalingPlanResourcesResponse is the response type for the
// DescribeScalingPlanResources API operation.
type DescribeScalingPlanResourcesResponse struct {
	*types.DescribeScalingPlanResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScalingPlanResources request.
func (r *DescribeScalingPlanResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
