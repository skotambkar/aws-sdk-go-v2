// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
)

const opModifyHapg = "ModifyHapg"

// ModifyHapgRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Modifies an existing high-availability partition group.
//
//    // Example sending a request using ModifyHapgRequest.
//    req := client.ModifyHapgRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ModifyHapg
func (c *Client) ModifyHapgRequest(input *types.ModifyHapgInput) ModifyHapgRequest {
	op := &aws.Operation{
		Name:       opModifyHapg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyHapgInput{}
	}

	req := c.newRequest(op, input, &types.ModifyHapgOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ModifyHapgMarshaler{Input: input}.GetNamedBuildHandler())

	return ModifyHapgRequest{Request: req, Input: input, Copy: c.ModifyHapgRequest}
}

// ModifyHapgRequest is the request type for the
// ModifyHapg API operation.
type ModifyHapgRequest struct {
	*aws.Request
	Input *types.ModifyHapgInput
	Copy  func(*types.ModifyHapgInput) ModifyHapgRequest
}

// Send marshals and sends the ModifyHapg API request.
func (r ModifyHapgRequest) Send(ctx context.Context) (*ModifyHapgResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyHapgResponse{
		ModifyHapgOutput: r.Request.Data.(*types.ModifyHapgOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyHapgResponse is the response type for the
// ModifyHapg API operation.
type ModifyHapgResponse struct {
	*types.ModifyHapgOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyHapg request.
func (r *ModifyHapgResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
