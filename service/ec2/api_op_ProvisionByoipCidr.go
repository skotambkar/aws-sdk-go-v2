// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opProvisionByoipCidr = "ProvisionByoipCidr"

// ProvisionByoipCidrRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Provisions an address range for use with your AWS resources through bring
// your own IP addresses (BYOIP) and creates a corresponding address pool. After
// the address range is provisioned, it is ready to be advertised using AdvertiseByoipCidr.
//
// AWS verifies that you own the address range and are authorized to advertise
// it. You must ensure that the address range is registered to you and that
// you created an RPKI ROA to authorize Amazon ASNs 16509 and 14618 to advertise
// the address range. For more information, see Bring Your Own IP Addresses
// (BYOIP) (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// Provisioning an address range is an asynchronous operation, so the call returns
// immediately, but the address range is not ready to use until its status changes
// from pending-provision to provisioned. To monitor the status of an address
// range, use DescribeByoipCidrs. To allocate an Elastic IP address from your
// address pool, use AllocateAddress with either the specific address from the
// address pool or the ID of the address pool.
//
//    // Example sending a request using ProvisionByoipCidrRequest.
//    req := client.ProvisionByoipCidrRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ProvisionByoipCidr
func (c *Client) ProvisionByoipCidrRequest(input *types.ProvisionByoipCidrInput) ProvisionByoipCidrRequest {
	op := &aws.Operation{
		Name:       opProvisionByoipCidr,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ProvisionByoipCidrInput{}
	}

	req := c.newRequest(op, input, &types.ProvisionByoipCidrOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.ProvisionByoipCidrMarshaler{Input: input}.GetNamedBuildHandler())

	return ProvisionByoipCidrRequest{Request: req, Input: input, Copy: c.ProvisionByoipCidrRequest}
}

// ProvisionByoipCidrRequest is the request type for the
// ProvisionByoipCidr API operation.
type ProvisionByoipCidrRequest struct {
	*aws.Request
	Input *types.ProvisionByoipCidrInput
	Copy  func(*types.ProvisionByoipCidrInput) ProvisionByoipCidrRequest
}

// Send marshals and sends the ProvisionByoipCidr API request.
func (r ProvisionByoipCidrRequest) Send(ctx context.Context) (*ProvisionByoipCidrResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ProvisionByoipCidrResponse{
		ProvisionByoipCidrOutput: r.Request.Data.(*types.ProvisionByoipCidrOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ProvisionByoipCidrResponse is the response type for the
// ProvisionByoipCidr API operation.
type ProvisionByoipCidrResponse struct {
	*types.ProvisionByoipCidrOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ProvisionByoipCidr request.
func (r *ProvisionByoipCidrResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
