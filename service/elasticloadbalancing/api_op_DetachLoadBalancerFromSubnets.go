// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

const opDetachLoadBalancerFromSubnets = "DetachLoadBalancerFromSubnets"

// DetachLoadBalancerFromSubnetsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Removes the specified subnets from the set of configured subnets for the
// load balancer.
//
// After a subnet is removed, all EC2 instances registered with the load balancer
// in the removed subnet go into the OutOfService state. Then, the load balancer
// balances the traffic among the remaining routable subnets.
//
//    // Example sending a request using DetachLoadBalancerFromSubnetsRequest.
//    req := client.DetachLoadBalancerFromSubnetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DetachLoadBalancerFromSubnets
func (c *Client) DetachLoadBalancerFromSubnetsRequest(input *types.DetachLoadBalancerFromSubnetsInput) DetachLoadBalancerFromSubnetsRequest {
	op := &aws.Operation{
		Name:       opDetachLoadBalancerFromSubnets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DetachLoadBalancerFromSubnetsInput{}
	}

	req := c.newRequest(op, input, &types.DetachLoadBalancerFromSubnetsOutput{})
	return DetachLoadBalancerFromSubnetsRequest{Request: req, Input: input, Copy: c.DetachLoadBalancerFromSubnetsRequest}
}

// DetachLoadBalancerFromSubnetsRequest is the request type for the
// DetachLoadBalancerFromSubnets API operation.
type DetachLoadBalancerFromSubnetsRequest struct {
	*aws.Request
	Input *types.DetachLoadBalancerFromSubnetsInput
	Copy  func(*types.DetachLoadBalancerFromSubnetsInput) DetachLoadBalancerFromSubnetsRequest
}

// Send marshals and sends the DetachLoadBalancerFromSubnets API request.
func (r DetachLoadBalancerFromSubnetsRequest) Send(ctx context.Context) (*DetachLoadBalancerFromSubnetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachLoadBalancerFromSubnetsResponse{
		DetachLoadBalancerFromSubnetsOutput: r.Request.Data.(*types.DetachLoadBalancerFromSubnetsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachLoadBalancerFromSubnetsResponse is the response type for the
// DetachLoadBalancerFromSubnets API operation.
type DetachLoadBalancerFromSubnetsResponse struct {
	*types.DetachLoadBalancerFromSubnetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachLoadBalancerFromSubnets request.
func (r *DetachLoadBalancerFromSubnetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
