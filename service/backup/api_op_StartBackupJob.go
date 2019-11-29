// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/backup/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opStartBackupJob = "StartBackupJob"

// StartBackupJobRequest returns a request value for making API operation for
// AWS Backup.
//
// Starts a job to create a one-time backup of the specified resource.
//
//    // Example sending a request using StartBackupJobRequest.
//    req := client.StartBackupJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/StartBackupJob
func (c *Client) StartBackupJobRequest(input *types.StartBackupJobInput) StartBackupJobRequest {
	op := &aws.Operation{
		Name:       opStartBackupJob,
		HTTPMethod: "PUT",
		HTTPPath:   "/backup-jobs",
	}

	if input == nil {
		input = &types.StartBackupJobInput{}
	}

	req := c.newRequest(op, input, &types.StartBackupJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.StartBackupJobMarshaler{Input: input}.GetNamedBuildHandler())

	return StartBackupJobRequest{Request: req, Input: input, Copy: c.StartBackupJobRequest}
}

// StartBackupJobRequest is the request type for the
// StartBackupJob API operation.
type StartBackupJobRequest struct {
	*aws.Request
	Input *types.StartBackupJobInput
	Copy  func(*types.StartBackupJobInput) StartBackupJobRequest
}

// Send marshals and sends the StartBackupJob API request.
func (r StartBackupJobRequest) Send(ctx context.Context) (*StartBackupJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartBackupJobResponse{
		StartBackupJobOutput: r.Request.Data.(*types.StartBackupJobOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartBackupJobResponse is the response type for the
// StartBackupJob API operation.
type StartBackupJobResponse struct {
	*types.StartBackupJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartBackupJob request.
func (r *StartBackupJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
