// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opRegisterDefaultPatchBaseline = "RegisterDefaultPatchBaseline"

// RegisterDefaultPatchBaselineRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Defines the default patch baseline for the relevant operating system.
//
// To reset the AWS predefined patch baseline as the default, specify the full
// patch baseline ARN as the baseline ID value. For example, for CentOS, specify
// arn:aws:ssm:us-east-2:733109147000:patchbaseline/pb-0574b43a65ea646ed instead
// of pb-0574b43a65ea646ed.
//
//    // Example sending a request using RegisterDefaultPatchBaselineRequest.
//    req := client.RegisterDefaultPatchBaselineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/RegisterDefaultPatchBaseline
func (c *Client) RegisterDefaultPatchBaselineRequest(input *types.RegisterDefaultPatchBaselineInput) RegisterDefaultPatchBaselineRequest {
	op := &aws.Operation{
		Name:       opRegisterDefaultPatchBaseline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterDefaultPatchBaselineInput{}
	}

	req := c.newRequest(op, input, &types.RegisterDefaultPatchBaselineOutput{})
	return RegisterDefaultPatchBaselineRequest{Request: req, Input: input, Copy: c.RegisterDefaultPatchBaselineRequest}
}

// RegisterDefaultPatchBaselineRequest is the request type for the
// RegisterDefaultPatchBaseline API operation.
type RegisterDefaultPatchBaselineRequest struct {
	*aws.Request
	Input *types.RegisterDefaultPatchBaselineInput
	Copy  func(*types.RegisterDefaultPatchBaselineInput) RegisterDefaultPatchBaselineRequest
}

// Send marshals and sends the RegisterDefaultPatchBaseline API request.
func (r RegisterDefaultPatchBaselineRequest) Send(ctx context.Context) (*RegisterDefaultPatchBaselineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterDefaultPatchBaselineResponse{
		RegisterDefaultPatchBaselineOutput: r.Request.Data.(*types.RegisterDefaultPatchBaselineOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterDefaultPatchBaselineResponse is the response type for the
// RegisterDefaultPatchBaseline API operation.
type RegisterDefaultPatchBaselineResponse struct {
	*types.RegisterDefaultPatchBaselineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterDefaultPatchBaseline request.
func (r *RegisterDefaultPatchBaselineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
