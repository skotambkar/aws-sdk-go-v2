// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateCustomerGateway = "CreateCustomerGateway"

// CreateCustomerGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Provides information to AWS about your VPN customer gateway device. The customer
// gateway is the appliance at your end of the VPN connection. (The device on
// the AWS side of the VPN connection is the virtual private gateway.) You must
// provide the Internet-routable IP address of the customer gateway's external
// interface. The IP address must be static and can be behind a device performing
// network address translation (NAT).
//
// For devices that use Border Gateway Protocol (BGP), you can also provide
// the device's BGP Autonomous System Number (ASN). You can use an existing
// ASN assigned to your network. If you don't have an ASN already, you can use
// a private ASN (in the 64512 - 65534 range).
//
// Amazon EC2 supports all 2-byte ASN numbers in the range of 1 - 65534, with
// the exception of 7224, which is reserved in the us-east-1 Region, and 9059,
// which is reserved in the eu-west-1 Region.
//
// For more information, see AWS Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html)
// in the AWS Site-to-Site VPN User Guide.
//
// You cannot create more than one customer gateway with the same VPN type,
// IP address, and BGP ASN parameter values. If you run an identical request
// more than one time, the first request creates the customer gateway, and subsequent
// requests return information about the existing customer gateway. The subsequent
// requests do not create new customer gateway resources.
//
//    // Example sending a request using CreateCustomerGatewayRequest.
//    req := client.CreateCustomerGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateCustomerGateway
func (c *Client) CreateCustomerGatewayRequest(input *types.CreateCustomerGatewayInput) CreateCustomerGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateCustomerGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateCustomerGatewayInput{}
	}

	req := c.newRequest(op, input, &types.CreateCustomerGatewayOutput{})
	return CreateCustomerGatewayRequest{Request: req, Input: input, Copy: c.CreateCustomerGatewayRequest}
}

// CreateCustomerGatewayRequest is the request type for the
// CreateCustomerGateway API operation.
type CreateCustomerGatewayRequest struct {
	*aws.Request
	Input *types.CreateCustomerGatewayInput
	Copy  func(*types.CreateCustomerGatewayInput) CreateCustomerGatewayRequest
}

// Send marshals and sends the CreateCustomerGateway API request.
func (r CreateCustomerGatewayRequest) Send(ctx context.Context) (*CreateCustomerGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCustomerGatewayResponse{
		CreateCustomerGatewayOutput: r.Request.Data.(*types.CreateCustomerGatewayOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCustomerGatewayResponse is the response type for the
// CreateCustomerGateway API operation.
type CreateCustomerGatewayResponse struct {
	*types.CreateCustomerGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCustomerGateway request.
func (r *CreateCustomerGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
