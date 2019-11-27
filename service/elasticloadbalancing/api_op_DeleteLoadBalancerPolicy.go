// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

const opDeleteLoadBalancerPolicy = "DeleteLoadBalancerPolicy"

// DeleteLoadBalancerPolicyRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Deletes the specified policy from the specified load balancer. This policy
// must not be enabled for any listeners.
//
//    // Example sending a request using DeleteLoadBalancerPolicyRequest.
//    req := client.DeleteLoadBalancerPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DeleteLoadBalancerPolicy
func (c *Client) DeleteLoadBalancerPolicyRequest(input *types.DeleteLoadBalancerPolicyInput) DeleteLoadBalancerPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteLoadBalancerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteLoadBalancerPolicyInput{}
	}

	req := c.newRequest(op, input, &types.DeleteLoadBalancerPolicyOutput{})
	return DeleteLoadBalancerPolicyRequest{Request: req, Input: input, Copy: c.DeleteLoadBalancerPolicyRequest}
}

// DeleteLoadBalancerPolicyRequest is the request type for the
// DeleteLoadBalancerPolicy API operation.
type DeleteLoadBalancerPolicyRequest struct {
	*aws.Request
	Input *types.DeleteLoadBalancerPolicyInput
	Copy  func(*types.DeleteLoadBalancerPolicyInput) DeleteLoadBalancerPolicyRequest
}

// Send marshals and sends the DeleteLoadBalancerPolicy API request.
func (r DeleteLoadBalancerPolicyRequest) Send(ctx context.Context) (*DeleteLoadBalancerPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLoadBalancerPolicyResponse{
		DeleteLoadBalancerPolicyOutput: r.Request.Data.(*types.DeleteLoadBalancerPolicyOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLoadBalancerPolicyResponse is the response type for the
// DeleteLoadBalancerPolicy API operation.
type DeleteLoadBalancerPolicyResponse struct {
	*types.DeleteLoadBalancerPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLoadBalancerPolicy request.
func (r *DeleteLoadBalancerPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
