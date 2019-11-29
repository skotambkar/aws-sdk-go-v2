// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opUpdateMaintenanceWindowTarget = "UpdateMaintenanceWindowTarget"

// UpdateMaintenanceWindowTargetRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Modifies the target of an existing maintenance window. You can change the
// following:
//
//    * Name
//
//    * Description
//
//    * Owner
//
//    * IDs for an ID target
//
//    * Tags for a Tag target
//
//    * From any supported tag type to another. The three supported tag types
//    are ID target, Tag target, and resource group. For more information, see
//    Target.
//
// If a parameter is null, then the corresponding field is not modified.
//
//    // Example sending a request using UpdateMaintenanceWindowTargetRequest.
//    req := client.UpdateMaintenanceWindowTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateMaintenanceWindowTarget
func (c *Client) UpdateMaintenanceWindowTargetRequest(input *types.UpdateMaintenanceWindowTargetInput) UpdateMaintenanceWindowTargetRequest {
	op := &aws.Operation{
		Name:       opUpdateMaintenanceWindowTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateMaintenanceWindowTargetInput{}
	}

	req := c.newRequest(op, input, &types.UpdateMaintenanceWindowTargetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateMaintenanceWindowTargetMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateMaintenanceWindowTargetRequest{Request: req, Input: input, Copy: c.UpdateMaintenanceWindowTargetRequest}
}

// UpdateMaintenanceWindowTargetRequest is the request type for the
// UpdateMaintenanceWindowTarget API operation.
type UpdateMaintenanceWindowTargetRequest struct {
	*aws.Request
	Input *types.UpdateMaintenanceWindowTargetInput
	Copy  func(*types.UpdateMaintenanceWindowTargetInput) UpdateMaintenanceWindowTargetRequest
}

// Send marshals and sends the UpdateMaintenanceWindowTarget API request.
func (r UpdateMaintenanceWindowTargetRequest) Send(ctx context.Context) (*UpdateMaintenanceWindowTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMaintenanceWindowTargetResponse{
		UpdateMaintenanceWindowTargetOutput: r.Request.Data.(*types.UpdateMaintenanceWindowTargetOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMaintenanceWindowTargetResponse is the response type for the
// UpdateMaintenanceWindowTarget API operation.
type UpdateMaintenanceWindowTargetResponse struct {
	*types.UpdateMaintenanceWindowTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMaintenanceWindowTarget request.
func (r *UpdateMaintenanceWindowTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
