// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opSendCommand = "SendCommand"

// SendCommandRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Runs commands on one or more managed instances.
//
//    // Example sending a request using SendCommandRequest.
//    req := client.SendCommandRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/SendCommand
func (c *Client) SendCommandRequest(input *types.SendCommandInput) SendCommandRequest {
	op := &aws.Operation{
		Name:       opSendCommand,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SendCommandInput{}
	}

	req := c.newRequest(op, input, &types.SendCommandOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.SendCommandMarshaler{Input: input}.GetNamedBuildHandler())

	return SendCommandRequest{Request: req, Input: input, Copy: c.SendCommandRequest}
}

// SendCommandRequest is the request type for the
// SendCommand API operation.
type SendCommandRequest struct {
	*aws.Request
	Input *types.SendCommandInput
	Copy  func(*types.SendCommandInput) SendCommandRequest
}

// Send marshals and sends the SendCommand API request.
func (r SendCommandRequest) Send(ctx context.Context) (*SendCommandResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendCommandResponse{
		SendCommandOutput: r.Request.Data.(*types.SendCommandOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendCommandResponse is the response type for the
// SendCommand API operation.
type SendCommandResponse struct {
	*types.SendCommandOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendCommand request.
func (r *SendCommandResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
