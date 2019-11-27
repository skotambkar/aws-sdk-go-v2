// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
)

const opDescribeBackups = "DescribeBackups"

// DescribeBackupsRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Describes backups. The results are ordered by time, with newest backups first.
// If you do not specify a BackupId or ServerName, the command returns all backups.
//
// This operation is synchronous.
//
// A ResourceNotFoundException is thrown when the backup does not exist. A ValidationException
// is raised when parameters of the request are not valid.
//
//    // Example sending a request using DescribeBackupsRequest.
//    req := client.DescribeBackupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DescribeBackups
func (c *Client) DescribeBackupsRequest(input *types.DescribeBackupsInput) DescribeBackupsRequest {
	op := &aws.Operation{
		Name:       opDescribeBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeBackupsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeBackupsOutput{})
	return DescribeBackupsRequest{Request: req, Input: input, Copy: c.DescribeBackupsRequest}
}

// DescribeBackupsRequest is the request type for the
// DescribeBackups API operation.
type DescribeBackupsRequest struct {
	*aws.Request
	Input *types.DescribeBackupsInput
	Copy  func(*types.DescribeBackupsInput) DescribeBackupsRequest
}

// Send marshals and sends the DescribeBackups API request.
func (r DescribeBackupsRequest) Send(ctx context.Context) (*DescribeBackupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBackupsResponse{
		DescribeBackupsOutput: r.Request.Data.(*types.DescribeBackupsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBackupsResponse is the response type for the
// DescribeBackups API operation.
type DescribeBackupsResponse struct {
	*types.DescribeBackupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBackups request.
func (r *DescribeBackupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
