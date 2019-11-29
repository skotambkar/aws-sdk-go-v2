// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opRevokeCacheSecurityGroupIngress = "RevokeCacheSecurityGroupIngress"

// RevokeCacheSecurityGroupIngressRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Revokes ingress from a cache security group. Use this operation to disallow
// access from an Amazon EC2 security group that had been previously authorized.
//
//    // Example sending a request using RevokeCacheSecurityGroupIngressRequest.
//    req := client.RevokeCacheSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/RevokeCacheSecurityGroupIngress
func (c *Client) RevokeCacheSecurityGroupIngressRequest(input *types.RevokeCacheSecurityGroupIngressInput) RevokeCacheSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opRevokeCacheSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RevokeCacheSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &types.RevokeCacheSecurityGroupIngressOutput{})
	return RevokeCacheSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.RevokeCacheSecurityGroupIngressRequest}
}

// RevokeCacheSecurityGroupIngressRequest is the request type for the
// RevokeCacheSecurityGroupIngress API operation.
type RevokeCacheSecurityGroupIngressRequest struct {
	*aws.Request
	Input *types.RevokeCacheSecurityGroupIngressInput
	Copy  func(*types.RevokeCacheSecurityGroupIngressInput) RevokeCacheSecurityGroupIngressRequest
}

// Send marshals and sends the RevokeCacheSecurityGroupIngress API request.
func (r RevokeCacheSecurityGroupIngressRequest) Send(ctx context.Context) (*RevokeCacheSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeCacheSecurityGroupIngressResponse{
		RevokeCacheSecurityGroupIngressOutput: r.Request.Data.(*types.RevokeCacheSecurityGroupIngressOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeCacheSecurityGroupIngressResponse is the response type for the
// RevokeCacheSecurityGroupIngress API operation.
type RevokeCacheSecurityGroupIngressResponse struct {
	*types.RevokeCacheSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeCacheSecurityGroupIngress request.
func (r *RevokeCacheSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
