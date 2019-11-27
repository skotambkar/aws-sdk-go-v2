// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

const opCreateTargetGroup = "CreateTargetGroup"

// CreateTargetGroupRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates a target group.
//
// To register targets with the target group, use RegisterTargets. To update
// the health check settings for the target group, use ModifyTargetGroup. To
// monitor the health of targets in the target group, use DescribeTargetHealth.
//
// To route traffic to the targets in a target group, specify the target group
// in an action using CreateListener or CreateRule.
//
// To delete a target group, use DeleteTargetGroup.
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple target groups with the same settings, each
// call succeeds.
//
// For more information, see Target Groups for Your Application Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html)
// in the Application Load Balancers Guide or Target Groups for Your Network
// Load Balancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html)
// in the Network Load Balancers Guide.
//
//    // Example sending a request using CreateTargetGroupRequest.
//    req := client.CreateTargetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/CreateTargetGroup
func (c *Client) CreateTargetGroupRequest(input *types.CreateTargetGroupInput) CreateTargetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateTargetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateTargetGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateTargetGroupOutput{})
	return CreateTargetGroupRequest{Request: req, Input: input, Copy: c.CreateTargetGroupRequest}
}

// CreateTargetGroupRequest is the request type for the
// CreateTargetGroup API operation.
type CreateTargetGroupRequest struct {
	*aws.Request
	Input *types.CreateTargetGroupInput
	Copy  func(*types.CreateTargetGroupInput) CreateTargetGroupRequest
}

// Send marshals and sends the CreateTargetGroup API request.
func (r CreateTargetGroupRequest) Send(ctx context.Context) (*CreateTargetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTargetGroupResponse{
		CreateTargetGroupOutput: r.Request.Data.(*types.CreateTargetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTargetGroupResponse is the response type for the
// CreateTargetGroup API operation.
type CreateTargetGroupResponse struct {
	*types.CreateTargetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTargetGroup request.
func (r *CreateTargetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
