// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opCreateSnapshotSchedule = "CreateSnapshotSchedule"

// CreateSnapshotScheduleRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates a new snapshot schedule.
//
//    // Example sending a request using CreateSnapshotScheduleRequest.
//    req := client.CreateSnapshotScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateSnapshotSchedule
func (c *Client) CreateSnapshotScheduleRequest(input *types.CreateSnapshotScheduleInput) CreateSnapshotScheduleRequest {
	op := &aws.Operation{
		Name:       opCreateSnapshotSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateSnapshotScheduleInput{}
	}

	req := c.newRequest(op, input, &types.CreateSnapshotScheduleOutput{})
	return CreateSnapshotScheduleRequest{Request: req, Input: input, Copy: c.CreateSnapshotScheduleRequest}
}

// CreateSnapshotScheduleRequest is the request type for the
// CreateSnapshotSchedule API operation.
type CreateSnapshotScheduleRequest struct {
	*aws.Request
	Input *types.CreateSnapshotScheduleInput
	Copy  func(*types.CreateSnapshotScheduleInput) CreateSnapshotScheduleRequest
}

// Send marshals and sends the CreateSnapshotSchedule API request.
func (r CreateSnapshotScheduleRequest) Send(ctx context.Context) (*CreateSnapshotScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSnapshotScheduleResponse{
		CreateSnapshotScheduleOutput: r.Request.Data.(*types.CreateSnapshotScheduleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSnapshotScheduleResponse is the response type for the
// CreateSnapshotSchedule API operation.
type CreateSnapshotScheduleResponse struct {
	*types.CreateSnapshotScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSnapshotSchedule request.
func (r *CreateSnapshotScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
