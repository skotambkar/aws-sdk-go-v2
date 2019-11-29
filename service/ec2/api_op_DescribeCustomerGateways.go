// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeCustomerGateways = "DescribeCustomerGateways"

// DescribeCustomerGatewaysRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your VPN customer gateways.
//
// For more information, see AWS Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html)
// in the AWS Site-to-Site VPN User Guide.
//
//    // Example sending a request using DescribeCustomerGatewaysRequest.
//    req := client.DescribeCustomerGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeCustomerGateways
func (c *Client) DescribeCustomerGatewaysRequest(input *types.DescribeCustomerGatewaysInput) DescribeCustomerGatewaysRequest {
	op := &aws.Operation{
		Name:       opDescribeCustomerGateways,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeCustomerGatewaysInput{}
	}

	req := c.newRequest(op, input, &types.DescribeCustomerGatewaysOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DescribeCustomerGatewaysMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeCustomerGatewaysRequest{Request: req, Input: input, Copy: c.DescribeCustomerGatewaysRequest}
}

// DescribeCustomerGatewaysRequest is the request type for the
// DescribeCustomerGateways API operation.
type DescribeCustomerGatewaysRequest struct {
	*aws.Request
	Input *types.DescribeCustomerGatewaysInput
	Copy  func(*types.DescribeCustomerGatewaysInput) DescribeCustomerGatewaysRequest
}

// Send marshals and sends the DescribeCustomerGateways API request.
func (r DescribeCustomerGatewaysRequest) Send(ctx context.Context) (*DescribeCustomerGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCustomerGatewaysResponse{
		DescribeCustomerGatewaysOutput: r.Request.Data.(*types.DescribeCustomerGatewaysOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCustomerGatewaysResponse is the response type for the
// DescribeCustomerGateways API operation.
type DescribeCustomerGatewaysResponse struct {
	*types.DescribeCustomerGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCustomerGateways request.
func (r *DescribeCustomerGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
