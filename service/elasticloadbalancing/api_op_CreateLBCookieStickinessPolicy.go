// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

const opCreateLBCookieStickinessPolicy = "CreateLBCookieStickinessPolicy"

// CreateLBCookieStickinessPolicyRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Generates a stickiness policy with sticky session lifetimes controlled by
// the lifetime of the browser (user-agent) or a specified expiration period.
// This policy can be associated only with HTTP/HTTPS listeners.
//
// When a load balancer implements this policy, the load balancer uses a special
// cookie to track the instance for each request. When the load balancer receives
// a request, it first checks to see if this cookie is present in the request.
// If so, the load balancer sends the request to the application server specified
// in the cookie. If not, the load balancer sends the request to a server that
// is chosen based on the existing load-balancing algorithm.
//
// A cookie is inserted into the response for binding subsequent requests from
// the same user to that server. The validity of the cookie is based on the
// cookie expiration time, which is specified in the policy configuration.
//
// For more information, see Duration-Based Session Stickiness (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-duration)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using CreateLBCookieStickinessPolicyRequest.
//    req := client.CreateLBCookieStickinessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/CreateLBCookieStickinessPolicy
func (c *Client) CreateLBCookieStickinessPolicyRequest(input *types.CreateLBCookieStickinessPolicyInput) CreateLBCookieStickinessPolicyRequest {
	op := &aws.Operation{
		Name:       opCreateLBCookieStickinessPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateLBCookieStickinessPolicyInput{}
	}

	req := c.newRequest(op, input, &types.CreateLBCookieStickinessPolicyOutput{})
	return CreateLBCookieStickinessPolicyRequest{Request: req, Input: input, Copy: c.CreateLBCookieStickinessPolicyRequest}
}

// CreateLBCookieStickinessPolicyRequest is the request type for the
// CreateLBCookieStickinessPolicy API operation.
type CreateLBCookieStickinessPolicyRequest struct {
	*aws.Request
	Input *types.CreateLBCookieStickinessPolicyInput
	Copy  func(*types.CreateLBCookieStickinessPolicyInput) CreateLBCookieStickinessPolicyRequest
}

// Send marshals and sends the CreateLBCookieStickinessPolicy API request.
func (r CreateLBCookieStickinessPolicyRequest) Send(ctx context.Context) (*CreateLBCookieStickinessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLBCookieStickinessPolicyResponse{
		CreateLBCookieStickinessPolicyOutput: r.Request.Data.(*types.CreateLBCookieStickinessPolicyOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLBCookieStickinessPolicyResponse is the response type for the
// CreateLBCookieStickinessPolicy API operation.
type CreateLBCookieStickinessPolicyResponse struct {
	*types.CreateLBCookieStickinessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLBCookieStickinessPolicy request.
func (r *CreateLBCookieStickinessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
