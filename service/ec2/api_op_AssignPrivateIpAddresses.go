// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opAssignPrivateIpAddresses = "AssignPrivateIpAddresses"

// AssignPrivateIpAddressesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Assigns one or more secondary private IP addresses to the specified network
// interface.
//
// You can specify one or more specific secondary IP addresses, or you can specify
// the number of secondary IP addresses to be automatically assigned within
// the subnet's CIDR block range. The number of secondary IP addresses that
// you can assign to an instance varies by instance type. For information about
// instance types, see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
// in the Amazon Elastic Compute Cloud User Guide. For more information about
// Elastic IP addresses, see Elastic IP Addresses (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// When you move a secondary private IP address to another network interface,
// any Elastic IP address that is associated with the IP address is also moved.
//
// Remapping an IP address is an asynchronous operation. When you move an IP
// address from one network interface to another, check network/interfaces/macs/mac/local-ipv4s
// in the instance metadata to confirm that the remapping is complete.
//
//    // Example sending a request using AssignPrivateIpAddressesRequest.
//    req := client.AssignPrivateIpAddressesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignPrivateIpAddresses
func (c *Client) AssignPrivateIpAddressesRequest(input *types.AssignPrivateIpAddressesInput) AssignPrivateIpAddressesRequest {
	op := &aws.Operation{
		Name:       opAssignPrivateIpAddresses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssignPrivateIpAddressesInput{}
	}

	req := c.newRequest(op, input, &types.AssignPrivateIpAddressesOutput{})
	return AssignPrivateIpAddressesRequest{Request: req, Input: input, Copy: c.AssignPrivateIpAddressesRequest}
}

// AssignPrivateIpAddressesRequest is the request type for the
// AssignPrivateIpAddresses API operation.
type AssignPrivateIpAddressesRequest struct {
	*aws.Request
	Input *types.AssignPrivateIpAddressesInput
	Copy  func(*types.AssignPrivateIpAddressesInput) AssignPrivateIpAddressesRequest
}

// Send marshals and sends the AssignPrivateIpAddresses API request.
func (r AssignPrivateIpAddressesRequest) Send(ctx context.Context) (*AssignPrivateIpAddressesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssignPrivateIpAddressesResponse{
		AssignPrivateIpAddressesOutput: r.Request.Data.(*types.AssignPrivateIpAddressesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssignPrivateIpAddressesResponse is the response type for the
// AssignPrivateIpAddresses API operation.
type AssignPrivateIpAddressesResponse struct {
	*types.AssignPrivateIpAddressesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssignPrivateIpAddresses request.
func (r *AssignPrivateIpAddressesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
