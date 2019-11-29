// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDisableGateway = "DisableGateway"

// DisableGatewayRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Disables a tape gateway when the gateway is no longer functioning. For example,
// if your gateway VM is damaged, you can disable the gateway so you can recover
// virtual tapes.
//
// Use this operation for a tape gateway that is not reachable or not functioning.
// This operation is only supported in the tape gateway type.
//
// Once a gateway is disabled it cannot be enabled.
//
//    // Example sending a request using DisableGatewayRequest.
//    req := client.DisableGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DisableGateway
func (c *Client) DisableGatewayRequest(input *types.DisableGatewayInput) DisableGatewayRequest {
	op := &aws.Operation{
		Name:       opDisableGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisableGatewayInput{}
	}

	req := c.newRequest(op, input, &types.DisableGatewayOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DisableGatewayMarshaler{Input: input}.GetNamedBuildHandler())

	return DisableGatewayRequest{Request: req, Input: input, Copy: c.DisableGatewayRequest}
}

// DisableGatewayRequest is the request type for the
// DisableGateway API operation.
type DisableGatewayRequest struct {
	*aws.Request
	Input *types.DisableGatewayInput
	Copy  func(*types.DisableGatewayInput) DisableGatewayRequest
}

// Send marshals and sends the DisableGateway API request.
func (r DisableGatewayRequest) Send(ctx context.Context) (*DisableGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableGatewayResponse{
		DisableGatewayOutput: r.Request.Data.(*types.DisableGatewayOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableGatewayResponse is the response type for the
// DisableGateway API operation.
type DisableGatewayResponse struct {
	*types.DisableGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableGateway request.
func (r *DisableGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
