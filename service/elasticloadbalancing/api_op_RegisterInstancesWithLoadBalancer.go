// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

const opRegisterInstancesWithLoadBalancer = "RegisterInstancesWithLoadBalancer"

// RegisterInstancesWithLoadBalancerRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Adds the specified instances to the specified load balancer.
//
// The instance must be a running instance in the same network as the load balancer
// (EC2-Classic or the same VPC). If you have EC2-Classic instances and a load
// balancer in a VPC with ClassicLink enabled, you can link the EC2-Classic
// instances to that VPC and then register the linked EC2-Classic instances
// with the load balancer in the VPC.
//
// Note that RegisterInstanceWithLoadBalancer completes when the request has
// been registered. Instance registration takes a little time to complete. To
// check the state of the registered instances, use DescribeLoadBalancers or
// DescribeInstanceHealth.
//
// After the instance is registered, it starts receiving traffic and requests
// from the load balancer. Any instance that is not in one of the Availability
// Zones registered for the load balancer is moved to the OutOfService state.
// If an Availability Zone is added to the load balancer later, any instances
// registered with the load balancer move to the InService state.
//
// To deregister instances from a load balancer, use DeregisterInstancesFromLoadBalancer.
//
// For more information, see Register or De-Register EC2 Instances (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-deregister-register-instances.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using RegisterInstancesWithLoadBalancerRequest.
//    req := client.RegisterInstancesWithLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/RegisterInstancesWithLoadBalancer
func (c *Client) RegisterInstancesWithLoadBalancerRequest(input *types.RegisterInstancesWithLoadBalancerInput) RegisterInstancesWithLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opRegisterInstancesWithLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterInstancesWithLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &types.RegisterInstancesWithLoadBalancerOutput{})
	return RegisterInstancesWithLoadBalancerRequest{Request: req, Input: input, Copy: c.RegisterInstancesWithLoadBalancerRequest}
}

// RegisterInstancesWithLoadBalancerRequest is the request type for the
// RegisterInstancesWithLoadBalancer API operation.
type RegisterInstancesWithLoadBalancerRequest struct {
	*aws.Request
	Input *types.RegisterInstancesWithLoadBalancerInput
	Copy  func(*types.RegisterInstancesWithLoadBalancerInput) RegisterInstancesWithLoadBalancerRequest
}

// Send marshals and sends the RegisterInstancesWithLoadBalancer API request.
func (r RegisterInstancesWithLoadBalancerRequest) Send(ctx context.Context) (*RegisterInstancesWithLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterInstancesWithLoadBalancerResponse{
		RegisterInstancesWithLoadBalancerOutput: r.Request.Data.(*types.RegisterInstancesWithLoadBalancerOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterInstancesWithLoadBalancerResponse is the response type for the
// RegisterInstancesWithLoadBalancer API operation.
type RegisterInstancesWithLoadBalancerResponse struct {
	*types.RegisterInstancesWithLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterInstancesWithLoadBalancer request.
func (r *RegisterInstancesWithLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
