// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workmail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opRegisterToWorkMail = "RegisterToWorkMail"

// RegisterToWorkMailRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Registers an existing and disabled user, group, or resource for Amazon WorkMail
// use by associating a mailbox and calendaring capabilities. It performs no
// change if the user, group, or resource is enabled and fails if the user,
// group, or resource is deleted. This operation results in the accumulation
// of costs. For more information, see Pricing (https://aws.amazon.com//workmail/pricing).
// The equivalent console functionality for this operation is Enable.
//
// Users can either be created by calling the CreateUser API operation or they
// can be synchronized from your directory. For more information, see DeregisterFromWorkMail.
//
//    // Example sending a request using RegisterToWorkMailRequest.
//    req := client.RegisterToWorkMailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/RegisterToWorkMail
func (c *Client) RegisterToWorkMailRequest(input *types.RegisterToWorkMailInput) RegisterToWorkMailRequest {
	op := &aws.Operation{
		Name:       opRegisterToWorkMail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterToWorkMailInput{}
	}

	req := c.newRequest(op, input, &types.RegisterToWorkMailOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RegisterToWorkMailMarshaler{Input: input}.GetNamedBuildHandler())

	return RegisterToWorkMailRequest{Request: req, Input: input, Copy: c.RegisterToWorkMailRequest}
}

// RegisterToWorkMailRequest is the request type for the
// RegisterToWorkMail API operation.
type RegisterToWorkMailRequest struct {
	*aws.Request
	Input *types.RegisterToWorkMailInput
	Copy  func(*types.RegisterToWorkMailInput) RegisterToWorkMailRequest
}

// Send marshals and sends the RegisterToWorkMail API request.
func (r RegisterToWorkMailRequest) Send(ctx context.Context) (*RegisterToWorkMailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterToWorkMailResponse{
		RegisterToWorkMailOutput: r.Request.Data.(*types.RegisterToWorkMailOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterToWorkMailResponse is the response type for the
// RegisterToWorkMail API operation.
type RegisterToWorkMailResponse struct {
	*types.RegisterToWorkMailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterToWorkMail request.
func (r *RegisterToWorkMailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
