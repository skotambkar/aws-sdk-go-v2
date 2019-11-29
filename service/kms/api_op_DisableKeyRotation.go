// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

const opDisableKeyRotation = "DisableKeyRotation"

// DisableKeyRotationRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Disables automatic rotation of the key material (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
// for the specified customer master key (CMK). You cannot perform this operation
// on a CMK in a different AWS account.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using DisableKeyRotationRequest.
//    req := client.DisableKeyRotationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/DisableKeyRotation
func (c *Client) DisableKeyRotationRequest(input *types.DisableKeyRotationInput) DisableKeyRotationRequest {
	op := &aws.Operation{
		Name:       opDisableKeyRotation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisableKeyRotationInput{}
	}

	req := c.newRequest(op, input, &types.DisableKeyRotationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DisableKeyRotationMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DisableKeyRotationRequest{Request: req, Input: input, Copy: c.DisableKeyRotationRequest}
}

// DisableKeyRotationRequest is the request type for the
// DisableKeyRotation API operation.
type DisableKeyRotationRequest struct {
	*aws.Request
	Input *types.DisableKeyRotationInput
	Copy  func(*types.DisableKeyRotationInput) DisableKeyRotationRequest
}

// Send marshals and sends the DisableKeyRotation API request.
func (r DisableKeyRotationRequest) Send(ctx context.Context) (*DisableKeyRotationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableKeyRotationResponse{
		DisableKeyRotationOutput: r.Request.Data.(*types.DisableKeyRotationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableKeyRotationResponse is the response type for the
// DisableKeyRotation API operation.
type DisableKeyRotationResponse struct {
	*types.DisableKeyRotationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableKeyRotation request.
func (r *DisableKeyRotationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
