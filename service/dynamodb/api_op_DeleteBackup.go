// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const opDeleteBackup = "DeleteBackup"

// DeleteBackupRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Deletes an existing backup of a table.
//
// You can call DeleteBackup at a maximum rate of 10 times per second.
//
//    // Example sending a request using DeleteBackupRequest.
//    req := client.DeleteBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DeleteBackup
func (c *Client) DeleteBackupRequest(input *types.DeleteBackupInput) DeleteBackupRequest {
	op := &aws.Operation{
		Name:       opDeleteBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteBackupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBackupOutput{})
	return DeleteBackupRequest{Request: req, Input: input, Copy: c.DeleteBackupRequest}
}

// DeleteBackupRequest is the request type for the
// DeleteBackup API operation.
type DeleteBackupRequest struct {
	*aws.Request
	Input *types.DeleteBackupInput
	Copy  func(*types.DeleteBackupInput) DeleteBackupRequest
}

// Send marshals and sends the DeleteBackup API request.
func (r DeleteBackupRequest) Send(ctx context.Context) (*DeleteBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupResponse{
		DeleteBackupOutput: r.Request.Data.(*types.DeleteBackupOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupResponse is the response type for the
// DeleteBackup API operation.
type DeleteBackupResponse struct {
	*types.DeleteBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackup request.
func (r *DeleteBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
