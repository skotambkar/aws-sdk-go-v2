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

const opDeleteVpnGateway = "DeleteVpnGateway"

// DeleteVpnGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified virtual private gateway. We recommend that before you
// delete a virtual private gateway, you detach it from the VPC and delete the
// VPN connection. Note that you don't need to delete the virtual private gateway
// if you plan to delete and recreate the VPN connection between your VPC and
// your network.
//
//    // Example sending a request using DeleteVpnGatewayRequest.
//    req := client.DeleteVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpnGateway
func (c *Client) DeleteVpnGatewayRequest(input *types.DeleteVpnGatewayInput) DeleteVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVpnGatewayOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DeleteVpnGatewayMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVpnGatewayRequest{Request: req, Input: input, Copy: c.DeleteVpnGatewayRequest}
}

// DeleteVpnGatewayRequest is the request type for the
// DeleteVpnGateway API operation.
type DeleteVpnGatewayRequest struct {
	*aws.Request
	Input *types.DeleteVpnGatewayInput
	Copy  func(*types.DeleteVpnGatewayInput) DeleteVpnGatewayRequest
}

// Send marshals and sends the DeleteVpnGateway API request.
func (r DeleteVpnGatewayRequest) Send(ctx context.Context) (*DeleteVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpnGatewayResponse{
		DeleteVpnGatewayOutput: r.Request.Data.(*types.DeleteVpnGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpnGatewayResponse is the response type for the
// DeleteVpnGateway API operation.
type DeleteVpnGatewayResponse struct {
	*types.DeleteVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpnGateway request.
func (r *DeleteVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
