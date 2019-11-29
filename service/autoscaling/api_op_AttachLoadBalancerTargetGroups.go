// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

const opAttachLoadBalancerTargetGroups = "AttachLoadBalancerTargetGroups"

// AttachLoadBalancerTargetGroupsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Attaches one or more target groups to the specified Auto Scaling group.
//
// To describe the target groups for an Auto Scaling group, use DescribeLoadBalancerTargetGroups.
// To detach the target group from the Auto Scaling group, use DetachLoadBalancerTargetGroups.
//
// With Application Load Balancers and Network Load Balancers, instances are
// registered as targets with a target group. With Classic Load Balancers, instances
// are registered with the load balancer. For more information, see Attaching
// a Load Balancer to Your Auto Scaling Group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using AttachLoadBalancerTargetGroupsRequest.
//    req := client.AttachLoadBalancerTargetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AttachLoadBalancerTargetGroups
func (c *Client) AttachLoadBalancerTargetGroupsRequest(input *types.AttachLoadBalancerTargetGroupsInput) AttachLoadBalancerTargetGroupsRequest {
	op := &aws.Operation{
		Name:       opAttachLoadBalancerTargetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AttachLoadBalancerTargetGroupsInput{}
	}

	req := c.newRequest(op, input, &types.AttachLoadBalancerTargetGroupsOutput{})
	return AttachLoadBalancerTargetGroupsRequest{Request: req, Input: input, Copy: c.AttachLoadBalancerTargetGroupsRequest}
}

// AttachLoadBalancerTargetGroupsRequest is the request type for the
// AttachLoadBalancerTargetGroups API operation.
type AttachLoadBalancerTargetGroupsRequest struct {
	*aws.Request
	Input *types.AttachLoadBalancerTargetGroupsInput
	Copy  func(*types.AttachLoadBalancerTargetGroupsInput) AttachLoadBalancerTargetGroupsRequest
}

// Send marshals and sends the AttachLoadBalancerTargetGroups API request.
func (r AttachLoadBalancerTargetGroupsRequest) Send(ctx context.Context) (*AttachLoadBalancerTargetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachLoadBalancerTargetGroupsResponse{
		AttachLoadBalancerTargetGroupsOutput: r.Request.Data.(*types.AttachLoadBalancerTargetGroupsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachLoadBalancerTargetGroupsResponse is the response type for the
// AttachLoadBalancerTargetGroups API operation.
type AttachLoadBalancerTargetGroupsResponse struct {
	*types.AttachLoadBalancerTargetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachLoadBalancerTargetGroups request.
func (r *AttachLoadBalancerTargetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
