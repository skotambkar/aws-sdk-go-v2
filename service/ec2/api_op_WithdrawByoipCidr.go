// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opWithdrawByoipCidr = "WithdrawByoipCidr"

// WithdrawByoipCidrRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Stops advertising an IPv4 address range that is provisioned as an address
// pool.
//
// You can perform this operation at most once every 10 seconds, even if you
// specify different address ranges each time.
//
// It can take a few minutes before traffic to the specified addresses stops
// routing to AWS because of BGP propagation delays.
//
//    // Example sending a request using WithdrawByoipCidrRequest.
//    req := client.WithdrawByoipCidrRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/WithdrawByoipCidr
func (c *Client) WithdrawByoipCidrRequest(input *types.WithdrawByoipCidrInput) WithdrawByoipCidrRequest {
	op := &aws.Operation{
		Name:       opWithdrawByoipCidr,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.WithdrawByoipCidrInput{}
	}

	req := c.newRequest(op, input, &types.WithdrawByoipCidrOutput{})
	return WithdrawByoipCidrRequest{Request: req, Input: input, Copy: c.WithdrawByoipCidrRequest}
}

// WithdrawByoipCidrRequest is the request type for the
// WithdrawByoipCidr API operation.
type WithdrawByoipCidrRequest struct {
	*aws.Request
	Input *types.WithdrawByoipCidrInput
	Copy  func(*types.WithdrawByoipCidrInput) WithdrawByoipCidrRequest
}

// Send marshals and sends the WithdrawByoipCidr API request.
func (r WithdrawByoipCidrRequest) Send(ctx context.Context) (*WithdrawByoipCidrResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &WithdrawByoipCidrResponse{
		WithdrawByoipCidrOutput: r.Request.Data.(*types.WithdrawByoipCidrOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// WithdrawByoipCidrResponse is the response type for the
// WithdrawByoipCidr API operation.
type WithdrawByoipCidrResponse struct {
	*types.WithdrawByoipCidrOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// WithdrawByoipCidr request.
func (r *WithdrawByoipCidrResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
