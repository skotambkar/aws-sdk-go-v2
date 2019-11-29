// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDetachVpnGateway = "DetachVpnGateway"

// DetachVpnGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Detaches a virtual private gateway from a VPC. You do this if you're planning
// to turn off the VPC and not use it anymore. You can confirm a virtual private
// gateway has been completely detached from a VPC by describing the virtual
// private gateway (any attachments to the virtual private gateway are also
// described).
//
// You must wait for the attachment's state to switch to detached before you
// can delete the VPC or attach a different VPC to the virtual private gateway.
//
//    // Example sending a request using DetachVpnGatewayRequest.
//    req := client.DetachVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachVpnGateway
func (c *Client) DetachVpnGatewayRequest(input *types.DetachVpnGatewayInput) DetachVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opDetachVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DetachVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &types.DetachVpnGatewayOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DetachVpnGatewayMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DetachVpnGatewayRequest{Request: req, Input: input, Copy: c.DetachVpnGatewayRequest}
}

// DetachVpnGatewayRequest is the request type for the
// DetachVpnGateway API operation.
type DetachVpnGatewayRequest struct {
	*aws.Request
	Input *types.DetachVpnGatewayInput
	Copy  func(*types.DetachVpnGatewayInput) DetachVpnGatewayRequest
}

// Send marshals and sends the DetachVpnGateway API request.
func (r DetachVpnGatewayRequest) Send(ctx context.Context) (*DetachVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachVpnGatewayResponse{
		DetachVpnGatewayOutput: r.Request.Data.(*types.DetachVpnGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachVpnGatewayResponse is the response type for the
// DetachVpnGateway API operation.
type DetachVpnGatewayResponse struct {
	*types.DetachVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachVpnGateway request.
func (r *DetachVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
