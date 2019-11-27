// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opGetBackupSelection = "GetBackupSelection"

// GetBackupSelectionRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns selection metadata and a document in JSON format that specifies a
// list of resources that are associated with a backup plan.
//
//    // Example sending a request using GetBackupSelectionRequest.
//    req := client.GetBackupSelectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupSelection
func (c *Client) GetBackupSelectionRequest(input *types.GetBackupSelectionInput) GetBackupSelectionRequest {
	op := &aws.Operation{
		Name:       opGetBackupSelection,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/{backupPlanId}/selections/{selectionId}",
	}

	if input == nil {
		input = &types.GetBackupSelectionInput{}
	}

	req := c.newRequest(op, input, &types.GetBackupSelectionOutput{})
	return GetBackupSelectionRequest{Request: req, Input: input, Copy: c.GetBackupSelectionRequest}
}

// GetBackupSelectionRequest is the request type for the
// GetBackupSelection API operation.
type GetBackupSelectionRequest struct {
	*aws.Request
	Input *types.GetBackupSelectionInput
	Copy  func(*types.GetBackupSelectionInput) GetBackupSelectionRequest
}

// Send marshals and sends the GetBackupSelection API request.
func (r GetBackupSelectionRequest) Send(ctx context.Context) (*GetBackupSelectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackupSelectionResponse{
		GetBackupSelectionOutput: r.Request.Data.(*types.GetBackupSelectionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackupSelectionResponse is the response type for the
// GetBackupSelection API operation.
type GetBackupSelectionResponse struct {
	*types.GetBackupSelectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackupSelection request.
func (r *GetBackupSelectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
