// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

const opSetLoadBalancerPoliciesForBackendServer = "SetLoadBalancerPoliciesForBackendServer"

// SetLoadBalancerPoliciesForBackendServerRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Replaces the set of policies associated with the specified port on which
// the EC2 instance is listening with a new set of policies. At this time, only
// the back-end server authentication policy type can be applied to the instance
// ports; this policy type is composed of multiple public key policies.
//
// Each time you use SetLoadBalancerPoliciesForBackendServer to enable the policies,
// use the PolicyNames parameter to list the policies that you want to enable.
//
// You can use DescribeLoadBalancers or DescribeLoadBalancerPolicies to verify
// that the policy is associated with the EC2 instance.
//
// For more information about enabling back-end instance authentication, see
// Configure Back-end Instance Authentication (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-create-https-ssl-load-balancer.html#configure_backendauth_clt)
// in the Classic Load Balancers Guide. For more information about Proxy Protocol,
// see Configure Proxy Protocol Support (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-proxy-protocol.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using SetLoadBalancerPoliciesForBackendServerRequest.
//    req := client.SetLoadBalancerPoliciesForBackendServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/SetLoadBalancerPoliciesForBackendServer
func (c *Client) SetLoadBalancerPoliciesForBackendServerRequest(input *types.SetLoadBalancerPoliciesForBackendServerInput) SetLoadBalancerPoliciesForBackendServerRequest {
	op := &aws.Operation{
		Name:       opSetLoadBalancerPoliciesForBackendServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetLoadBalancerPoliciesForBackendServerInput{}
	}

	req := c.newRequest(op, input, &types.SetLoadBalancerPoliciesForBackendServerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.SetLoadBalancerPoliciesForBackendServerMarshaler{Input: input}.GetNamedBuildHandler())

	return SetLoadBalancerPoliciesForBackendServerRequest{Request: req, Input: input, Copy: c.SetLoadBalancerPoliciesForBackendServerRequest}
}

// SetLoadBalancerPoliciesForBackendServerRequest is the request type for the
// SetLoadBalancerPoliciesForBackendServer API operation.
type SetLoadBalancerPoliciesForBackendServerRequest struct {
	*aws.Request
	Input *types.SetLoadBalancerPoliciesForBackendServerInput
	Copy  func(*types.SetLoadBalancerPoliciesForBackendServerInput) SetLoadBalancerPoliciesForBackendServerRequest
}

// Send marshals and sends the SetLoadBalancerPoliciesForBackendServer API request.
func (r SetLoadBalancerPoliciesForBackendServerRequest) Send(ctx context.Context) (*SetLoadBalancerPoliciesForBackendServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetLoadBalancerPoliciesForBackendServerResponse{
		SetLoadBalancerPoliciesForBackendServerOutput: r.Request.Data.(*types.SetLoadBalancerPoliciesForBackendServerOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetLoadBalancerPoliciesForBackendServerResponse is the response type for the
// SetLoadBalancerPoliciesForBackendServer API operation.
type SetLoadBalancerPoliciesForBackendServerResponse struct {
	*types.SetLoadBalancerPoliciesForBackendServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetLoadBalancerPoliciesForBackendServer request.
func (r *SetLoadBalancerPoliciesForBackendServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
