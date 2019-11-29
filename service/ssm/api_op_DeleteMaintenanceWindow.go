// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDeleteMaintenanceWindow = "DeleteMaintenanceWindow"

// DeleteMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Deletes a maintenance window.
//
//    // Example sending a request using DeleteMaintenanceWindowRequest.
//    req := client.DeleteMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteMaintenanceWindow
func (c *Client) DeleteMaintenanceWindowRequest(input *types.DeleteMaintenanceWindowInput) DeleteMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opDeleteMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &types.DeleteMaintenanceWindowOutput{})
	return DeleteMaintenanceWindowRequest{Request: req, Input: input, Copy: c.DeleteMaintenanceWindowRequest}
}

// DeleteMaintenanceWindowRequest is the request type for the
// DeleteMaintenanceWindow API operation.
type DeleteMaintenanceWindowRequest struct {
	*aws.Request
	Input *types.DeleteMaintenanceWindowInput
	Copy  func(*types.DeleteMaintenanceWindowInput) DeleteMaintenanceWindowRequest
}

// Send marshals and sends the DeleteMaintenanceWindow API request.
func (r DeleteMaintenanceWindowRequest) Send(ctx context.Context) (*DeleteMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMaintenanceWindowResponse{
		DeleteMaintenanceWindowOutput: r.Request.Data.(*types.DeleteMaintenanceWindowOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMaintenanceWindowResponse is the response type for the
// DeleteMaintenanceWindow API operation.
type DeleteMaintenanceWindowResponse struct {
	*types.DeleteMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMaintenanceWindow request.
func (r *DeleteMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
