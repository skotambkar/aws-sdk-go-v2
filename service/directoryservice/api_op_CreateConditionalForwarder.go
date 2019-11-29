// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
)

const opCreateConditionalForwarder = "CreateConditionalForwarder"

// CreateConditionalForwarderRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Creates a conditional forwarder associated with your AWS directory. Conditional
// forwarders are required in order to set up a trust relationship with another
// domain. The conditional forwarder points to the trusted domain.
//
//    // Example sending a request using CreateConditionalForwarderRequest.
//    req := client.CreateConditionalForwarderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateConditionalForwarder
func (c *Client) CreateConditionalForwarderRequest(input *types.CreateConditionalForwarderInput) CreateConditionalForwarderRequest {
	op := &aws.Operation{
		Name:       opCreateConditionalForwarder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateConditionalForwarderInput{}
	}

	req := c.newRequest(op, input, &types.CreateConditionalForwarderOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateConditionalForwarderMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateConditionalForwarderRequest{Request: req, Input: input, Copy: c.CreateConditionalForwarderRequest}
}

// CreateConditionalForwarderRequest is the request type for the
// CreateConditionalForwarder API operation.
type CreateConditionalForwarderRequest struct {
	*aws.Request
	Input *types.CreateConditionalForwarderInput
	Copy  func(*types.CreateConditionalForwarderInput) CreateConditionalForwarderRequest
}

// Send marshals and sends the CreateConditionalForwarder API request.
func (r CreateConditionalForwarderRequest) Send(ctx context.Context) (*CreateConditionalForwarderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConditionalForwarderResponse{
		CreateConditionalForwarderOutput: r.Request.Data.(*types.CreateConditionalForwarderOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConditionalForwarderResponse is the response type for the
// CreateConditionalForwarder API operation.
type CreateConditionalForwarderResponse struct {
	*types.CreateConditionalForwarderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConditionalForwarder request.
func (r *CreateConditionalForwarderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
