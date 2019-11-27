// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opDeleteSnapshotSchedule = "DeleteSnapshotSchedule"

// DeleteSnapshotScheduleRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Deletes a snapshot schedule.
//
//    // Example sending a request using DeleteSnapshotScheduleRequest.
//    req := client.DeleteSnapshotScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DeleteSnapshotSchedule
func (c *Client) DeleteSnapshotScheduleRequest(input *types.DeleteSnapshotScheduleInput) DeleteSnapshotScheduleRequest {
	op := &aws.Operation{
		Name:       opDeleteSnapshotSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteSnapshotScheduleInput{}
	}

	req := c.newRequest(op, input, &types.DeleteSnapshotScheduleOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteSnapshotScheduleRequest{Request: req, Input: input, Copy: c.DeleteSnapshotScheduleRequest}
}

// DeleteSnapshotScheduleRequest is the request type for the
// DeleteSnapshotSchedule API operation.
type DeleteSnapshotScheduleRequest struct {
	*aws.Request
	Input *types.DeleteSnapshotScheduleInput
	Copy  func(*types.DeleteSnapshotScheduleInput) DeleteSnapshotScheduleRequest
}

// Send marshals and sends the DeleteSnapshotSchedule API request.
func (r DeleteSnapshotScheduleRequest) Send(ctx context.Context) (*DeleteSnapshotScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSnapshotScheduleResponse{
		DeleteSnapshotScheduleOutput: r.Request.Data.(*types.DeleteSnapshotScheduleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSnapshotScheduleResponse is the response type for the
// DeleteSnapshotSchedule API operation.
type DeleteSnapshotScheduleResponse struct {
	*types.DeleteSnapshotScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSnapshotSchedule request.
func (r *DeleteSnapshotScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
