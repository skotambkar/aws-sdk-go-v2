// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opApplyPendingMaintenanceAction = "ApplyPendingMaintenanceAction"

// ApplyPendingMaintenanceActionRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Applies a pending maintenance action to a resource (for example, to a replication
// instance).
//
//    // Example sending a request using ApplyPendingMaintenanceActionRequest.
//    req := client.ApplyPendingMaintenanceActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ApplyPendingMaintenanceAction
func (c *Client) ApplyPendingMaintenanceActionRequest(input *types.ApplyPendingMaintenanceActionInput) ApplyPendingMaintenanceActionRequest {
	op := &aws.Operation{
		Name:       opApplyPendingMaintenanceAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ApplyPendingMaintenanceActionInput{}
	}

	req := c.newRequest(op, input, &types.ApplyPendingMaintenanceActionOutput{})
	return ApplyPendingMaintenanceActionRequest{Request: req, Input: input, Copy: c.ApplyPendingMaintenanceActionRequest}
}

// ApplyPendingMaintenanceActionRequest is the request type for the
// ApplyPendingMaintenanceAction API operation.
type ApplyPendingMaintenanceActionRequest struct {
	*aws.Request
	Input *types.ApplyPendingMaintenanceActionInput
	Copy  func(*types.ApplyPendingMaintenanceActionInput) ApplyPendingMaintenanceActionRequest
}

// Send marshals and sends the ApplyPendingMaintenanceAction API request.
func (r ApplyPendingMaintenanceActionRequest) Send(ctx context.Context) (*ApplyPendingMaintenanceActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ApplyPendingMaintenanceActionResponse{
		ApplyPendingMaintenanceActionOutput: r.Request.Data.(*types.ApplyPendingMaintenanceActionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ApplyPendingMaintenanceActionResponse is the response type for the
// ApplyPendingMaintenanceAction API operation.
type ApplyPendingMaintenanceActionResponse struct {
	*types.ApplyPendingMaintenanceActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ApplyPendingMaintenanceAction request.
func (r *ApplyPendingMaintenanceActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
