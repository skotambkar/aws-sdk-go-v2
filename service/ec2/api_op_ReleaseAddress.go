// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opReleaseAddress = "ReleaseAddress"

// ReleaseAddressRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Releases the specified Elastic IP address.
//
// [EC2-Classic, default VPC] Releasing an Elastic IP address automatically
// disassociates it from any instance that it's associated with. To disassociate
// an Elastic IP address without releasing it, use DisassociateAddress.
//
// [Nondefault VPC] You must use DisassociateAddress to disassociate the Elastic
// IP address before you can release it. Otherwise, Amazon EC2 returns an error
// (InvalidIPAddress.InUse).
//
// After releasing an Elastic IP address, it is released to the IP address pool.
// Be sure to update your DNS records and any servers or devices that communicate
// with the address. If you attempt to release an Elastic IP address that you
// already released, you'll get an AuthFailure error if the address is already
// allocated to another AWS account.
//
// [EC2-VPC] After you release an Elastic IP address for use in a VPC, you might
// be able to recover it. For more information, see AllocateAddress.
//
//    // Example sending a request using ReleaseAddressRequest.
//    req := client.ReleaseAddressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReleaseAddress
func (c *Client) ReleaseAddressRequest(input *types.ReleaseAddressInput) ReleaseAddressRequest {
	op := &aws.Operation{
		Name:       opReleaseAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ReleaseAddressInput{}
	}

	req := c.newRequest(op, input, &types.ReleaseAddressOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ReleaseAddressRequest{Request: req, Input: input, Copy: c.ReleaseAddressRequest}
}

// ReleaseAddressRequest is the request type for the
// ReleaseAddress API operation.
type ReleaseAddressRequest struct {
	*aws.Request
	Input *types.ReleaseAddressInput
	Copy  func(*types.ReleaseAddressInput) ReleaseAddressRequest
}

// Send marshals and sends the ReleaseAddress API request.
func (r ReleaseAddressRequest) Send(ctx context.Context) (*ReleaseAddressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReleaseAddressResponse{
		ReleaseAddressOutput: r.Request.Data.(*types.ReleaseAddressOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReleaseAddressResponse is the response type for the
// ReleaseAddress API operation.
type ReleaseAddressResponse struct {
	*types.ReleaseAddressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReleaseAddress request.
func (r *ReleaseAddressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
