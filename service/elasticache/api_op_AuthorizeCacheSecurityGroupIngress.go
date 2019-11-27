// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opAuthorizeCacheSecurityGroupIngress = "AuthorizeCacheSecurityGroupIngress"

// AuthorizeCacheSecurityGroupIngressRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Allows network ingress to a cache security group. Applications using ElastiCache
// must be running on Amazon EC2, and Amazon EC2 security groups are used as
// the authorization mechanism.
//
// You cannot authorize ingress from an Amazon EC2 security group in one region
// to an ElastiCache cluster in another region.
//
//    // Example sending a request using AuthorizeCacheSecurityGroupIngressRequest.
//    req := client.AuthorizeCacheSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/AuthorizeCacheSecurityGroupIngress
func (c *Client) AuthorizeCacheSecurityGroupIngressRequest(input *types.AuthorizeCacheSecurityGroupIngressInput) AuthorizeCacheSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opAuthorizeCacheSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AuthorizeCacheSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &types.AuthorizeCacheSecurityGroupIngressOutput{})
	return AuthorizeCacheSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.AuthorizeCacheSecurityGroupIngressRequest}
}

// AuthorizeCacheSecurityGroupIngressRequest is the request type for the
// AuthorizeCacheSecurityGroupIngress API operation.
type AuthorizeCacheSecurityGroupIngressRequest struct {
	*aws.Request
	Input *types.AuthorizeCacheSecurityGroupIngressInput
	Copy  func(*types.AuthorizeCacheSecurityGroupIngressInput) AuthorizeCacheSecurityGroupIngressRequest
}

// Send marshals and sends the AuthorizeCacheSecurityGroupIngress API request.
func (r AuthorizeCacheSecurityGroupIngressRequest) Send(ctx context.Context) (*AuthorizeCacheSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AuthorizeCacheSecurityGroupIngressResponse{
		AuthorizeCacheSecurityGroupIngressOutput: r.Request.Data.(*types.AuthorizeCacheSecurityGroupIngressOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AuthorizeCacheSecurityGroupIngressResponse is the response type for the
// AuthorizeCacheSecurityGroupIngress API operation.
type AuthorizeCacheSecurityGroupIngressResponse struct {
	*types.AuthorizeCacheSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AuthorizeCacheSecurityGroupIngress request.
func (r *AuthorizeCacheSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
