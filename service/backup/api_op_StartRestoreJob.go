// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opStartRestoreJob = "StartRestoreJob"

// StartRestoreJobRequest returns a request value for making API operation for
// AWS Backup.
//
// Recovers the saved resource identified by an Amazon Resource Name (ARN).
//
// If the resource ARN is included in the request, then the last complete backup
// of that resource is recovered. If the ARN of a recovery point is supplied,
// then that recovery point is restored.
//
//    // Example sending a request using StartRestoreJobRequest.
//    req := client.StartRestoreJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/StartRestoreJob
func (c *Client) StartRestoreJobRequest(input *types.StartRestoreJobInput) StartRestoreJobRequest {
	op := &aws.Operation{
		Name:       opStartRestoreJob,
		HTTPMethod: "PUT",
		HTTPPath:   "/restore-jobs",
	}

	if input == nil {
		input = &types.StartRestoreJobInput{}
	}

	req := c.newRequest(op, input, &types.StartRestoreJobOutput{})
	return StartRestoreJobRequest{Request: req, Input: input, Copy: c.StartRestoreJobRequest}
}

// StartRestoreJobRequest is the request type for the
// StartRestoreJob API operation.
type StartRestoreJobRequest struct {
	*aws.Request
	Input *types.StartRestoreJobInput
	Copy  func(*types.StartRestoreJobInput) StartRestoreJobRequest
}

// Send marshals and sends the StartRestoreJob API request.
func (r StartRestoreJobRequest) Send(ctx context.Context) (*StartRestoreJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartRestoreJobResponse{
		StartRestoreJobOutput: r.Request.Data.(*types.StartRestoreJobOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartRestoreJobResponse is the response type for the
// StartRestoreJob API operation.
type StartRestoreJobResponse struct {
	*types.StartRestoreJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartRestoreJob request.
func (r *StartRestoreJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
