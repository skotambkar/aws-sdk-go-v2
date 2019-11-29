// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opUpdateTerminationProtection = "UpdateTerminationProtection"

// UpdateTerminationProtectionRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Updates termination protection for the specified stack. If a user attempts
// to delete a stack with termination protection enabled, the operation fails
// and the stack remains unchanged. For more information, see Protecting a Stack
// From Being Deleted (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html)
// in the AWS CloudFormation User Guide.
//
// For nested stacks (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html),
// termination protection is set on the root stack and cannot be changed directly
// on the nested stack.
//
//    // Example sending a request using UpdateTerminationProtectionRequest.
//    req := client.UpdateTerminationProtectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/UpdateTerminationProtection
func (c *Client) UpdateTerminationProtectionRequest(input *types.UpdateTerminationProtectionInput) UpdateTerminationProtectionRequest {
	op := &aws.Operation{
		Name:       opUpdateTerminationProtection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateTerminationProtectionInput{}
	}

	req := c.newRequest(op, input, &types.UpdateTerminationProtectionOutput{})
	return UpdateTerminationProtectionRequest{Request: req, Input: input, Copy: c.UpdateTerminationProtectionRequest}
}

// UpdateTerminationProtectionRequest is the request type for the
// UpdateTerminationProtection API operation.
type UpdateTerminationProtectionRequest struct {
	*aws.Request
	Input *types.UpdateTerminationProtectionInput
	Copy  func(*types.UpdateTerminationProtectionInput) UpdateTerminationProtectionRequest
}

// Send marshals and sends the UpdateTerminationProtection API request.
func (r UpdateTerminationProtectionRequest) Send(ctx context.Context) (*UpdateTerminationProtectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTerminationProtectionResponse{
		UpdateTerminationProtectionOutput: r.Request.Data.(*types.UpdateTerminationProtectionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTerminationProtectionResponse is the response type for the
// UpdateTerminationProtection API operation.
type UpdateTerminationProtectionResponse struct {
	*types.UpdateTerminationProtectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTerminationProtection request.
func (r *UpdateTerminationProtectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
