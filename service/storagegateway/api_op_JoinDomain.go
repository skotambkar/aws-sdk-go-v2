// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opJoinDomain = "JoinDomain"

// JoinDomainRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Adds a file gateway to an Active Directory domain. This operation is only
// supported for file gateways that support the SMB file protocol.
//
//    // Example sending a request using JoinDomainRequest.
//    req := client.JoinDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/JoinDomain
func (c *Client) JoinDomainRequest(input *types.JoinDomainInput) JoinDomainRequest {
	op := &aws.Operation{
		Name:       opJoinDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.JoinDomainInput{}
	}

	req := c.newRequest(op, input, &types.JoinDomainOutput{})
	return JoinDomainRequest{Request: req, Input: input, Copy: c.JoinDomainRequest}
}

// JoinDomainRequest is the request type for the
// JoinDomain API operation.
type JoinDomainRequest struct {
	*aws.Request
	Input *types.JoinDomainInput
	Copy  func(*types.JoinDomainInput) JoinDomainRequest
}

// Send marshals and sends the JoinDomain API request.
func (r JoinDomainRequest) Send(ctx context.Context) (*JoinDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &JoinDomainResponse{
		JoinDomainOutput: r.Request.Data.(*types.JoinDomainOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// JoinDomainResponse is the response type for the
// JoinDomain API operation.
type JoinDomainResponse struct {
	*types.JoinDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// JoinDomain request.
func (r *JoinDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
