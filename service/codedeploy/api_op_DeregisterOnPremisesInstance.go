// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opDeregisterOnPremisesInstance = "DeregisterOnPremisesInstance"

// DeregisterOnPremisesInstanceRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Deregisters an on-premises instance.
//
//    // Example sending a request using DeregisterOnPremisesInstanceRequest.
//    req := client.DeregisterOnPremisesInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeregisterOnPremisesInstance
func (c *Client) DeregisterOnPremisesInstanceRequest(input *types.DeregisterOnPremisesInstanceInput) DeregisterOnPremisesInstanceRequest {
	op := &aws.Operation{
		Name:       opDeregisterOnPremisesInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeregisterOnPremisesInstanceInput{}
	}

	req := c.newRequest(op, input, &types.DeregisterOnPremisesInstanceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeregisterOnPremisesInstanceMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeregisterOnPremisesInstanceRequest{Request: req, Input: input, Copy: c.DeregisterOnPremisesInstanceRequest}
}

// DeregisterOnPremisesInstanceRequest is the request type for the
// DeregisterOnPremisesInstance API operation.
type DeregisterOnPremisesInstanceRequest struct {
	*aws.Request
	Input *types.DeregisterOnPremisesInstanceInput
	Copy  func(*types.DeregisterOnPremisesInstanceInput) DeregisterOnPremisesInstanceRequest
}

// Send marshals and sends the DeregisterOnPremisesInstance API request.
func (r DeregisterOnPremisesInstanceRequest) Send(ctx context.Context) (*DeregisterOnPremisesInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterOnPremisesInstanceResponse{
		DeregisterOnPremisesInstanceOutput: r.Request.Data.(*types.DeregisterOnPremisesInstanceOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterOnPremisesInstanceResponse is the response type for the
// DeregisterOnPremisesInstance API operation.
type DeregisterOnPremisesInstanceResponse struct {
	*types.DeregisterOnPremisesInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterOnPremisesInstance request.
func (r *DeregisterOnPremisesInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
