// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opDeleteDBInstanceAutomatedBackup = "DeleteDBInstanceAutomatedBackup"

// DeleteDBInstanceAutomatedBackupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes automated backups based on the source instance's DbiResourceId value
// or the restorable instance's resource ID.
//
//    // Example sending a request using DeleteDBInstanceAutomatedBackupRequest.
//    req := client.DeleteDBInstanceAutomatedBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBInstanceAutomatedBackup
func (c *Client) DeleteDBInstanceAutomatedBackupRequest(input *types.DeleteDBInstanceAutomatedBackupInput) DeleteDBInstanceAutomatedBackupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBInstanceAutomatedBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteDBInstanceAutomatedBackupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDBInstanceAutomatedBackupOutput{})
	return DeleteDBInstanceAutomatedBackupRequest{Request: req, Input: input, Copy: c.DeleteDBInstanceAutomatedBackupRequest}
}

// DeleteDBInstanceAutomatedBackupRequest is the request type for the
// DeleteDBInstanceAutomatedBackup API operation.
type DeleteDBInstanceAutomatedBackupRequest struct {
	*aws.Request
	Input *types.DeleteDBInstanceAutomatedBackupInput
	Copy  func(*types.DeleteDBInstanceAutomatedBackupInput) DeleteDBInstanceAutomatedBackupRequest
}

// Send marshals and sends the DeleteDBInstanceAutomatedBackup API request.
func (r DeleteDBInstanceAutomatedBackupRequest) Send(ctx context.Context) (*DeleteDBInstanceAutomatedBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBInstanceAutomatedBackupResponse{
		DeleteDBInstanceAutomatedBackupOutput: r.Request.Data.(*types.DeleteDBInstanceAutomatedBackupOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBInstanceAutomatedBackupResponse is the response type for the
// DeleteDBInstanceAutomatedBackup API operation.
type DeleteDBInstanceAutomatedBackupResponse struct {
	*types.DeleteDBInstanceAutomatedBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBInstanceAutomatedBackup request.
func (r *DeleteDBInstanceAutomatedBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
