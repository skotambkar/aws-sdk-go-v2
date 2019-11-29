// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opUpdateMaintenanceStartTime = "UpdateMaintenanceStartTime"

// UpdateMaintenanceStartTimeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates a gateway's weekly maintenance start time information, including
// day and time of the week. The maintenance time is the time in your gateway's
// time zone.
//
//    // Example sending a request using UpdateMaintenanceStartTimeRequest.
//    req := client.UpdateMaintenanceStartTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateMaintenanceStartTime
func (c *Client) UpdateMaintenanceStartTimeRequest(input *types.UpdateMaintenanceStartTimeInput) UpdateMaintenanceStartTimeRequest {
	op := &aws.Operation{
		Name:       opUpdateMaintenanceStartTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateMaintenanceStartTimeInput{}
	}

	req := c.newRequest(op, input, &types.UpdateMaintenanceStartTimeOutput{})
	return UpdateMaintenanceStartTimeRequest{Request: req, Input: input, Copy: c.UpdateMaintenanceStartTimeRequest}
}

// UpdateMaintenanceStartTimeRequest is the request type for the
// UpdateMaintenanceStartTime API operation.
type UpdateMaintenanceStartTimeRequest struct {
	*aws.Request
	Input *types.UpdateMaintenanceStartTimeInput
	Copy  func(*types.UpdateMaintenanceStartTimeInput) UpdateMaintenanceStartTimeRequest
}

// Send marshals and sends the UpdateMaintenanceStartTime API request.
func (r UpdateMaintenanceStartTimeRequest) Send(ctx context.Context) (*UpdateMaintenanceStartTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMaintenanceStartTimeResponse{
		UpdateMaintenanceStartTimeOutput: r.Request.Data.(*types.UpdateMaintenanceStartTimeOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMaintenanceStartTimeResponse is the response type for the
// UpdateMaintenanceStartTime API operation.
type UpdateMaintenanceStartTimeResponse struct {
	*types.UpdateMaintenanceStartTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMaintenanceStartTime request.
func (r *UpdateMaintenanceStartTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
