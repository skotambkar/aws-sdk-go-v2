// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opDescribeRestoreJob = "DescribeRestoreJob"

// DescribeRestoreJobRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns metadata associated with a restore job that is specified by a job
// ID.
//
//    // Example sending a request using DescribeRestoreJobRequest.
//    req := client.DescribeRestoreJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeRestoreJob
func (c *Client) DescribeRestoreJobRequest(input *types.DescribeRestoreJobInput) DescribeRestoreJobRequest {
	op := &aws.Operation{
		Name:       opDescribeRestoreJob,
		HTTPMethod: "GET",
		HTTPPath:   "/restore-jobs/{restoreJobId}",
	}

	if input == nil {
		input = &types.DescribeRestoreJobInput{}
	}

	req := c.newRequest(op, input, &types.DescribeRestoreJobOutput{})
	return DescribeRestoreJobRequest{Request: req, Input: input, Copy: c.DescribeRestoreJobRequest}
}

// DescribeRestoreJobRequest is the request type for the
// DescribeRestoreJob API operation.
type DescribeRestoreJobRequest struct {
	*aws.Request
	Input *types.DescribeRestoreJobInput
	Copy  func(*types.DescribeRestoreJobInput) DescribeRestoreJobRequest
}

// Send marshals and sends the DescribeRestoreJob API request.
func (r DescribeRestoreJobRequest) Send(ctx context.Context) (*DescribeRestoreJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRestoreJobResponse{
		DescribeRestoreJobOutput: r.Request.Data.(*types.DescribeRestoreJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRestoreJobResponse is the response type for the
// DescribeRestoreJob API operation.
type DescribeRestoreJobResponse struct {
	*types.DescribeRestoreJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRestoreJob request.
func (r *DescribeRestoreJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
