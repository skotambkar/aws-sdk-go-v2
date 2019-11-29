// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDeregisterTargetFromMaintenanceWindow = "DeregisterTargetFromMaintenanceWindow"

// DeregisterTargetFromMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Removes a target from a maintenance window.
//
//    // Example sending a request using DeregisterTargetFromMaintenanceWindowRequest.
//    req := client.DeregisterTargetFromMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterTargetFromMaintenanceWindow
func (c *Client) DeregisterTargetFromMaintenanceWindowRequest(input *types.DeregisterTargetFromMaintenanceWindowInput) DeregisterTargetFromMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opDeregisterTargetFromMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeregisterTargetFromMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &types.DeregisterTargetFromMaintenanceWindowOutput{})
	return DeregisterTargetFromMaintenanceWindowRequest{Request: req, Input: input, Copy: c.DeregisterTargetFromMaintenanceWindowRequest}
}

// DeregisterTargetFromMaintenanceWindowRequest is the request type for the
// DeregisterTargetFromMaintenanceWindow API operation.
type DeregisterTargetFromMaintenanceWindowRequest struct {
	*aws.Request
	Input *types.DeregisterTargetFromMaintenanceWindowInput
	Copy  func(*types.DeregisterTargetFromMaintenanceWindowInput) DeregisterTargetFromMaintenanceWindowRequest
}

// Send marshals and sends the DeregisterTargetFromMaintenanceWindow API request.
func (r DeregisterTargetFromMaintenanceWindowRequest) Send(ctx context.Context) (*DeregisterTargetFromMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTargetFromMaintenanceWindowResponse{
		DeregisterTargetFromMaintenanceWindowOutput: r.Request.Data.(*types.DeregisterTargetFromMaintenanceWindowOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTargetFromMaintenanceWindowResponse is the response type for the
// DeregisterTargetFromMaintenanceWindow API operation.
type DeregisterTargetFromMaintenanceWindowResponse struct {
	*types.DeregisterTargetFromMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterTargetFromMaintenanceWindow request.
func (r *DeregisterTargetFromMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
