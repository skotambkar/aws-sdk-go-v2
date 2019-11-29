// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opRegisterTaskWithMaintenanceWindow = "RegisterTaskWithMaintenanceWindow"

// RegisterTaskWithMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Adds a new task to a maintenance window.
//
//    // Example sending a request using RegisterTaskWithMaintenanceWindowRequest.
//    req := client.RegisterTaskWithMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/RegisterTaskWithMaintenanceWindow
func (c *Client) RegisterTaskWithMaintenanceWindowRequest(input *types.RegisterTaskWithMaintenanceWindowInput) RegisterTaskWithMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opRegisterTaskWithMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterTaskWithMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &types.RegisterTaskWithMaintenanceWindowOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RegisterTaskWithMaintenanceWindowMarshaler{Input: input}.GetNamedBuildHandler())

	return RegisterTaskWithMaintenanceWindowRequest{Request: req, Input: input, Copy: c.RegisterTaskWithMaintenanceWindowRequest}
}

// RegisterTaskWithMaintenanceWindowRequest is the request type for the
// RegisterTaskWithMaintenanceWindow API operation.
type RegisterTaskWithMaintenanceWindowRequest struct {
	*aws.Request
	Input *types.RegisterTaskWithMaintenanceWindowInput
	Copy  func(*types.RegisterTaskWithMaintenanceWindowInput) RegisterTaskWithMaintenanceWindowRequest
}

// Send marshals and sends the RegisterTaskWithMaintenanceWindow API request.
func (r RegisterTaskWithMaintenanceWindowRequest) Send(ctx context.Context) (*RegisterTaskWithMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterTaskWithMaintenanceWindowResponse{
		RegisterTaskWithMaintenanceWindowOutput: r.Request.Data.(*types.RegisterTaskWithMaintenanceWindowOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterTaskWithMaintenanceWindowResponse is the response type for the
// RegisterTaskWithMaintenanceWindow API operation.
type RegisterTaskWithMaintenanceWindowResponse struct {
	*types.RegisterTaskWithMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterTaskWithMaintenanceWindow request.
func (r *RegisterTaskWithMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
