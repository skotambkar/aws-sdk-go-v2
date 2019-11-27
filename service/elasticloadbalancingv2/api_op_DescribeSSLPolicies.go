// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

const opDescribeSSLPolicies = "DescribeSSLPolicies"

// DescribeSSLPoliciesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the specified policies or all policies used for SSL negotiation.
//
// For more information, see Security Policies (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
// in the Application Load Balancers Guide.
//
//    // Example sending a request using DescribeSSLPoliciesRequest.
//    req := client.DescribeSSLPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeSSLPolicies
func (c *Client) DescribeSSLPoliciesRequest(input *types.DescribeSSLPoliciesInput) DescribeSSLPoliciesRequest {
	op := &aws.Operation{
		Name:       opDescribeSSLPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeSSLPoliciesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSSLPoliciesOutput{})
	return DescribeSSLPoliciesRequest{Request: req, Input: input, Copy: c.DescribeSSLPoliciesRequest}
}

// DescribeSSLPoliciesRequest is the request type for the
// DescribeSSLPolicies API operation.
type DescribeSSLPoliciesRequest struct {
	*aws.Request
	Input *types.DescribeSSLPoliciesInput
	Copy  func(*types.DescribeSSLPoliciesInput) DescribeSSLPoliciesRequest
}

// Send marshals and sends the DescribeSSLPolicies API request.
func (r DescribeSSLPoliciesRequest) Send(ctx context.Context) (*DescribeSSLPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSSLPoliciesResponse{
		DescribeSSLPoliciesOutput: r.Request.Data.(*types.DescribeSSLPoliciesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSSLPoliciesResponse is the response type for the
// DescribeSSLPolicies API operation.
type DescribeSSLPoliciesResponse struct {
	*types.DescribeSSLPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSSLPolicies request.
func (r *DescribeSSLPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
